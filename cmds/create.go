package cmds

import (
	"fmt"
	// "os/exec"

	"github.com/1peripheral/pyvenv/services"
	"github.com/spf13/cobra"
)


func createEnvCmd() *cobra.Command {
  var cmd = &cobra.Command{
    Use: "create [name] [path]",
    Short: "Create a new python virtual environment",
    Args: cobra.RangeArgs(1, 2),
    PreRun: func(cmd *cobra.Command, args []string) {
      // TODO: load the file where the envs are stored
    },
    Run: func(cmd *cobra.Command, args []string) {
      name := args[0] 
      path := args[1] 

      err := services.AddEnv(name, path)
      if err != nil {
        fmt.Println(err.Error())
        return
      }

      fmt.Printf("Creating a new virtual env name : %s\n", name)
      // Creating python virtual env
      // console_cmd := exec.Command(pyCmd, "-m", "venv", path)
      // err := console_cmd.Run()
      // if err != nil {
      //   fmt.Println("Error :\n", err)
      // }
    },
  }

  cmd.SetUsageTemplate("Usage : create <path>\n")

  return cmd
}
