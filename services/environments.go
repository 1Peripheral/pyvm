package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const ENVIRONMENTS_FILENAME = ".pyvm.json"

// var Environments = struct {
//   Entries map[string]interface{}
//   Count int
// }{}
var Environments = make(map[string]interface{})

func InitEnv() {
  homeDir, _ := os.UserHomeDir()

  filePath := filepath.Join(homeDir, ENVIRONMENTS_FILENAME)

  // Initialise file if it does not exist
  if _, err := os.Stat(filePath) ; os.IsNotExist(err) {
    initialData := Environments
    fileContent, err := json.Marshal(initialData)
    if err != nil {
      fmt.Println("Failed to initiate environments data file .")
      os.Exit(1)
    }

    envsFile, err:= os.OpenFile(filePath, os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
      fmt.Println("Failed to initiate environments data file .")
      os.Exit(1)
    }
    defer envsFile.Close()

    _, err = envsFile.Write(fileContent)
    if err != nil {
      fmt.Println("Error while writing to the data file")
      os.Exit(1)
    }
    return
  }
  envsFile, err := os.OpenFile(filePath, os.O_RDWR, 0644)
  if err != nil {
    fmt.Println("Failed to initiate environments data file .")
    os.Exit(1)
  }
  defer envsFile.Close()

  decoder := json.NewDecoder(envsFile)
  err = decoder.Decode(&Environments)
  if err != nil && err.Error() != "EOF" {
    fmt.Println("Failed to open the data file.")
    os.Exit(1)
  }
}

func AddEnv(name, path string) error {
  if Environments == nil {
    Environments = make(map[string]interface{})
  }
  _, exists := Environments[name]
  if exists {
    return fmt.Errorf("An environment already exists using this name : %s\n", name)
  }

  Environments[name] = path
  return nil
}

func SaveChanges() {
  homeDir, _ := os.UserHomeDir()
  filePath := filepath.Join(homeDir, ENVIRONMENTS_FILENAME)

  envsFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)

  if err != nil {
    fmt.Println("Failed to open environments data file for saving.")
    os.Exit(1)
  }

  defer envsFile.Close()

  fileContent, err := json.Marshal(Environments)
  if err != nil {
    fmt.Println("Failed to marshal Environments data.")
    os.Exit(1)
  }

  _, err = envsFile.Write(fileContent)
  if err != nil {
    fmt.Println("Error while saving to the environments data file.")
    os.Exit(1)
  }
}
