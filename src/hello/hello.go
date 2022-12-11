package main

import (
  "fmt"

  "github.com/tmp-friends/go-tutorial/src/greetings"
)

func main() {
  message := greetings.Hello("tmp")
  fmt.Println(message)
}
