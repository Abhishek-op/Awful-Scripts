package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// usage
	if len(os.Args) == 1 || os.Args[1] == "-h" {
		fmt.Printf("Ussages: go run main.go image.png file.zip")
		os.Exit(1)
	}

	fileJpg := os.Args[1]
	fileZip := os.Args[2]

	if !strings.HasSuffix(fileJpg, "png") || !strings.HasSuffix(fileZip, "zip") {
		fmt.Printf("Ussages: go run main.go file.png file.zip")
		os.Exit(1)
	}
	firstFile, err := os.Open(fileJpg)
	if err != nil {
		log.Fatal(err)
	}
	defer firstFile.Close()
	secondFile, err := os.Open(fileZip)
	if err != nil {
		log.Fatal(err)
	}
	defer secondFile.Close()
	newFile, err := os.Create("stego_image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
  var sep string;
  sep = "----------------";
	_, err = io.Copy(newFile, firstFile)
	if err != nil {
		log.Fatal(err)
	}
  newFile.WriteString(sep);
	_, err = io.Copy(newFile, secondFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Output image is created")
}