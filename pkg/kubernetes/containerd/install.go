package containerd

import (
	"context"
	"fmt"
	"os"
	"runtime"

	_ "embed"

	"github.com/containerd/containerd"
	"github.com/pterm/pterm"
	"github.com/wandb/server/pkg/download"
	"github.com/wandb/server/pkg/files"
	"github.com/wandb/server/pkg/linux/systemd"
)

const GithubRepo = "https://github.com/containerd/containerd"

//go:embed containerd.service
var service string

func DownloadURL(version string) string {
	return fmt.Sprintf(
		"%s/releases/download/v%s/containerd-%s-linux-%s.tar.gz",
		GithubRepo,
		version,
		version,
		runtime.GOARCH,
	)
}

func Download(version string, path string) error {
	return download.HTTPDownloadAndSave(DownloadURL(version), path)
}

func unzip(tarFile string) error {
	return files.ExtractTarGz(tarFile, "/usr/local")
}

func setupSystemd() error {
	err := os.WriteFile(
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

func Install(tarFile string) {
	err := unzip(tarFile)
	pterm.Fatal.PrintOnError(err, "failed to unzip containerd tar file")
	
	err = setupSystemd()
	pterm.Fatal.PrintOnError(err, "failed to setup containerd systemd service")
}

func IsInstalled() bool {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return false
	}
	if err = client.Close(); err != nil {
		return false
	}
	return true
}

func Client() (*containerd.Client) {
	client, err := containerd.New("/run/containerd/containerd.sock")
	pterm.Fatal.PrintOnError(err, "failed to create containerd client")
	return client
}

func Version() string {
	client := Client()
	defer client.Close()

	info, err := client.Version(context.Background())
	pterm.Fatal.PrintOnError(err, "failed to get containerd version")

	return info.Version
}