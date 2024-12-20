package containerd

import (
	"context"

	"github.com/containerd/containerd"
	"github.com/pterm/pterm"
)

func Client() *containerd.Client {
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
