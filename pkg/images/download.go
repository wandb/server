package images

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/wandb/server/pkg/dependency"
	"github.com/wandb/server/pkg/kubernetes/containerd"
	"github.com/wandb/server/pkg/kubernetes/docker"
)

// Download downloads an image and saves it to a file.
func Download(image string, filename string) error {
	_, err := exec.LookPath("docker")
	if err == nil {
		return docker.DownloadImage(image, filename)
	}

	if containerd.IsInstalled() {
		return containerd.DownloadImage(image, filename)
	}

	return fmt.Errorf("no supported container runtime found")
}

func NewPackage(image string, dest string) dependency.Package {
	return &Package{image, dest}
}

type Package struct {
	image string
	dest string
}

func (p Package) Download() error {
	return Download(p.image, p.imageFile())
}

func (p Package) imageFile() string {
	img := strings.Split(p.image, ":")[0]
	v := p.Version()
	file := path.Join(p.dest, img, v+".tar")
	os.MkdirAll(filepath.Dir(file), 0755)
	return file
}

// Install implements dependency.Package
func (*Package) Install() error {
	panic("unimplemented")
}

// Name implements dependency.Package
func (p Package) Name() string {
	return strings.Split(p.image, ":")[0]
}

// Version implements dependency.Package
func (p Package) Version() string {
	image := strings.Split(p.image, ":")[1]
	return strings.Trim(image, "v")
}
