package install

import (
	"path"
	"runtime"

	"github.com/sikalabs/install-slu/cmd/root"
	"github.com/sikalabs/install-slu/install"
	"github.com/spf13/cobra"
)

var CmdFlagBinDir string
var CmdFlagBinName string
var CmdFlagSluVersion string
var CmdFlagOS string
var CmdFlagArch string

var Cmd = &cobra.Command{
	Use:     "install",
	Short:   "install slu",
	Aliases: []string{"i"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if CmdFlagSluVersion == "latest" {
			CmdFlagSluVersion = install.GetLatestRelease()
		}
		install.Install(
			CmdFlagSluVersion,
			CmdFlagOS,
			CmdFlagArch,
			path.Join(CmdFlagBinDir, CmdFlagBinName),
		)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().StringVarP(
		&CmdFlagSluVersion,
		"version",
		"v",
		"latest",
		"slu version",
	)
	Cmd.PersistentFlags().StringVarP(
		&CmdFlagBinDir,
		"bin-dir",
		"d",
		"/usr/local/bin",
		"Binary dir",
	)
	Cmd.PersistentFlags().StringVarP(
		&CmdFlagBinName,
		"bin-name",
		"n",
		"slu",
		"Name of the binary",
	)
	Cmd.PersistentFlags().StringVarP(
		&CmdFlagOS,
		"os",
		"o",
		runtime.GOOS,
		"OS",
	)
	Cmd.PersistentFlags().StringVarP(
		&CmdFlagArch,
		"arch",
		"a",
		runtime.GOARCH,
		"Architecture",
	)
}
