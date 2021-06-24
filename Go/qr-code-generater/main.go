package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

  message := `
_________________________________________________________________________
¦¦¦__¦¦¦¦__¯¦¦¦¦¦__¦¦¦¦___¦¦¦¯¦¦¦¦¦¦___¦¦¦__¯¦¦__¯¦¦__¯¦__¦__¦¦¦___¦¦¦__¯
¦¦¦¦¦¦¦¦¦¯¯_¦¦¦¦¦¦¯¯¦¦¦___¦¦¦¦¦¦¦¦¦¦___¦¦¦¯¯_¦¦¯¯¦¦¦¯¯¦¦¦¦¦¦¦¦¦¦___¦¦¦¯¯_
¦¦__¦¯¦¦¦¦¦¦¦¦¦¦¦¯¯_¦¦¦¯¯¯¦¦¦¦¦_¦¦¦¦¯¯¯¦¦¦¦¦¦¦¦¦¦¦¦¦¦¦¦¦¦¦¦¦¦¦¦¦¯¯¯¦¦¦¦¦¦
¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯

`
  fmt.Printf("%s", message)
  fmt.Printf("Enter text or url:  ")
  var st string;
  fmt.Scan(&st)
  var Url string;
  Url = "https://image-charts.com/chart?chs=150x150&cht=qr&chl="+st+"&choe=UTF-8"
	err := DownloadFile("QR.jpg", Url)
	if err != nil {
		panic(err)
	}
	fmt.Println("QR_CODE is Generaated as QR.jpg")
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