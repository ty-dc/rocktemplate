package types

type MybookManager interface {
	RunWebhookServer(webhookPort int, tlsDir string)
	RunInformer(leaseName, leaseNameSpace string)
}
