package kubeadm

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"time"

	"github.com/pterm/pterm"
	"github.com/wandb/server/pkg/dependency"
	"github.com/wandb/server/pkg/discover/networking"
	"github.com/wandb/server/pkg/files"
)

func DownloadURL(version string) string {
	return fmt.Sprintf(
		"https://dl.k8s.io/release/v%s/bin/linux/%s/kubeadm",
		version,
		runtime.GOARCH,
	)
}

func Init() {
	health, _ := pterm.DefaultSpinner.Start("Waiting for kubeadm to be in healthy state")
	retry := 0
	for !IsHealthy() {
		retry += 1
		if retry > 10 {
			health.Fail("Kubeadm unhealthy after 10 retries. Exiting.")
			os.Exit(1)
			return
		}
		time.Sleep(1 * time.Second)
	}
	health.Success("Kubeadm is healthy")
}

func IsHealthy() bool {
	privateIP, _ := networking.GetPrivateIP()
	url := fmt.Sprintf("https://%s:%s/healthz", privateIP, "6443")

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		Proxy:             nil, // Disable the use of a proxy, similar to --noproxy "*"
		DisableKeepAlives: true,
		// Skip certificate validation, similar to --insecure flag
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Timeout:   time.Second * 30,
		Transport: transport,
	}

	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	// Read and discard the response body as required
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return false
	}

	return true
}

func NewPackage(version string, dest string) *KubeadmPackage {
	return &KubeadmPackage{version, dest}
}

type KubeadmPackage struct {
	version string
	dest    string
}

func (p KubeadmPackage) path() string {
	pa := path.Join(p.dest, p.Name(), p.version)
	os.MkdirAll(pa, 0755)
	return pa
}

func (p KubeadmPackage) Init() error {
	privateIP, _ := networking.GetPrivateIP()
	controlEndpoint := privateIP + ":6443"
	cmd := exec.Command(
		"kubeadm", "init",
		"--control-plane-endpoint="+controlEndpoint,
		"--pod-network-cidr=",
		"--ignore-preflight-errors=all")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (p KubeadmPackage) Version() string {
	return p.version
}

func (p KubeadmPackage) Install() error {
	binary := path.Join(p.path(), "kubeadm")
	err := files.CopyFile(binary, "/usr/local/bin/kubeadm")
	if err != nil {
		return err
	}

	err = os.Chmod("/usr/local/bin/kubeadm", 0755)
	if err != nil {
		return err
	}

	err = LoadModules()
	if err != nil {
		return err
	}

	return LoadSystemdModules()
}

func (p KubeadmPackage) Download() error {
	return dependency.HTTPDownloadAndSave(
		DownloadURL(p.version),
		path.Join(p.path(), "kubeadm"),
	)
}

func (p KubeadmPackage) Name() string {
	return "kubeadm"
}
