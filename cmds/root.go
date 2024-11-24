package cmds

import (
	"github.com/spf13/cobra"
)

var rootcmd = &cobra.Command{
  Use: "pyvenv",
  Run: func(cmd *cobra.Command, args []string) {
    cmd.Usage();
  },
}

func Execute() {
  rootcmd.AddCommand(createEnvCmd())

  rootcmd.Execute();
  
}
