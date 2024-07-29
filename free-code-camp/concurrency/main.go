package main

import (
	"fmt"
	"time"
)

func sendMail(message string) {
	func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
		}()
	go fmt.Printf("Email sent: '%s'\n", message)
}



func main() {
	
	sendMail("hi")
	sendMail("hello")
	sendMail("how are you")
}