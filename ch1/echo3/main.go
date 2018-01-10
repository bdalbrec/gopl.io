// Echo3 prints out its command-line arguments
package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  fmt.Println(strings.Join(os.Args[1:], " "))
}
