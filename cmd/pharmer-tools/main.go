package main

import (
	"os"

	"github.com/pharmer/cloud/pkg/cmds"
	_ "github.com/pharmer/cloud/pkg/credential/cloud"
	"kmodules.xyz/client-go/logs"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmds.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
