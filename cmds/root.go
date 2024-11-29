package cmds

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func checkPython() bool {
  commands := []string{"python", "python3", "py"}

  for _, cmdName := range commands {
    cmd := exec.Command(cmdName, "--version")
    err := cmd.Run()
    if err == nil {
      return true
    }
  }

  fmt.Println("Python is not installed.")
  return false
}

var rootcmd = &cobra.Command{
  Use: "pyvenv",
  Run: func(cmd *cobra.Command, args []string) {
    // Checking if python exists
    if checkPython() == false {
      os.Exit(1)
    } 

    cmd.Usage();
  },
}

func Execute() {
  rootcmd.AddCommand(createEnvCmd())
  rootcmd.Execute();
}
