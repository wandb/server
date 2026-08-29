package containerd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	_ "embed"

	"github.com/pterm/pterm"
	download "github.com/wandb/server/pkg/dependency"
	"github.com/wandb/server/pkg/linux/systemd"
)

const GithubRepo = "https://github.com/containerd/containerd"

func DownloadURL(version string) string {
	return fmt.Sprintf(
		"%s/releases/download/v%s/containerd-%s-linux-%s.tar.gz",
		GithubRepo,
		version,
		version,
		runtime.GOARCH,
	)
}

// go:embed containerd.service
var service string
func setupSystemd() error {
	pterm.Info.Println("Configuring containerd in systemd")

	err := os.MkdirAll("/etc/systemd/system", 0755)
	if err != nil {
		return err
	}
	err = os.WriteFile(
		"/etc/systemd/system/containerd.service",
		[]byte(service),
		0600,
	)
	if err != nil {
		return err
	}

	err = systemd.ReloadDemon()
	if err != nil {
		return err
	}

	err = systemd.Enable("containerd")
	if err != nil {
		return err
	}

	err = systemd.Restart("containerd")
	if err != nil {
		return err
	}
	return nil
}

func NewPackage(version string, dest string) download.Package {
	return &ContainerdPackage{version, dest}
}

type ContainerdPackage struct{
	version string
	dest string
}

func (p ContainerdPackage) path() string {
	pa := path.Join(p.dest, p.Name(), p.version)
	os.MkdirAll(pa, 0755)
	return pa
}

func (p ContainerdPackage) Version() string {
	return p.version
}

func (p ContainerdPackage) Install() error {
	tar := path.Join(p.path(), "containerd.tar.gz")
	err := exec.Command("tar", "-xzf", tar, "-C", "/usr/local").Run()
	if err != nil {
		return err
	}
	return setupSystemd()
}

func (p ContainerdPackage) Download() error {
	return download.HTTPDownloadAndSave(
		DownloadURL(p.version),
		path.Join(p.path(), "containerd.tar.gz"),
	)
}

func (p *ContainerdPackage) Name() string {
	return "containerd"
}