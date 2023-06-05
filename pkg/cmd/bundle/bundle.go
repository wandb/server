package bundle

import (
	"sync"

	"github.com/pterm/pterm"
	"github.com/wandb/server/pkg/kubernetes/cni"
	"github.com/wandb/server/pkg/kubernetes/containerd"
	"github.com/wandb/server/pkg/kubernetes/crictl"
	"github.com/wandb/server/pkg/kubernetes/kubeadm"
	"github.com/wandb/server/pkg/kubernetes/kubectl"
	"github.com/wandb/server/pkg/kubernetes/kubelet"
	"github.com/wandb/server/pkg/kubernetes/runc"
)

func DownloadPackages() {
	packages := []func()error{
		func() error { return containerd.Download("1.7.1", "./packages/containerd.tar.gz") },
		func() error { return crictl.Download("1.27.0", "./packages/crictl.tar.gz") },
		func() error { return runc.Download("1.1.7", "./packages/runc") },
		func() error { return cni.Download("1.3.0", "./packages/cni-plugins.tar.gz") },
		func() error { return kubeadm.Download("1.27.2", "./packages/kubeadm") },
		func() error { return kubelet.Download("1.27.2", "./packages/kubelet") },
		func() error { return kubectl.Download("1.27.2", "./packages/kubectl") },
	}

	progressbar, _ := pterm.DefaultProgressbar.
		WithTotal(len(packages)).
		WithTitle("Downloading packages").
		Start()
	
	wg := sync.WaitGroup{}
	wg.Add(len(packages))
	for _, pkg := range packages {
		go func(download func()error) {
			err := download()
			progressbar.Increment()
			wg.Done()
			pterm.Error.PrintOnError(err)
		}(pkg)
	}
	wg.Wait()

	progressbar.Stop()
}
