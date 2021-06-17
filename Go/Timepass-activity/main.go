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
  
    )                                                  
 ( /(                 (           )             )      
 )\())  (  (  (       )\       ( /((   )  (  ( /((     
((_)\  ))\ )\))(   ((((_)(  (  )\())\ /(( )\ )\())\ )  
 _((_)/((_|(_)()\   )\ _ )\ )\(_))((_|_))((_|_))(()/(  
| \| (_)) _(()((_)  (_)_\(_|(_) |_ (_))((_|_) |_ )(_)) 
| .  / -_)\ V  V /   / _ \/ _||  _|| \ V /| |  _| || | 
|_|\_\___| \_/\_/   /_/ \_\__| \__||_|\_/ |_|\__|\_, | 
                                                 |__/  
                  
`

    fmt.Printf("%s", message)
    response, err := http.Get("https://www.boredapi.com/api/activity")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    type Response struct {
    Name    string    `json:"activity"`
    Type    string    `json:"type"`
    participants    string    `json:"participants"`

    }
    var responseObject Response
    json.Unmarshal(responseData, &responseObject)
    fmt.Printf("[==]Don't be bored, Here is your activity::::::     ")
    fmt.Println(responseObject.Name)
    fmt.Printf("[-]Activity_Type: ")
    fmt.Println(responseObject.Type)
    fmt.Printf("[-]Required Participants: ")
    fmt.Println(responseObject.participants)
}