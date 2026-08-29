package swap

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

func SwapIsOn() bool {
	cmd := exec.Command("swapon", "--summary")
	output, _ := cmd.Output()
	return strings.Contains(string(output), " ")
}

func SwapIsEnabled() bool {
	return swapFstabEnabled()
}

func swapFstabEnabled() bool {
	f, _ := os.Open("/etc/fstab")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") && strings.Contains(line, "swap") {
			return true
		}
	}
	return false
}

func swapFstabDisable() {
	exec.Command("sed", "--in-place=.bak", "/\\bswap\\b/ s/^/#/", "/etc/fstab").Run()
}

func MustSweepoff() {
	if SwapIsOn() || SwapIsEnabled() {
		pterm.Warning.Println("This application is incompatible with memory swapping enabled.")
		disable, _ := pterm.DefaultInteractiveConfirm.WithDefaultValue(true).Show()
		if disable {
			exec.Command("swapoff", "--all").Run()

			if swapFstabEnabled() {
				swapFstabDisable()
				pterm.Info.Println("/etc/fstab has been updated to disable swap")
				pterm.Warning.Println(
					"Changes have been made to /etc/fstab." + 
					"We recommend reviewing them after completing this installation to " +
					"ensure mounts are correctly configured.")
				time.Sleep(5 * time.Second)
			}
		} else {
			os.Exit(1)
		}
	}
}