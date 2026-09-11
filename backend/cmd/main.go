package main

import (
	"fmt"
	"os"
	_ "time/tzdata"

	"github.com/AikidoSec/firewall-go/zen"
	"github.com/pocket-id/pocket-id/backend/internal/cmds"
	"github.com/pocket-id/pocket-id/backend/internal/common"
)

// @title Pocket ID API
// @version 1.0
// @description.markdown

func main() {
	if err := zen.Protect(); err != nil {
		fmt.Fprintf(os.Stderr, "aikido zen error: %v\n", err)
		os.Exit(1)
	}

	if err := common.ValidateEnvConfig(&common.EnvConfig); err != nil {
		fmt.Fprintf(os.Stderr, "config error: %v\n", err)
		os.Exit(1)
	}
	cmds.Execute()
}
