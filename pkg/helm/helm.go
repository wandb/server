package helm

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

const (
	secretsStorageDriver = "secrets"
)

var (
	noopLogger = func(_ string, _ ...interface{}) {}
)

func InitConfig(namespace string) (*cli.EnvSettings, *action.Configuration, error) {
	settings := cli.New()
	settings.SetNamespace(namespace)
	config := new(action.Configuration)
	err := config.Init(
		settings.RESTClientGetter(),
		settings.Namespace(),
		secretsStorageDriver,
		noopLogger,
	)
	return settings, config, err
}
