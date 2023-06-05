package cni

import (
	"fmt"
	"os"
	"runtime"

	"github.com/pterm/pterm"
	"github.com/wandb/server/pkg/download"
	"github.com/wandb/server/pkg/files"
)

const GithubRepo = "https://github.com/containernetworking/plugins"

func DownloadURL(version string) string {
	return fmt.Sprintf(
		"%s/releases/download/v%s/cni-plugins-linux-%s-v%s.tgz",
		GithubRepo,
		version,
		runtime.GOARCH,
		version,
	)
}

func Download(version string, path string) error {
	return download.HTTPDownloadAndSave(DownloadURL(version), path)
}

func Install(tarFile string) {
	pterm.Info.Printf("Installing cni plugins from %s\n", tarFile)
	os.MkdirAll("/opt/cni/bin", 0755)
	files.ExtractTarGz(tarFile, "/opt/cni/bin")
}