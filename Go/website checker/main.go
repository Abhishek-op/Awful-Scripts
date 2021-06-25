package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
  fmt.Print("[+]Enter url:  ")
	var target string;
  fmt.Scan(&target)
	u, err := url.ParseRequestURI(target)
	if err != nil {
		log.Fatal(err)
	}
	pwdDir, _ := os.Getwd()
	fLog, err := os.OpenFile(pwdDir+`/`+u.Hostname()+`.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
  fmt.Println("[-]Log file is created for "+target)
  fmt.Println("[-]Logs are collecting")
  fmt.Println("[-]Press ctrl+c to exit")
  
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()
	log.SetOutput(fLog)

	for {
		resp, err := http.Get(u.String())
		if err != nil {
			log.Printf("Error connection %s\n", err)
			return
		}

		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			log.Printf("Error. %s http-status: %d\n", u.String(), resp.StatusCode)
			return
		}

		log.Printf("Online. %s http-status: %d\n", u.String(), resp.StatusCode)

		time.Sleep(3 * time.Second)
	}
}
