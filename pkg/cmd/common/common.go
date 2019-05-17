package common

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Execute(cmd *exec.Cmd) {

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err == nil {
		return
	}

	if exiterr, ok := err.(*exec.ExitError); ok {
		// This works on both Unix and Windows. Although package
		// syscall is generally platform dependent, WaitStatus is
		// defined for both Unix and Windows and in both cases has
		// an ExitStatus() method with the same signature.
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			os.Exit(status.ExitStatus())
		}
	} else {
		fmt.Printf("Failed to run kubectl: %v\n", err)
		return
	}
}

func CollectFlags(cmd *cobra.Command) []string {
	flags := make(map[string]string)
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Changed {
			flags[flag.Name], _ = cmd.Flags().GetString(flag.Name)
		}
	})

	var sCmd = []string{}
	for name, value := range flags {
		sCmd = append(sCmd, fmt.Sprintf("--%s=%s", name, value))
	}
	return sCmd
}
