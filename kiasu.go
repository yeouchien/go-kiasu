package main

import (
  "github.com/go-martini/martini"
  "github.com/robfig/cron"
  "fmt"
  "os"
  "os/exec"
)

func main() {
  m := martini.Classic()
  c := cron.New()
  c.AddFunc("@every 20s", func() { 
    fmt.Println("kia") 
    KiasuAppend()
    KiasuCommit()
    KiasuPush()
  })
  c.Start()
  m.Run()
}

func KiasuAppend() {
  f, _ := os.OpenFile("kiasu.log", os.O_APPEND|os.O_WRONLY, 0600)
  f.WriteString("kiasu\n")
}

func KiasuCommit() {
  exec.Command("git", "commit", "-am", "I kiasu everyday").Output()
}

func KiasuPush() {
  exec.Command("git", "push", "origin", "master").Output()
}
