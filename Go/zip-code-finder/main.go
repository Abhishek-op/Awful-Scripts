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
  ___________ _____        ______ _____ _   _ _____  ______ _____  
 |___  /_   _|  __ \      |  ____|_   _| \ | |  __ \|  ____|  __ \ 
    / /  | | | |__) |_____| |__    | | |  \| | |  | | |__  | |__) |
   / /   | | |  ___/______|  __|   | | | .   | |  | |  __| |  _  / 
  / /__ _| |_| |          | |     _| |_| |\  | |__| | |____| | \ \ 
 /_____|_____|_|          |_|    |_____|_| \_|_____/|______|_|  \_\
                                                                   
                                                                   
`

    fmt.Printf("%s", message)
    fmt.Printf("Enter countryname(shortform):  ")
    var country string;
    fmt.Scan(&country)
    fmt.Printf("Enter zip code:  ")
    var code string;
    fmt.Scan(&code)
    response, err := http.Get("https://api.zippopotam.us/"+country+"/"+code)

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

        responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    type res struct { 
      Con    string    `json:"place name"`
      Region    string    `json:"longitude"`
      City    string    `json:"state"`
      Zip    string    `json:"latitude"`

    }
    type Response struct {
      Response []res `json:"places"`
    }  
    
    
    var responseObject Response
    json.Unmarshal(responseData, &responseObject)
    fmt.Println("Place latitude state latitude ")
    fmt.Println(responseObject)
}