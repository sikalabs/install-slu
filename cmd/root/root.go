
package root

import (
	"github.com/sikalabs/install-slu/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "install-slu",
	Short: "install-slu, " + version.Version,
}
