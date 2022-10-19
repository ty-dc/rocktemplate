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

func SetupExampleWebhook(webhookPort int, tlsDir string, logger *zap.Logger) {
	logger.Info("setup webhook")

	schema := runtime.NewScheme()
	if e := crd.AddToScheme(schema); e != nil {
		logger.Sugar().Fatalf("failed to add crd schema, reason=%v", e)
	}
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 schema,
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
	e := ctrl.NewWebhookManagedBy(mgr).
		For(&crd.Mybook{}).
		WithDefaulter(r).
		WithValidator(r).
		Complete()
	if e != nil {
		logger.Sugar().Fatalf("failed to NewWebhookManagedBy, reason=%v", e)
	}

}
