package cmds

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/1peripheral/pyvm/utils"
	"github.com/spf13/cobra"
)

func addEnv() *cobra.Command {
  var cmd  = &cobra.Command{
    Use: "add [name] [path]",
    Short: "Add an existing virtual environment to the manager",
    Run: func(cmd *cobra.Command, args []string) {
      if len(args) < 2 {
        cmd.Usage()
        return
      }

      name := args[0]
      path, err := filepath.Abs(args[1])
      if err != nil {
        fmt.Println("Error: Incorrect path")
      }

      //TODO: check if the path is a valid python virtual env
      info, err := os.Stat(path)
      if err != nil {
        if os.IsNotExist(err) {
          fmt.Println("Error: The path you provided does not exist .")
          return
        }
        fmt.Println("Error: An error occured when adding " + path)
        return
      }
      if !info.IsDir() {
        fmt.Println("Error: " + path + " is not a directory")
        return
      }
    
      utils.AddEnv(name, path)
      fmt.Printf("Virtual Environment '%s' has been added \n", name)
    },
  }

  cmd.SetUsageTemplate("Usage : pyvm add [name] [path]\n")

  return cmd
}
