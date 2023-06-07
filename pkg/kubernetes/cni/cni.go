package cni

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/wandb/server/pkg/dependency"
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

func NewPackage(version string, dest string) dependency.Package {
	return &CNIPluginPackage{version, dest}
}

type CNIPluginPackage struct {
	version string
	dest    string
}

func (p CNIPluginPackage) Version() string {
	return p.version
}

func (p CNIPluginPackage) path() string {
	pa := path.Join(p.dest, p.Name(), p.version)
	os.MkdirAll(pa, 0755)
	return pa
}

func (p CNIPluginPackage) Install() error {
	err := os.MkdirAll("/opt/cni/bin", 0755)
	if err != nil {
		return err
	}
	tar := path.Join(p.path(), "cni-plugins.tgz")
	return exec.Command("tar", "-xzf", tar, "-C", "/opt/cni/bin").Run()
}

func (p CNIPluginPackage) Download() error {
	return dependency.HTTPDownloadAndSave(
		DownloadURL(p.version),
		path.Join(p.path(), "cni-plugins.tgz"),
	)
}

func (p CNIPluginPackage) Name() string {
	return "cni-plugins"
}
