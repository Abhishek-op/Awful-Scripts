package main

import (
	"fmt"
	"strings"
)
func main() {
  var target string;
  fmt.Println("[+]]Press 1 to encrypt, 2 to deccrypt ")
  fmt.Println("[-]Enter your choise:  ")
	fmt.Scan(&target)
  if target == "1" {
    var data string;
    fmt.Printf("[-]Enter data: ")
	  fmt.Scan(&data)
    var key int;
    fmt.Printf("[-]Enter key: ")
	  fmt.Scan(&key)
    fmt.Println(encrypt(data, key))
  } else if target == "2" {
    var data string;
    fmt.Printf("[-]Enter encrypted text: ")
	  fmt.Scan(&data)
    var key int;
    fmt.Printf("[-]Enter key: ")
	  fmt.Scan(&key)
    fmt.Println(decrypt(data, key))
  }
}
func encrypt(data string, key int) string {
	result := strings.Map(func(r rune) rune {
		return caesar(r, -key)
	}, data)
	return result
}

func decrypt(data string, key int) string {
	result := strings.Map(func(r rune) rune {
		return caesar(r, +key)
	}, data)
	return result
}

func caesar(r rune, shift int) rune {
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}