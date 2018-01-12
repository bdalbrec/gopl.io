// this doesn't work. the only time measured is the delay time 
package main

import (
  "fmt"
  "os"
  "strings"
  "time"
)

func main() {
  start := time.Now()
  time.Sleep(5 * time.Second)
  // for loop method
  var s, sep string
  for i := 1; i < len(os.Args); i++ { // start at 1. Args[0] is the program name with fullpath
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Printf("For loop - %s - took %f \n", s, time.Since(start).Seconds())

  // range method
  start = time.Now()
  var s2, sep2 string
  for _, t := range(os.Args[1:]){
    s2 += sep2 + t
    sep2 = " "
    }
  fmt.Printf("For loop - %s - took %f \n", s2, time.Since(start).Seconds())

  // join method
  start = time.Now()
  fmt.Printf("For loop - %s - took %f \n", strings.Join(os.Args[1:], " "), time.Since(start).Seconds())
}
