package cmds

import (
	"fmt"
	"os/exec"

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
      
      // Creating python virtual env
      console_cmd := exec.Command("echo", "Hello, There Partner")
      output, err := console_cmd.Output()
      if err != nil {
        fmt.Println("Error :\n", err)
      }
      fmt.Println("Environment Created:\n", string(output))
    },
  }

  cmd.SetUsageTemplate("Usage : pyvenv create <name> <opt:path>\n")

  return cmd
}
