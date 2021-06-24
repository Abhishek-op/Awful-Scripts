package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "encoding/json"
)

func main() {
    message := `
   ____      _      ____          ___     _   _    U  ___ u _____  U _____ u 
U | __")uU  /"\  u |  _"\        / " \ U |"|u| |    \/"_ \/|_ " _| \| ___"|/ 
 \|  _ \/ \/ _ \/ /| | | |      | |"| | \| |\| |    | | | |  | |    |  _|"   
  | |_) | / ___ \ U| |_| |\    /| |_| |\ | |_| |.-,_| |_| | /| |\   | |___   
  |____/ /_/   \_\ |____/ u    U \__\_\u<<\___/  \_)-\___/ u |_|U   |_____|  
 _|| \\_  \\    >>  |||_          \\// (__) )(        \\   _// \\_  <<   >>  
(__) (__)(__)  (__)(__)_)        (_(__)    (__)      (__) (__) (__)(__) (__) 
                           
`

    fmt.Printf("%s", message)
    response, err := http.Get("https://breaking-bad-quotes.herokuapp.com/v1/quotes/5")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    type Response struct {
    Name    string    `json:"author"`
    Punch    string    `json:"quote"`
    

    }
    var responseObject []Response
    json.Unmarshal(responseData, &responseObject)
    fmt.Println("Here are some breaking bad quotes")
    fmt.Println("")
    for _, user := range responseObject {
          fmt.Println("Quote:  ", user.Punch)
          fmt.Println("Auther: ", user.Name) 
        }
}