package crictl

import (
	"fmt"
	"runtime"

	"github.com/wandb/server/pkg/download"
	"github.com/wandb/server/pkg/files"
)

const GithubRepo = "https://github.com/kubernetes-sigs/cri-tools"

func DownloadURL(version string) string {
	return fmt.Sprintf(
		"%s/releases/download/v%s/crictl-v%s-linux-%s.tar.gz",
		GithubRepo,
		version,
		version,
		runtime.GOARCH,
	)
}

func Download(version string, path string) error {
	return download.HTTPDownloadAndSave(DownloadURL(version), path)
}

func Install(tarFile string) {
	files.ExtractTarGz(tarFile, "/usr/local/bin")
}