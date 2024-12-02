package cmds

import (
  "fmt"

	"github.com/1peripheral/pyvenv/utils"
	"github.com/spf13/cobra"
)

func listPackagesCmd() *cobra.Command {
  var cmd  = &cobra.Command{
    Use: "packages [name]",
    Short: "lists the installed packages on a virtual env",
    Run: func(cmd *cobra.Command, args []string) {
      name := args[0]
      path, err := utils.GetPath(name)
      if err != nil {
        fmt.Println(err.Error())
        return
      }
      utils.ExecuteCmd("ls " + path + "/site-packages")
    },
  }

  return cmd
}
