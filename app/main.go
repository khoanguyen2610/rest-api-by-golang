package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "Ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go func() {
		c <- "Start channel"
		fmt.Println("Fired")
	}()
	fmt.Println(<-c)
	fmt.Println("Done")
}
