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
  
         )     )     (     
      ( /(  ( /(     )\ )  
   (  )\()) )\())(  (()/(  
   )\((_)\|((_)\ )\  /(_)) 
  ((_) ((_)_ ((_|(_)(_))   
 _ | |/ _ \ |/ /| __/ __|  
| || | (_) |' < | _|\__ \  
 \__/ \___/_|\_\|___|___/  
                           

`

    fmt.Printf("%s", message)
    response, err := http.Get("https://official-joke-api.appspot.com/jokes/ten")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    type Response struct {
    Name    string    `json:"setup"`
    Punch    string    `json:"punchline"`
    

    }
    var responseObject []Response
    json.Unmarshal(responseData, &responseObject)
    for _, user := range responseObject {
          fmt.Println("==>", user.Name)
          fmt.Println("Answer: ", user.Punch) 
        }
}