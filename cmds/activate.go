package cmds

import "github.com/spf13/cobra"

func activateEnv() (*cobra.Command) {
  var cmd = &cobra.Command{
    Use: "activate [name]",
    Short: "Activate a virtual environment",
    Run: func(cmd *cobra.Command, args []string) {
      if len(args) < 1 {
        cmd.Usage()
        return
      }
    },
  }

  cmd.SetUsageTemplate("Usage: pyvenv activate [name]")

  return cmd
}
