package services

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const ENVIRONMENTS_FILENAME = ".pyvenv.json"

var Environments struct {
  Entries map[string]string
  Count int
}

func InitEnv() {
  homeDir, _ := os.UserHomeDir()

  filePath := filepath.Join(homeDir, ENVIRONMENTS_FILENAME)
  
  envsFile, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
  if err != nil {
    panic(1)
  } 
}
