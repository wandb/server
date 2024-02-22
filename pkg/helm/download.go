package helm

import (
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/releaseutil"
	"helm.sh/helm/v3/pkg/repo"
	"k8s.io/kubectl/pkg/scheme"

	"k8s.io/apimachinery/pkg/runtime"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	v1 "k8s.io/api/core/v1"
)

func DownloadChart(url string, name string, version string) (string, error) {
	entry := new(repo.Entry)
	entry.URL = url
	entry.Name = name

	file := repo.NewFile()
	file.Update(entry)

	settings := cli.New()
	providers := getter.All(settings)
	chartRepo, err := repo.NewChartRepository(entry, providers)
	if err != nil {
		return "", err
	}

	_, err = chartRepo.DownloadIndexFile()
	if err != nil {
		return "", err
	}

	chartURL, err := repo.FindChartInRepoURL(
		entry.URL, entry.Name, version,
		"", "", "",
		providers,
	)
	if err != nil {
		return "", err
	}

	_, cfg, err := InitConfig("")
	if err != nil {
		return "", err
	}

	client := downloader.ChartDownloader{
		Verify:  downloader.VerifyNever,
		Getters: getter.All(settings),
		Options: []getter.Option{
			// TODO: Add support for other auth methods
		},
		RegistryClient:   cfg.RegistryClient,
		RepositoryConfig: settings.RepositoryConfig,
		RepositoryCache:  settings.RepositoryCache,
	}

	dest := "./bundle/chart"
	os.MkdirAll(dest, 0755)

	saved, _, err := client.DownloadTo(chartURL, version, dest)
	if err != nil {
		return "", err
	}

	return saved, err
}

const DefaultReleaseName = "wandb"

func GetRuntimeObjects(chartPath string, vals map[string]interface{}) ([]runtime.Object, error) {
	_, c, _ := InitConfig("")

    chart, err := loader.Load(chartPath)
	if err != nil {
		return nil, err
	}

	install := action.NewInstall(c)
	install.ReleaseName = DefaultReleaseName
	install.DryRun = true
	install.Replace = true
	install.ClientOnly = true

	release, err := install.Run(chart, vals)
	if err != nil {
		return nil, err
	}

	manifests := releaseutil.SplitManifests(release.Manifest)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	runtimeObejcts := []runtime.Object{}
	for _, yaml := range manifests {
		obj, _, err := decode([]byte(yaml), nil, nil)
		if err != nil {
			return nil, err
		}

		runtimeObejcts = append(runtimeObejcts, obj)
	}

	return runtimeObejcts, nil
}

func ExtractImages(obj []runtime.Object) []string {
	images := []string{}
	for _, o := range obj {
		images = append(images, ExtractImage(o)...)
	}
	return images
}

func ExtractImage(obj runtime.Object) []string {
	var images []string
	switch typedObj := obj.(type) {
    case *v1.Pod:
        for _, container := range typedObj.Spec.Containers {
            images = append(images, container.Image)
        }
	case *v1.ReplicationController:
        for _, container := range typedObj.Spec.Template.Spec.Containers {
            images = append(images, container.Image)
        }
    case *appsv1.ReplicaSet:
        for _, container := range typedObj.Spec.Template.Spec.Containers {
            images = append(images, container.Image)
        }
	case *appsv1.Deployment:
        for _, container := range typedObj.Spec.Template.Spec.Containers {
            images = append(images, container.Image)
        }
    case *appsv1.StatefulSet:
        for _, container := range typedObj.Spec.Template.Spec.Containers {
            images = append(images, container.Image)
        }
    case *appsv1.DaemonSet:
        for _, container := range typedObj.Spec.Template.Spec.Containers {
            images = append(images, container.Image)
        }
    case *batchv1.Job:
        for _, container := range typedObj.Spec.Template.Spec.Containers {
            images = append(images, container.Image)
        }
    case *batchv1beta1.CronJob:
        for _, container := range typedObj.Spec.JobTemplate.Spec.Template.Spec.Containers {
            images = append(images, container.Image)
        }
	}
	return images
}