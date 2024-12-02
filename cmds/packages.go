package cmds

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/1peripheral/pyvenv/utils"
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
      path, err := utils.GetPath(name)
      if err != nil {
        fmt.Println(err.Error())
        return
      }

      switch runtime.GOOS {
      case "linux": {
        path = filepath.Join(path, "/bin")
      }
      case "windows": {
        path = filepath.Join(path, "/Scripts")
        fmt.Println(path)
      }
      default: {}
      }

      output, err := utils.ExecuteCmd(path + "/pip list")
      if err != nil {
        fmt.Println(err.Error())
        return
      }
      fmt.Println(output)
    },
  }

  return cmd
}
