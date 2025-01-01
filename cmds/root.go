package cmds

import (
	"fmt"
	"os"

	"github.com/1peripheral/pyvm/utils"
	"github.com/spf13/cobra"
)

var pyCmd = ""

func checkPython() bool {
  var err error
  commands := []string{"python", "python3", "py"}
  for _, cmdName := range commands {
    _, err = utils.ExecuteCmd(cmdName + " --version")
    if err == nil {
      pyCmd = cmdName
      return true
    }
  }

  fmt.Println("Python is not installed.")
  return false
}

var rootcmd = &cobra.Command{
  Use: "pyvm",
  PersistentPreRun: func(cmd *cobra.Command, args []string) {
    // Checking if python exists
    if checkPython() == false {
      os.Exit(1)
    } 
    utils.InitEnv()
  },
  PersistentPostRun: func(cmd *cobra.Command, args []string) {
    utils.SaveChanges()
  },
  Run: func(cmd *cobra.Command, args []string) {
    cmd.Usage();
  },
}

func Execute() {
  rootcmd.AddCommand(
    createEnvCmd(),
    addEnv(),
    deleteEnv(),
    listEnvCmd(),
    listPackagesCmd(),
    activateCmd(),
    moveCmd(),
    renameCmd(),
  )
  rootcmd.Execute();
}
