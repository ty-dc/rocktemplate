package cmd

import (
	"context"
	"fmt"
	"github.com/spidernet-io/rocktemplate/pkg/k8s/client/informers/externalversions"
	"github.com/spidernet-io/rocktemplate/pkg/lease"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"os"
	"time"
)

type ExampleInformer struct {
	logger *zap.Logger
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
	clientset, err := kubernetes.NewForConfig(config) // 初始化 client
	if err != nil {
		s.logger.Sugar().Fatalf("failed to NewForConfig, reason=%v", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	leaseNamespace := "kube-system"
	leaseName := "testlease"
	rlogger := s.logger.Named(fmt.Sprintf("lease %s/%s", leaseNamespace, leaseName))
	getLease, lossLease, err := lease.NewLeaseElector(ctx, leaseNamespace, leaseName, os.Hostname(), rlogger)
	if err != nil {
		s.logger.Sugar().Fatalf("failed to generate lease, reason=%v ", err)
	}
	<-getLease

	stopInfomer := make(chan struct{})
	go func(lossLease chan struct{}) {
		<-lossLease
		close(stopInfomer)
	}(lossLease)

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

func SetupExampleInformer(logger *zap.Logger) {
	s := ExampleInformer{
		logger: logger,
	}
	go func {
		for {
			s.RunInformer()
			time.Sleep(time.Duration(5) * time.Second)
		}
	}()
}
