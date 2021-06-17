package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "log"
    "encoding/json"
)

func main() {
    message := `
  
 (   (           (    (     (       (   (     
 )\ ))\ )   (    )\ ) )\ )  )\ )    )\ ))\ )  
(()/(()/(   )\  (()/((()/( (()/((  (()/(()/(  
 /(_))(_)|(((_)( /(_))/(_)) /(_))\  /(_))(_)) 
(_))(_))  )\ _ )(_))_(_))_ (_))((_)(_))(_))   
|_ _| _ \ (_)_\(_)   \|   \| _ \ __/ __/ __|  
 | ||  _/  / _ \ | |) | |) |   / _|\__ \__ \  
|___|_|   /_/ \_\|___/|___/|_|_\___|___/___/  
                                              
 
`

    fmt.Printf("%s", message)
    var target string;
    fmt.Printf("[==]Enter ip address: ")
	  fmt.Scan(&target)
    response, err := http.Get("http://ip-api.com/json/"+target)

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    type Response struct {
    Con    string    `json:"country"`
    Region    string    `json:"regionName"`
    City    string    `json:"city"`
    Zip    string    `json:"zip"`
    Lat    string    `json:"lat"`
    Lon    string    `json:"lon"`
    Time    string    `json:"timezone"`
    ISP    string    `json:"isp"`
    Org    string    `json:"org"`
    As    string    `json:"as"`

    }
    var responseObject Response
    json.Unmarshal(responseData, &responseObject)
    fmt.Printf("[-]Country: ")
    fmt.Println(responseObject.Con)
    fmt.Printf("[-]Region: ")
    fmt.Println(responseObject.Region)
    fmt.Printf("[-]City: ")
    fmt.Println(responseObject.City)
    fmt.Printf("[-]Zip: ")
    fmt.Println(responseObject.Zip)
    fmt.Printf("[-]Lat: ")
    fmt.Println(responseObject.Lat)
    fmt.Printf("[-]Lon: ")
    fmt.Println(responseObject.Lon)
    fmt.Printf("[-]Time: ")
    fmt.Println(responseObject.Time)
    fmt.Printf("[-]ISP: ")
    fmt.Println(responseObject.ISP)
    fmt.Printf("[-]ORG: ")
    fmt.Println(responseObject.Org)
    fmt.Printf("[-]As: ")
    fmt.Println(responseObject.As)
}