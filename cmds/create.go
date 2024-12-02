package cmds

import (
	"fmt"
	"path/filepath"

	"github.com/1peripheral/pyvm/utils"
	"github.com/spf13/cobra"
)


func createEnvCmd() *cobra.Command {
  var cmd = &cobra.Command{
    Use: "create [name] [path]",
    Short: "Create a new python virtual environment",
    Args: cobra.RangeArgs(1, 2),
    Run: func(cmd *cobra.Command, args []string) {
      // TODO: Prettify the output
      if len(args) < 2 {
        cmd.Usage()
        return
      }

      name := args[0] 
      path := args[1] 

      absolutePath, err := filepath.Abs(path)
      if err != nil {
        fmt.Println("Incorrect path")
      }

      if utils.DoesEnvExist(name) {
        fmt.Println("Environment with that name already exists")
        return
      }

      // Creating python virtual env
      if _, err := utils.ExecuteCmd(pyCmd + " -m venv " + path); err != nil {
        fmt.Println("Encountered an error when creating the virtual env")
        fmt.Println(err.Error())
        return
      }
      err = utils.AddEnv(name, absolutePath)
      if err != nil {
        fmt.Println(err.Error())
        return
      }
      fmt.Printf("Virtual environment '%s' has been created\n", name)
    },
  }

  cmd.SetUsageTemplate("Usage : pyvenv create [name] [path]\n")

  return cmd
}
