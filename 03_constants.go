package main

import "fmt"
import "math"

const c string = "I AM A CONSTANT"

func main() { 
  fmt.Println("Lesson 3: Constants")
  fmt.Println(c)

  const n = 500000

  const d = 3e20 / n

  fmt.Println("Value of d:", d)
  fmt.Println("Value of `int64(d)`:", int64(d))
  fmt.Println("Value of `math.Sin(d)`:", math.Sin(d))
}
