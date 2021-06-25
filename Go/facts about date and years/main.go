package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func main() {
    message := `
  _____        _           ___     __                __           _       
 |  __ \      | |         / \ \   / /               / _|         | |      
 | |  | | __ _| |_ ___   / / \ \_/ /__  __ _ _ __  | |_ __ _  ___| |_ ___ 
 | |  | |/ _  | __/ _ \ / /   \   / _ \/ _  | '__| |  _/ _ |/ __| __/ __|
 | |__| | (_| | |_  __// /     | |  __/ (_| | |    | || (_| | (__| |_\__ \
 |_____/ \__,_|\__\___/_/      |_|\___|\__,_|_|    |_| \__,_|\___|\__|___/
                                                                          
`

    fmt.Printf("%s", message)
    var target string;
    fmt.Printf("[==]Press 1 for date facts.Press 2 for year facts.Enter your choice: ")
	  fmt.Scan(&target)
    if target == "1" {
      var tar string;
      fmt.Printf("[=]Enter date(formet:= /month/date): ")
	    fmt.Scan(&tar)
      response, err := http.Get("http://numbersapi.com"+tar+"/date")
      if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
      }
      responseData, err := ioutil.ReadAll(response.Body)
      if err != nil {
        log.Fatal(err)
      }
      responseString := string(responseData)
      fmt.Println(responseString)
    } else if target == "2" {
      var tar string;
      fmt.Printf("[=]Enter year: ")
	    fmt.Scan(&tar)
      response, err := http.Get("http://numbersapi.com/"+tar+"/year")
      if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
      }
      responseData, err := ioutil.ReadAll(response.Body)
      if err != nil {
        log.Fatal(err)
      }
      responseString := string(responseData)
      fmt.Println(responseString)

    }
}