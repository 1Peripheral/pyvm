package cmds

import (
	"github.com/1peripheral/pyvm/utils"
	"github.com/spf13/cobra"
)

func listEnvCmd() *cobra.Command {
  var cmd  = &cobra.Command{
    Use: "list",
    Short: "Lists the available python environments",
    Run: func(cmd *cobra.Command, args []string) {
      utils.PrintEnvs()
    },
  }

  return cmd
}
