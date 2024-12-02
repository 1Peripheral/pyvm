package cmds

import (
	"fmt"

	"github.com/1peripheral/pyvm/utils"
	"github.com/spf13/cobra"
)

func deleteEnv() *cobra.Command {
  var cmd  = &cobra.Command{
    Use: "delete [name]",
    Short: "lists the available python environments",
    Run: func(cmd *cobra.Command, args []string) {
      if len(args) < 1 {
        cmd.Usage()
        return
      }
      name := args[0]
      err := utils.DeleteEnv(name)
      if err != nil {
        fmt.Println(err.Error())
        return
      }
      fmt.Println("Environment " + name + " has been delete")
    },
  }

  cmd.SetUsageTemplate("Usage : pyvenv delete [name]\n")

  return cmd
}
