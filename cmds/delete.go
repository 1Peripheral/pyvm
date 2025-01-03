package cmds

import (
	"fmt"

	"github.com/1peripheral/pyvm/utils"
	"github.com/spf13/cobra"
)

func deleteEnv() *cobra.Command {
  var cmd  = &cobra.Command{
    Use: "rm [name]",
    Short: "Delete a virtual environment",
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
      fmt.Println("Environment " + name + " has been deleted")
    },
  }

  cmd.SetUsageTemplate("Usage : pyvm rm [name]\n")

  return cmd
}
