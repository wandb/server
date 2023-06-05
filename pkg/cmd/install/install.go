package isntall

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

func Install() {
	packages := []func(){
		func() { containerd.Install("./packages/containerd.tar.gz") },
		func() { crictl.Install("./packages/crictl.tar.gz") },
		func() { runc.Install("./packages/runc") },
		func() { cni.Install("./packages/cni-plugins.tar.gz") },
		func() { kubeadm.Install("./packages/kubeadm") },
		func() { kubelet.Install("./packages/kubelet") },
		func() { kubectl.Install("./packages/kubectl") },
	}

	progressbar, _ := pterm.DefaultProgressbar.
		WithTotal(len(packages)).
		WithTitle("Downloading packages").
		Start()
	
	wg := sync.WaitGroup{}
	wg.Add(len(packages))
	for _, pkg := range packages {
		go func(download func()) {
			download()
			progressbar.Increment()
			wg.Done()
		}(pkg)
	}
	wg.Wait()

	progressbar.Stop()
}