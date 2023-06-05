package kubectl

import (
	"fmt"
	"os"
	"runtime"

	"github.com/wandb/server/pkg/download"
	"github.com/wandb/server/pkg/files"
)

func DownloadURL(version string) string {
	return fmt.Sprintf(
		"https://dl.k8s.io/release/v%s/bin/linux/%s/kubectl",
		version,
		runtime.GOARCH,
	)
}

func Download(version string, path string) error {
	return download.HTTPDownloadAndSave(DownloadURL(version), path)
}

func Install(binary string) {
	files.CopyFile(binary, "/usr/local/kubectl")
	os.Chmod("/usr/local/kubeadm", 0755)
}