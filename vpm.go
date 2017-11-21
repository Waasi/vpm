package main

import(
  "fmt"
  "os"
  "os/user"
  "os/exec"
  "path/filepath"
  "strings"
  "io/ioutil"
)

func main() {
  currentUser, _ := user.Current()
  config := filepath.Join(currentUser.HomeDir, ".vim/pack/plugins/start")

  if len(os.Args) < 2 {
    fmt.Println("please specify a command")
    os.Exit(1)
  }

  switch os.Args[1] {
  case "init":
    os.MkdirAll(config, os.ModePerm)
    fmt.Println("initialized")
  case "list":
    plugins, err := ioutil.ReadDir(config)
    if err != nil { os.Exit(1) }

    for _, plugin := range plugins {
      fmt.Println(plugin.Name())
    }
  case "add":
    plugin := os.Args[2]

    fmt.Printf("JJ: %s", invalid(plugin))
    if invalid(plugin) {
      fmt.Println("please specify a valid git url")
      os.Exit(1)
    }

    path := filepath.Join(config, name(plugin))
    exec.Command("git", "clone", plugin, path).Run()
    fmt.Printf("added plugin %s\n", name(plugin))
  case "remove":
    plugin := os.Args[2]
    path := filepath.Join(config, plugin)

    if _, err := os.Stat(path); err == nil {
      if err := os.RemoveAll(path); err == nil {
        fmt.Printf("removed plugin %s\n", plugin)
      } else {
        fmt.Printf("unable to remove plugin %s\n", plugin)
      }
    } else {
      fmt.Printf("plugin %s already removed\n", plugin)
    }
  default:
    fmt.Println("please specify a valid command (init | add | remove)")
    os.Exit(1)
  }
}

func name(plugin string) string {
  nameWithExtension := strings.Split(plugin, "/")
  return strings.Trim(nameWithExtension[len(nameWithExtension) - 1], ".git")
}

func invalid(plugin string) bool {
  url := strings.Split(plugin, "/")
  last := strings.Split(url[len(url) - 1], ".")
  return url[0] != "https:" || last[len(last) - 1] != "git"
}
