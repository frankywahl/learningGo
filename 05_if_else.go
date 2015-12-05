package main 

import "fmt"

func main() {
  if 7 % 2 == 0{ 
    fmt.Println("7 is even") 
  } else { 
    fmt.Println("7 is odd") 
  }

  if 8 % 4 == 0 { 
    fmt.Println("8 is dividable by 4")
  }

  if 8 % 3 == 0 { fmt.Println("8 is dividable by 3") } else { fmt.Println("8 is not dividable by 3") }

  if num := 9; num < 0 { 
    fmt.Println(num, "is negative") 
  } else if num < 10 { 
    fmt.Println(num, "has 1 digit") 
  } else { 
    fmt.Println(num, "has multiple digits") 
  }


}

