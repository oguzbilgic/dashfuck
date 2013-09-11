package main

import (
  "io/ioutil"
  "strings"
  "os"
  "log"
)

func main() {
    dashfuckByte, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
      log.Fatal(err)
    }

    dashfuckString := string(dashfuckByte)

    // Replate en dash
    brainfuckString := strings.Replace(dashfuckString, "—", "+", -1)

    // Replate em dash
    brainfuckString = strings.Replace(brainfuckString, "–", ".", -1)

    // Replate swung dash
    brainfuckString = strings.Replace(brainfuckString, "~", ">", -1)

    brainfuckByte := []byte(brainfuckString)
    err = ioutil.WriteFile("main.bf", brainfuckByte, os.ModePerm)
    if err != nil {
      log.Fatal(err)
    }
}
