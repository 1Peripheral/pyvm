package cmds

import (
	"fmt"

	"github.com/1peripheral/pyvm/utils"
	"github.com/spf13/cobra"
)

func renameCmd() *cobra.Command {
  var cmd  = &cobra.Command{
    Use: "rename [old_name] [new_name]",
    Short: "Rename a virtual env",
    Run: func(cmd *cobra.Command, args []string) {
      if len(args) < 2 {
        cmd.Usage()
        return
      }
      oldName := args[0]
      newName := args[1]
      err := utils.RenameEnv(oldName, newName)
      if err != nil {
        fmt.Println(err.Error())
      }
      fmt.Printf("%s renamed to %s\n", oldName, newName);
    },
  }

  cmd.SetUsageTemplate("Usage : pyvm rename [old_name] [new_name]\n")

  return cmd
}
