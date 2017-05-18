package main
import "fmt"
func main () {
  var i int = 3 
  var p *int
  p = &i
  fmt.Println(*p)
  z := *p * *p
  fmt.Println(z)
  *p = 33 
  fmt.Println(i)
}

