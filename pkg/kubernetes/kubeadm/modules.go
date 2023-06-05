package kubeadm

import (
	"os"
	"os/exec"

	_ "embed"

	"github.com/wandb/server/pkg/linux"
)

//go:embed k8s.conf
var modulesConfig string

func LoadModules() {
	linux.Modprobe("overlay")
	linux.Modprobe("br_netfilter")

	linux.Modprobe("ip_tables")
	linux.Modprobe("ip6_tables")

	linux.Modprobe("ip_vs")
	linux.Modprobe("ip_vs_rr")
	linux.Modprobe("ip_vs_wrr")
	linux.Modprobe("ip_vs_sh")

	linux.Modprobe("nf_conntrack")

	os.MkdirAll("/etc/modules-load.d", 0755)
	os.WriteFile(
		"/etc/modules-load.d/k8s.conf",
		[]byte(modulesConfig),
		0600,
	)
}

//go:embed ip.conf
var systemdConf string

func LoadSystemdModules() {
	os.MkdirAll("/etc/sysctl.d", 0755)
	os.WriteFile(
		"/etc/sysctl -a.d/k8s-ipv4.conf",
		[]byte(systemdConf),
		0600,
	)
	reloadSysctl()
}

func reloadSysctl() error {
	return exec.Command("sysctl", "--system").Run()
}