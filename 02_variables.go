package main

import "fmt"

func main() { 
  fmt.Println("Lesson 2: Variables")

  var a string = "Hello"
  fmt.Println(a)

  var b, c int = 1, 2
  fmt.Println(b, c)

  var d = true
  fmt.Println(d)

  var e int
  fmt.Println("Value of `var e int`:", e)

  // Shorthand for var f string = "short"
  f := "short" 
  fmt.Println("Value of `f:= \"short\"", f)
}
