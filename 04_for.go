package main

import "fmt"

func main(){ 
  var i int = 1
  for i <= 3 { 
    fmt.Println("Iteration:", i)
    i += 1
  }

  fmt.Println("---")

  for j := 7; j < 9; j++ { 
    fmt.Println("Iteration:", j)
  }

  fmt.Println("---")

  // Infinite Loop
  for { 
    fmt.Println("An infinite loop")
    break
  }
}
