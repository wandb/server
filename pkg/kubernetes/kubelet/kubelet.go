package kubelet

import (
	"fmt"
	"os"
	"runtime"

	_ "embed"

	"github.com/wandb/server/pkg/download"
	"github.com/wandb/server/pkg/files"
	"github.com/wandb/server/pkg/linux/systemd"
)

func DownloadURL(version string) string {
	return fmt.Sprintf(
		"https://dl.k8s.io/release/v%s/bin/linux/%s/kubelet",
		version,
		runtime.GOARCH,
	)
}

func Download(version string, path string) error {
	return download.HTTPDownloadAndSave(DownloadURL(version), path)
}

//go:embed 10-kube.conf
var service string
func Install(binary string) {
	files.CopyFile(binary, "/usr/local/kubelet")
	os.Chmod("/usr/local/kubelete", 0755)

	os.MkdirAll("/etc/systemd/system/kubelet.service.d", 0755)
	os.WriteFile(
		"/etc/systemd/system/kubelet.service.d/10-kube.conf",
		[]byte(service),
		0600,
	)

	systemd.ReloadDemon()
	systemd.Enable("kubelete")
	systemd.Restart("kubelete")
}