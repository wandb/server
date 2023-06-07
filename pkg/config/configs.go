package config

import (
	"github.com/spf13/viper"
	"github.com/wandb/server/pkg/dependency"
	"github.com/wandb/server/pkg/kubernetes/addons/flannel"
	"github.com/wandb/server/pkg/kubernetes/cni"
	"github.com/wandb/server/pkg/kubernetes/conntrack"
	"github.com/wandb/server/pkg/kubernetes/containerd"
	"github.com/wandb/server/pkg/kubernetes/crictl"
	"github.com/wandb/server/pkg/kubernetes/helm"
	"github.com/wandb/server/pkg/kubernetes/kubeadm"
	"github.com/wandb/server/pkg/kubernetes/kubectl"
	"github.com/wandb/server/pkg/kubernetes/kubelet"
	"github.com/wandb/server/pkg/kubernetes/runc"
)

type config struct {
	Airgap bool

	Dir string

	Kubernetes struct {
		Version string
		Images  []string
	}

	CNI        string
	Containerd string
	Crictl     string
	Flannel    string
	Helm       string
	Conntract  string
	Runc       string
}

var Config config

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	viper.SetEnvPrefix("INSTALLER")
	viper.AutomaticEnv()

	viper.AddConfigPath(".")
	viper.ReadInConfig()

	viper.SetDefault("dir", ".")
	viper.SetDefault("airgap", false)

	viper.Unmarshal(&Config)
}

func KubernetesPackages() []dependency.Package {
	return []dependency.Package{
		kubeadm.NewPackage(Config.Kubernetes.Version, Config.Dir),
		kubectl.NewPackage(Config.Kubernetes.Version, Config.Dir),
		kubelet.NewPackage(Config.Kubernetes.Version, Config.Dir),

		cni.NewPackage(Config.CNI, Config.Dir),
		containerd.NewPackage(Config.Containerd, Config.Dir),

		crictl.NewPackage(Config.Crictl, Config.Dir),
		runc.NewPackage(Config.Runc, Config.Dir),
		conntrack.NewPackage(Config.Conntract, Config.Dir),
		flannel.NewPackage(Config.Flannel, Config.Dir),
		helm.NewPackage(Config.Helm, Config.Dir),
	}
}

func KubernetesAddonPackages() []dependency.Package {
	return []dependency.Package{
		flannel.NewPackage(Config.Flannel, Config.Dir),
	}
}
