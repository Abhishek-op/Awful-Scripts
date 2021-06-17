package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
message := `
  _                        
 |_) _  ._ _|_             
 |  (_) |   |_             
                           
  __                       
 (_   _  _. ._  ._   _  ._ 
 __) (_ (_| | | | | (/_ |  
                          
`

  fmt.Printf("%s", message)
  var target string;
  fmt.Printf("[%%]Enter ip address: ")
	fmt.Scan(&target)


	activeThreads := 0
	doneChannel := make(chan bool)

	for port := 0; port <= 65535; port++ {
		go testTCPConnection(target, port, doneChannel)
		activeThreads++
	}
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
}

func testTCPConnection(ip string, port int, doneChannel chan bool) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port),
		time.Second*10)
	if err == nil {
		fmt.Printf("[-]Port %d: Open\n", port)
	}
	doneChannel <- true
}