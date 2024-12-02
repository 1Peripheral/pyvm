package cmds

import (
	"fmt"
	"os"

	"github.com/1peripheral/pyvenv/utils"
	"github.com/spf13/cobra"
)

var pyCmd = "python3"

func checkPython() bool {
  var err error
  commands := []string{"python3", "python", "py"}
  for _, cmdName := range commands {
    err = utils.ExecuteCmd(pyCmd + " --version")
    if err == nil {
      pyCmd = cmdName
      return true
    }
    fmt.Println(err.Error())
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
    listEnvCmd(),
    listPackagesCmd(),
    deleteEnv(),
  )
  rootcmd.Execute();
}
