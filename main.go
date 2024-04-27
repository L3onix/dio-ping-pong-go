package main

import (
  "fmt"
  "time"
)

func main() {
  bola := make(chan string)

  go func () {
    for true {
      bola <- "ping"
    }
  } ()
  go func () {
    for true {
      msg := <- bola
      fmt.Println(msg)
      time.Sleep(time.Second * 1)
    }
  } ()
  go func () {
    for true {
      bola <- "pong"
    }
  } ()

  var input string
  fmt.Scanln(&input)
}
