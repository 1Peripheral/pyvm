package utils

import (
  "fmt"

	"os/exec"
	"strings"
)

func ExecuteCmd(cmd string) (string, error) {
  cmdArgs := strings.Fields(cmd)
  if len(cmdArgs) < 1 {
    return "", fmt.Errorf("No args given")
  }

  consoleCmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
  output, err := consoleCmd.Output()
  if err != nil {
    return "", err
  }

  return string(output), nil 
}
