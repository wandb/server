package systemd

import "os/exec"

func ReloadDemon() error {
	return exec.Command("systemctl", "daemon-reload").Run()
}

func Enable(service string) error {
	return exec.Command("systemctl", "enable", service).Run()
}

func Restart(service string) error {
	return exec.Command("systemctl", "restart", service).Run()
}