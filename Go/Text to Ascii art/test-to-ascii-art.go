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
  
  _            _            _                         _____  _____ _____ _____ 
 | |          | |          | |                 /\    / ____|/ ____|_   _|_   _|
 | |_ _____  __ |_   ______| |_ ___    ______ /  \  | (___ | |      | |   | |  
 | __/ _ \ \/ / __| |______| __/ _ \  |______/ /\ \  \___ \| |      | |   | |  
 | |_  __/>  <| |_         | |_ (_) |       / ____ \ ____) | |____ _| |_ _| |_ 
  \__\___/_/\_\\__|         \__\___/       /_/    \_\_____/ \_____|_____|_____|              
`

    fmt.Printf("%s", message)
    fmt.Println("")
    fmt.Println("")
    fmt.Println("Welcome to awfulscripts")
    fmt.Println("")
    fmt.Println("")
    
    fmt.Printf("Enter text(If multiple word then use + to add them(HI+GO)):  ")
    var target string;
    fmt.Scan(&target)
    response, err := http.Get("http://artii.herokuapp.com/make?text="+target)

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    s, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    responseString := string(s)
    fmt.Print(responseString)

    
}