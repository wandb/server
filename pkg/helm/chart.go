package helm

import (
	"fmt"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/releaseutil"
)


func isInstalled(config *action.Configuration, releaseName string) bool {
	h, err := config.Releases.History(releaseName)
	if err != nil || len(h) < 1 {
		return false
	}

	releaseutil.Reverse(h, releaseutil.SortByRevision)
	rel := h[0]
	st := rel.Info.Status

	return st != release.StatusUninstalled
}

func Apply(
	namespace string,
	releaseName string,
	chart *chart.Chart,
	values map[string]interface{},
) (*release.Release, error) {
	if err := chartutil.ValidateReleaseName(releaseName); err != nil {
		return nil, fmt.Errorf("release name %q", releaseName)
	}

	_, config, err := InitConfig(namespace)
	if err != nil {
		return nil, err
	}

	if isInstalled(config, releaseName) {
		return upgrade(config, namespace, releaseName, chart, values)
	}
	return install(config, namespace, releaseName, chart, values)
}

func install(
	config *action.Configuration,
	namespace string,
	releaseName string,
	chart *chart.Chart,
	values map[string]interface{},
) (*release.Release, error) {
	client := action.NewInstall(config)
	client.ReleaseName = releaseName
	client.Namespace = namespace
	client.CreateNamespace = true
	return client.Run(chart, values)
}

func upgrade(
	config *action.Configuration,
	namespace string,
	releaseName string,
	chart *chart.Chart,
	values map[string]interface{},
) (*release.Release, error) {
	client := action.NewUpgrade(config)
	client.Namespace = namespace
	return client.Run(releaseName, chart, values)
}