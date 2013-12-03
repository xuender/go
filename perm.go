package main

import (
  "fmt"
  "github.com/ntns/goitertools/itertools"
)

func main(){
  fmt.Println("xxx")
  a := itertools.Permutations([]int{1, 2, 3}, 2)
  fmt.Println(a)
}
