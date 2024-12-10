package cmds

import (
	"github.com/1peripheral/pyvm/utils"
	"github.com/spf13/cobra"
)

func activateCmd() *cobra.Command {
  var cmd  = &cobra.Command{
    Use: "activate [name]",
    Short: "Activates a virutal environment by its name",
    Run: func(cmd *cobra.Command, args []string) {
      if len(args) < 1 {
        cmd.Usage()
        return
      }
      name := args[0]
      utils.ActivateEnv(name)
    },
  }

  cmd.SetUsageTemplate("Usage : pyvm activate [name]\n")

  return cmd
}
