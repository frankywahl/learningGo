package main

import "fmt"
import "time"

func main() {
  i := 2

  fmt.Println("write", i, "as") 

  switch i { 
  case 1: 
    fmt.Println("one")
  case 2: 
    fmt.Println("two")
  case 3: 
    fmt.Println("three")
  }

  switch time.Now().Weekday() { 
  case time.Saturday, time.Sunday: 
    fmt.Println("Weekend")
  default: 
    fmt.Println("Week-day")
  }

  fmt.Println(time.Weekday(0))

}
