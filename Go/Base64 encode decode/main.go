package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		usage(os.Args[0])
	}

	switch fl := os.Args[1]; fl {
	case "encode":
		encodedString := base64.StdEncoding.EncodeToString([]byte(os.Args[2]))
		fmt.Printf("Encoded data is:   %s", encodedString)
	case "decode":
		decodedData, err := base64.StdEncoding.DecodeString(os.Args[2])
		if err != nil {
			log.Fatal("Error decoding data. ", err)
		}
		fmt.Printf(" Decoded data is:  %s\n", decodedData)
	default:
		usage(os.Args[0])
	}
}

func usage(name string) {
	fmt.Fprintf(os.Stderr, "Ussages:%s go run main.go encode your_data \n", name)
	fmt.Fprintf(os.Stderr, "OR %s go run main.go decode your_data \n", name)
	os.Exit(1)
}