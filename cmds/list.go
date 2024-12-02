package cmds

import (
	"github.com/1peripheral/pyvenv/utils"
	"github.com/spf13/cobra"
)



func listEnvCmd() *cobra.Command {
  var cmd  = &cobra.Command{
    Use: "list",
    Short: "lists the available python environments",
    Run: func(cmd *cobra.Command, args []string) {
      utils.PrintEnvs()
    },
  }

  return cmd
}
