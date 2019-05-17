package main

import (
	"github.com/srvaroa/kubectl-fiaas-plugin/pkg/cmd/pods"
)

func main() {
	pods.NewCommand().Execute()
}
