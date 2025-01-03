package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const ENVIRONMENTS_FILENAME = ".pyvm.json"

var Environments = make(map[string]string)

func InitEnv() {
  homeDir, _ := os.UserHomeDir()

  filePath := filepath.Join(homeDir, ENVIRONMENTS_FILENAME)

  // Initialize file if it does not exist
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
    Environments = make(map[string]string)
  }

  if DoesEnvExist(name) {
    return fmt.Errorf("An environment already exists using this name : %s\n", name)
  }

  Environments[name] = path
  return nil
}

func DoesEnvExist(name string) bool {
  _, exists := Environments[name]
  return exists
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

func DeleteEnv(name string) error {
  path, ok := Environments[name]
  if !ok {
    return fmt.Errorf("Name not existant")
  }

  if err := os.RemoveAll(path); err != nil {
    return fmt.Errorf("Failed to remove the env path :" + err.Error())
  }
  delete(Environments, name)
  return nil
}

func PrintEnvs() {
  fmt.Println("Available Python Virtual Environments :")
  for name, path:= range Environments {
    fmt.Printf("%s  ->  %s\n", name, path)
  }
}

func GetPath(name string) (string, error) {
  path, exists := Environments[name]
  if !exists {
    return "", fmt.Errorf("Name not existant")
  }
  return path, nil
}

func ListPackages(name string) error {
  path, err := GetPath(name)
  if err != nil {
    return err
  }

  switch runtime.GOOS {
  case "linux": {
    path = filepath.Join(path, "/bin")
  }
  case "windows": {
    path = filepath.Join(path, "/Scripts")
  }
  default: {}
  }

  output, err := ExecuteCmd(path + "/pip list")
  if err != nil {
    return err
  }
  fmt.Println(output)
  return nil
}

func ActivateEnv(name string) error {
  venvPath, exists := Environments[name] 
  if !exists {
    return fmt.Errorf("Name not existant")
  }

	var activateScript string
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		activateScript = fmt.Sprintf("%s\\Scripts\\activate.bat", venvPath)
		cmd = exec.Command("cmd.exe", "/K", activateScript)
	case "linux", "darwin":
		activateScript = fmt.Sprintf("%s/bin/activate", venvPath)
    fmt.Println(activateScript)
		cmd = exec.Command("bash", "-i", "-c", fmt.Sprintf("source %s && exec $SHELL", activateScript))

	default:
    return fmt.Errorf("Unsupported OS: %s\n", runtime.GOOS)

	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin


	if err := cmd.Run(); err != nil {
    return fmt.Errorf("Failed to execute command: %v\n", err)
	}

  return nil
}

func RenameEnv(oldName, newName string) error {
  path, err := GetPath(oldName)
  if err != nil {
    return err
  }

  delete(Environments, oldName)
  AddEnv(newName, path)

  return nil
}

func MoveEnv(name, newPath string) error {
  path, err := GetPath(name)
  if err != nil {
    return err
  }

  err = os.Rename(path, newPath)
  if err != nil {
    return err
  }
  
  return nil
}
