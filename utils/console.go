package utils

import (
  "fmt"

	"os/exec"
	"strings"
)

func ExecuteCmd(cmd string) error {
  cmdArgs := strings.Fields(cmd)
  if len(cmdArgs) < 1 {
    return fmt.Errorf("No args given")
  }

  consoleCmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
  err := consoleCmd.Run()
  if err != nil {
    return err
  }

  return nil 
}
