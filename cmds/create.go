package cmds

import (
	"fmt"
	"path/filepath"

	"github.com/1peripheral/pyvenv/utils"
	"github.com/spf13/cobra"
)


func createEnvCmd() *cobra.Command {
  var cmd = &cobra.Command{
    Use: "create [name] [path]",
    Short: "Create a new python virtual environment",
    Args: cobra.RangeArgs(1, 2),
    Run: func(cmd *cobra.Command, args []string) {
      name := args[0] 
      path := args[1] 

      absolutePath, err := filepath.Abs(path)
      if err != nil {
        fmt.Println("Incorrect path")
      }

      // err = utils.AddEnv(name, absolutePath)
      if utils.DoesEnvExist(name) {
        fmt.Println("Environemtn with that name already exists")
        return
      }

      fmt.Printf("Creating a new virtual env name : %s\n", name)
      // Creating python virtual env
      if err := utils.ExecuteCmd(pyCmd + " -m venv " + path); err != nil {
        fmt.Println("Encountered an error when creating the virtual env")
        fmt.Println(err.Error())
        return
      }
      err = utils.AddEnv(name, absolutePath)
      if err != nil {
        fmt.Println(err.Error())
        return
      }
    },
  }

  cmd.SetUsageTemplate("Usage : create <path>\n")

  return cmd
}
