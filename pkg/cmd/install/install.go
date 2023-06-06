package install

import (
	"sync"

	"github.com/pterm/pterm"
)

func InstallPackages() {
	packages := []func(){
		// func() { containerd.Install("./packages/containerd.tar.gz") },
		// func() { crictl.Install("./packages/crictl.tar.gz") },
		// func() { runc.Install("./packages/runc") },
		// func() { cni.Install("./packages/cni-plugins.tar.gz") },
		// func() { kubeadm.Install("./packages/kubeadm") },
		// func() { kubelet.Install("./packages/kubelet") },
		// func() { kubectl.Install("./packages/kubectl") },
		// func() { conntrack.Install("./packages/conntrack.deb") },
	}

	progressbar, _ := pterm.DefaultProgressbar.
		WithTotal(len(packages)).
		WithTitle("Installing packages").
		Start()
	
	wg := sync.WaitGroup{}
	wg.Add(len(packages))
	for _, pkg := range packages {
		go func(install func()) {
			install()
			progressbar.Increment()
			wg.Done()
		}(pkg)
	}
	wg.Wait()

	progressbar.Stop()
}