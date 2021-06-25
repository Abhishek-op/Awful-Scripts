package main

import (
	"fmt"
  "time"
	"net/http"
)
func date(w http.ResponseWriter, req *http.Request) {
  t := time.Now()
  fmt.Printf("today date is = %s", t)
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/date", date)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8080", nil)
}