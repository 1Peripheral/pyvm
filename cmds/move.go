package cmds

import (
	"fmt"

	"github.com/1peripheral/pyvm/utils"
	"github.com/spf13/cobra"
)

func moveCmd() *cobra.Command {
  var cmd  = &cobra.Command{
    Use: "mv [name] [new_path]",
    Short: "Move a virtual env",
    Run: func(cmd *cobra.Command, args []string) {
      if len(args) < 2 {
        cmd.Usage()
        return
      }
      name := args[0]
      newPath := args[1]
      err := utils.MoveEnv(name, newPath)
      if err != nil {
        fmt.Println(err.Error())
      }
      fmt.Printf("%s moved to %s\n", name, newPath);
    },
  }

  cmd.SetUsageTemplate("Usage : pyvm mv [name] [new_path]\n")

  return cmd
}
