package helm

import (
	"os"

	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
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

	dest := "./charts"
	os.MkdirAll(dest, 0755)
	saved, _, err := client.DownloadTo(chartURL, version, dest)
	if err != nil {
		return "", err
	}
	return saved, err
}
