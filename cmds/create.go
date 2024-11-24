package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)


func createEnvCmd() *cobra.Command {
  var cmd = &cobra.Command{
    Use: "create [name] [path]",
    Short: "Create a new python virtual environment",
    Args: cobra.RangeArgs(1, 2),
    Run: func(cmd *cobra.Command, args []string) {
      name := args[0] 
      path := "."
      if len(args) > 1 {
        path = args[1]
      }
      fmt.Printf("Creating a new virtual env name : %s at %s\n", name, path)
    },
  }

  cmd.SetUsageTemplate("Usage : pyvenv create <name> <opt:path>\n")

  return cmd
}
