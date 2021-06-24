package main

import (
	"crypto/sha1"
	"fmt"
)
func main() {
  fmt.Println("Enter password")
	var target string
  fmt.Scanln(&target)
	h := sha1.New()
	h.Write([]byte(target))
	bs := h.Sum(nil)
	fmt.Println(target)
	fmt.Printf("%x\n", bs)
}