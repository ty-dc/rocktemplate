package cmd

import (
	"context"
	"fmt"
	crd "github.com/spidernet-io/rocktemplate/pkg/k8s/apis/rocktemplate.spidernet.io/v1"
	"go.uber.org/zap"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"time"
)

type webhookhander struct {
	logger *zap.Logger
}

var _ webhook.CustomValidator = (*webhookhander)(nil)

// mutating webhook
func (s *webhookhander) Default(ctx context.Context, obj runtime.Object) error {
	logger := s.logger.Named("mutating wehbook")

	r, ok := obj.(*crd.Mybook)
	if !ok {
		s := fmt.Sprintf("failed to get obj")
		logger.Error(s)
		return apierrors.NewBadRequest(s)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return nil

}

func (s *webhookhander) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	logger := s.logger.Named("validating create wehbook")

	r, ok := obj.(*crd.Mybook)
	if !ok {
		s := fmt.Sprintf("failed to get obj")
		logger.Error(s)
		return apierrors.NewBadRequest(s)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return nil
}

func (s *webhookhander) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	logger := s.logger.Named("validating update wehbook")

	old, ok := oldObj.(*crd.Mybook)
	if !ok {
		s := fmt.Sprintf("failed to get oldObj")
		logger.Error(s)
		return apierrors.NewBadRequest(s)
	}
	new, ok := newObj.(*crd.Mybook)
	if !ok {
		s := fmt.Sprintf("failed to get newObj")
		logger.Error(s)
		return apierrors.NewBadRequest(s)
	}
	logger.Sugar().Infof("oldObj: %+v", old)
	logger.Sugar().Infof("newObj: %+v", new)

	return nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type
func (s *webhookhander) ValidateDelete(ctx context.Context, obj runtime.Object) error {
	logger := s.logger.Named("validating delete wehbook")

	r, ok := obj.(*crd.Mybook)
	if !ok {
		s := fmt.Sprintf("failed to get obj")
		logger.Error(s)
		return apierrors.NewBadRequest(s)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return nil
}

// func runWebhookServer(mux *http.ServeMux, webhookPort int, keyPath, certPath string, logger *zap.Logger) {
//
// 	logger.Sugar().Infof("setup webhook on port %v, with tls under %v", webhookPort, tlsDir)
//
// 	ctx, cancel := context.WithCancel(context.Background())
//
// 	// tls
// 	certWatcher, err := certwatcher.New(certPath, keyPath)
// 	if err != nil {
// 		logger.Sugar().Fatalf("failed to certwatcher, reason=%v", err)
// 	}
// 	go func() {
// 		certWatcher.Start(ctx)
// 	}()
//
// 	tlscfg := &tls.Config{
// 		// NextProtos:     []string{"h2"},
// 		GetCertificate:     certWatcher.GetCertificate,
// 		MinVersion:         tls.VersionTLS12,
// 		InsecureSkipVerify: true,
// 	}
//
// 	srv := &http.Server{
// 		Addr:              fmt.Sprintf(":%v", webhookPort),
// 		TLSConfig:         tlscfg,
// 		Handler:           mux,
// 		MaxHeaderBytes:    1 << 20,
// 		IdleTimeout:       90 * time.Second, // matches http.DefaultTransport keep-alive timeout
// 		ReadHeaderTimeout: 32 * time.Second,
// 	}
//
// 	s := "wehbhook server exit"
// 	if e := srv.ListenAndServe(); e != nil {
// 		s += fmt.Sprintf(", reason=%v", e)
// 	}
// 	// cancel tls watch
// 	cancel()
// 	logger.Error(s)
//
// }
//
// type WebhookValidating struct {
// 	logger *zap.Logger
// }
//
// func SetuptExampleWebhook(webhookPort int, keyPath, certPath string, logger *zap.Logger) {
//
// 	schema := runtime.NewScheme()
// 	if e := crd.AddToScheme(schema); e != nil {
// 		logger.Sugar().Fatalf("failed to add crd schema, reason=%v", e)
// 	}
// 	t := admission.Webhook{
// 		Handler:
// 	}
// 	hook, err := admission.StandaloneWebhook(t, admission.StandaloneOptions{
// 		Scheme: schema,
// 		Logger: logger.Named("validating wehbhook"),
// 	})
// 	if err != nil {
// 		logger.Sugar().Fatalf("failed to StandaloneWebhook, reason=%v", err)
// 	}
//
// 	mux := http.NewServeMux()
// 	mux.Handle("/failing", hook)
//
// 	runWebhookServer(mux, webhookPort, keyPath, certPath, logger)
//
// }

// https://github.com/kubernetes-sigs/controller-runtime/blob/master/pkg/builder/example_webhook_test.go
// https://github.com/kubernetes-sigs/controller-runtime/blob/master/pkg/builder/webhook_test.go
func SetupExampleWebhook(webhookPort int, tlsDir string, logger *zap.Logger) {
	logger.Sugar().Infof("setup webhook on port %v, with tls under %v", webhookPort, tlsDir)

	scheme := runtime.NewScheme()
	if e := crd.AddToScheme(scheme); e != nil {
		logger.Sugar().Fatalf("failed to add crd scheme, reason=%v", e)
	}
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		LeaderElection:         false,
		MetricsBindAddress:     "0",
		HealthProbeBindAddress: "0",
		// webhook port
		Port: webhookPort,
		// directory that contains the webhook server key and certificate, The server key and certificate must be named tls.key and tls.crt
		CertDir: tlsDir,
	})
	if err != nil {
		logger.Sugar().Fatalf("failed to NewManager, reason=%v", err)
	}

	r := &webhookhander{
		logger: logger,
	}
	// the mutating route path : "/mutate-" + strings.ReplaceAll(gvk.Group, ".", "-") + "-" + gvk.Version + "-" + strings.ToLower(gvk.Kind)
	// the validate route path : "/validate-" + strings.ReplaceAll(gvk.Group, ".", "-") + "-" + gvk.Version + "-" + strings.ToLower(gvk.Kind)
	e := ctrl.NewWebhookManagedBy(mgr).
		For(&crd.Mybook{}).
		WithDefaulter(r).
		WithValidator(r).
		Complete()
	if e != nil {
		logger.Sugar().Fatalf("failed to NewWebhookManagedBy, reason=%v", e)
	}

	// server := mgr.GetWebhookServer()
	// mgr.Start()

	go func() {
		logger.Info("start wehbhook server")
		if err := mgr.Start(context.Background()); err != nil {
			logger.Sugar().Errorf("wehbhook down, reason=%v", err)
		}
		time.Sleep(time.Second)
	}()

}
