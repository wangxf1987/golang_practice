package main

import (
    "fmt"
    "github.com/pkg/errors"
)

func main() {
   var a = []int{1,3,4}
   fmt.Println(a[:])
   fmt.Println(errors.Wrapf("test"))
}
