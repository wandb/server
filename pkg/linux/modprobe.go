package linux

import "os/exec"

func Modprobe(mod string) error {
	return exec.Command("modprobe", mod).Run()
}