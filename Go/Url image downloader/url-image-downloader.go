package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

  message := `
 +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
 |u|r|l|-|i|m|a|g|e|-|d|o|w|n|l|o|a|d|e|r|
 +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
`
  fmt.Printf("%s", message)
  fmt.Printf("Enter image url:  ")
  var fileUrl string;
  fmt.Scan(&fileUrl)
	err := DownloadFile("downloded.jpg", fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + fileUrl)
}
func DownloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}