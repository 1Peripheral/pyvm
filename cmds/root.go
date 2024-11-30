package cmds

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/1peripheral/pyvenv/services"
	"github.com/spf13/cobra"
)

var pyCmd = "python3"

func checkPython() bool {
  commands := []string{"python", "python3", "py"}

  for _, cmdName := range commands {
    cmd := exec.Command(cmdName, "--version")
    err := cmd.Run()
    if err == nil {
      pyCmd = cmdName
      return true
    }
  }

  fmt.Println("Python is not installed.")
  return false
}

var rootcmd = &cobra.Command{
  Use: "pyvenv",
  PersistentPreRun: func(cmd *cobra.Command, args []string) {
    // Checking if python exists
    if checkPython() == false {
      os.Exit(1)
    } 
    services.InitEnv()
  },
  PersistentPostRun: func(cmd *cobra.Command, args []string) {
    services.SaveChanges()
  },
  Run: func(cmd *cobra.Command, args []string) {
    cmd.Usage();
  },
}

func Execute() {
  rootcmd.AddCommand(createEnvCmd())
  rootcmd.Execute();
}
