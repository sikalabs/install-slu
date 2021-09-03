package cmd

import (
	"github.com/sikalabs/install-slu/cmd/root"
	_ "github.com/sikalabs/install-slu/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
