package kubeadm

import (
	"fmt"
	"os"
	"runtime"

	"github.com/wandb/server/pkg/download"
	"github.com/wandb/server/pkg/files"
)

func DownloadURL(version string) string {
	return fmt.Sprintf(
		"https://dl.k8s.io/release/v%s/bin/linux/%s/kubeadm",
		version,
		runtime.GOARCH,
	)
}

func Download(version string, path string) error {
	return download.HTTPDownloadAndSave(DownloadURL(version), path)
}

func Install(binary string) {
	files.CopyFile(binary, "/usr/local/kubeadm")
	os.Chmod("/usr/local/kubeadm", 0755)

	LoadModules()
	LoadSystemdModules()
}