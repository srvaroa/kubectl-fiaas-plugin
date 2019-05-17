package main

import (
	"github.com/srvaroa/kubectl-fiaas-plugin/pkg/cmd/logs"
)

func main() {
	logs.NewCommand().Execute()
}
