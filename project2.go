//  Code to find solution for project euler problem part deux!!

package main

import "fmt"

func main() {
var sum, x1, x2, flag int = 2, 1, 2, 0  //sum variable starts at 2

  for x1 < 4000000 && x2 < 4000000 {

    if flag == 0 {
      x1 += x2
      if (x1 % 2) == 0 {
        sum += x1
      }
      flag++
      fmt.Println("x1:" , x1)

    } else if flag == 1 {

      x2 += x1
      if (x2 % 2) == 0 {
        sum += x2
      }
      flag--
      fmt.Println("x2:" , x2)
    }
  }

  fmt.Println(sum)
}
