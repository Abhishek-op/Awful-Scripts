package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 1)
			c <- false
		}

		time.Sleep(time.Second * 1)
		c <- true
	}()

	go func() {
		for {
			timer := time.NewTimer(time.Second * 5)
			defer timer.Stop()
			select {
			case b := <-c:
				if b == false {
					fmt.Println(time.Now(), "timer is running")
					continue
				}
				fmt.Println(time.Now(), ":completed")
				return
			case <-timer.C:
				fmt.Println(time.Now(), ":timer expired")
				continue
			}
		}
	}()
	var s string
	fmt.Scanln(&s)
}