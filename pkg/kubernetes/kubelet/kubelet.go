package kubelet

import (
	"fmt"
	"os"
	"path"
	"runtime"

	_ "embed"

	"github.com/wandb/server/pkg/dependency"
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

func NewPackage(version string, dest string) dependency.Package {
	return &KubeletPackage{version, dest}
}

type KubeletPackage struct {
	version string
	dest string
}

func (p KubeletPackage) Version() string {
	return p.version
}

func (p KubeletPackage) path() string {
	pa := path.Join(p.dest, p.Name(), p.version)
	os.MkdirAll(pa, 0755)
	return pa
}

//go:embed 10-kube.conf
var kubeletConf string
//go:embed kubelet.service
var kubeletService string

func (p KubeletPackage) Install() error {
	binary := path.Join(p.path(), "kubelet")
	files.CopyFile(binary, "/usr/bin/kubelet")
	os.Chmod("/usr/bin/kubelet", 0755)

	os.MkdirAll("/etc/systemd/system/kubelet.service.d", 0755)
	os.WriteFile(
		"/etc/systemd/system/kubelet.service.d/10-kube.conf",
		[]byte(kubeletConf),
		0600,
	)

	os.WriteFile(
		"/etc/systemd/system/kubelet.service.d/10-kube.conf",
		[]byte(kubeletService),
		0600,
	)
	systemd.ReloadDemon()
	systemd.Enable("kubelet")
	return systemd.Restart("kubelet")
}

func (p KubeletPackage) Download() error {
	return dependency.HTTPDownloadAndSave(
		DownloadURL(p.version),
		path.Join(p.path(), "kubelet"),
	)
}

func (p KubeletPackage) Name() string {
	return "kubelet"
}