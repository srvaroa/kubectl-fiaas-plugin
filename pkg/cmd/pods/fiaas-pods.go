package pods

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/srvaroa/kubectl-fiaas-plugin/pkg/cmd/common"
)

const (
	Desc = `List pods from a FIAAS application`
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kubectl-fiaas-pods",
		Short: "List pods from a FIAAS application",
		Run:   run,
	}
	var n string
	cmd.Flags().StringVarP(&n, "namespace", "n", "",
		"the namespace where your app is deployed")
	return cmd
}

func run(cmd *cobra.Command, args []string) {
	var app string
	if len(args) > 0 {
		app = args[0]
	} else {
		fmt.Println("No app name specified")
		return
	}

	var sCmd = []string{"get", "pod", fmt.Sprintf("-lapp=%s", app)}
	sCmd = append(sCmd, common.CollectFlags(cmd)...)
	common.Execute(exec.Command("kubectl", sCmd...))
}
