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
NN   NN                     bb                        fff                tt          
NNN  NN uu   uu mm mm mmmm  bb        eee  rr rr     ff     aa aa   cccc tt     sss  
NN N NN uu   uu mmm  mm  mm bbbbbb  ee   e rrr  r    ffff  aa aaa cc     tttt  s     
NN  NNN uu   uu mmm  mm  mm bb   bb eeeee  rr        ff   aa  aaa cc     tt     sss  
NN   NN  uuuu u mmm  mm  mm bbbbbb   eeeee rr        ff    aaa aa  ccccc  tttt     s 
                                                                                sss   
`

    fmt.Printf("%s", message)
    var target string;
    fmt.Printf("[==]Enter number: ")
	  fmt.Scan(&target)
    response, err := http.Get("http://numbersapi.com/"+target)

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