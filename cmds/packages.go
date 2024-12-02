package cmds

import (
	"fmt"

	"github.com/1peripheral/pyvm/utils"
	"github.com/spf13/cobra"
)

func listPackagesCmd() *cobra.Command {
  var cmd  = &cobra.Command{
    Use: "packages [name]",
    Short: "lists the installed packages on a virtual env",
    Run: func(cmd *cobra.Command, args []string) {
      if len(args) < 1 {
        cmd.Usage()
        return
      }

      name := args[0]
      err := utils.ListPackages(name)
      if err != nil {
        fmt.Println(err.Error())
        return
      }
    },
  }

  return cmd
}
