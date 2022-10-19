package cmd

import (
	"context"
	"fmt"
	crdclientset "github.com/spidernet-io/rocktemplate/pkg/k8s/client/clientset/versioned"
	"github.com/spidernet-io/rocktemplate/pkg/k8s/client/informers/externalversions"
	"github.com/spidernet-io/rocktemplate/pkg/lease"
	"go.uber.org/zap"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"time"
)

type ExampleInformer struct {
	logger         *zap.Logger
	leaseName      string
	leaseNameSpace string
}

func (s *ExampleInformer) informerAddHandler(obj interface{}) {
	s.logger.Sugar().Infof("crd add: %+v", obj)
}

func (s *ExampleInformer) informerUpdateHandler(oldObj interface{}, newObj interface{}) {
	s.logger.Sugar().Infof("crd update old: %+v", oldObj)
	s.logger.Sugar().Infof("crd update new: %+v", newObj)
}

func (s *ExampleInformer) informerDeleteHandler(obj interface{}) {
	s.logger.Sugar().Infof("crd delete: %+v", obj)
}

func (s *ExampleInformer) RunInformer() {

	config, err := rest.InClusterConfig()
	if err != nil {
		s.logger.Sugar().Fatalf("failed to InClusterConfig, reason=%v", err)
	}
	clientset, err := crdclientset.NewForConfig(config) // 初始化 client
	if err != nil {
		s.logger.Sugar().Fatalf("failed to NewForConfig, reason=%v", err)
		return
	}

	stopInfomer := make(chan struct{})

	if len(s.leaseName) > 0 && len(s.leaseNameSpace) > 0 {
		s.logger.Sugar().Infof("try to get lease %s/%s to run informer", s.leaseNameSpace, s.leaseName)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		rlogger := s.logger.Named(fmt.Sprintf("lease %s/%s", s.leaseNameSpace, s.leaseName))
		// id, _ := os.Hostname()
		id := globalConfig.PodName
		getLease, lossLease, err := lease.NewLeaseElector(ctx, s.leaseNameSpace, s.leaseName, id, rlogger)
		if err != nil {
			s.logger.Sugar().Fatalf("failed to generate lease, reason=%v ", err)
		}
		<-getLease
		s.logger.Sugar().Infof("succeed to get lease %s/%s to run informer", s.leaseNameSpace, s.leaseName)

		go func(lossLease chan struct{}) {
			<-lossLease
			close(stopInfomer)
			s.logger.Sugar().Warnf("lease %s/%s is loss, informer is broken", s.leaseNameSpace, s.leaseName)
		}(lossLease)
	}

	// setup informer
	s.logger.Info("begin to setup informer")
	factory := externalversions.NewSharedInformerFactory(clientset, 0)
	inform := factory.Rocktemplate().V1().Mybooks().Informer()
	inform.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    s.informerAddHandler,
		UpdateFunc: s.informerUpdateHandler,
		DeleteFunc: s.informerDeleteHandler,
	})
	inform.Run(stopInfomer)

}

func SetupExampleInformer(leaseName, leaseNameSpace string, logger *zap.Logger) {
	s := ExampleInformer{
		logger:         logger,
		leaseName:      leaseName,
		leaseNameSpace: leaseNameSpace,
	}
	go func() {
		for {
			s.RunInformer()
			time.Sleep(time.Duration(5) * time.Second)
		}
	}()
}
