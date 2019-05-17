package logs

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/srvaroa/kubectl-fiaas-plugin/pkg/cmd/common"
)

const (
	Desc = `Logs from all pods in FIAAS application`
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kubectl-fiaas-logs",
		Short: "Show logs from a FIAAS application",
		Run:   run,
	}
	var n string
	var s string
	cmd.Flags().StringVarP(&n, "namespace", "n", "",
		"the namespace where your app is deployed")
	cmd.Flags().StringVarP(&s, "since", "s", "",
		"Only return logs newer than a relative duration like 5s, 2m, or 3h.  Defaults to all logs.")
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

	var sCmd = []string{"logs", fmt.Sprintf("-lapp=%s", app)}
	sCmd = append(sCmd, common.CollectFlags(cmd)...)
	common.Execute(exec.Command("kubectl", sCmd...))
}
