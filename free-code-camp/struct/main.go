package main

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
	user
}

type user struct {
	name   string
	number int
}



func canSendMessage(mToSend messageToSend) bool {

	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.sender.number ==0  {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}

	fmt.Println(mToSend)
	return true
}

type rect struct {
	width int 
	height int
}


type authenticationInfo struct {
	username string
	password string
}

func (authInfo authenticationInfo) authenticate () {
	if authInfo.username == "a" && authInfo.password == "123" {
		fmt.Println("login successfully")
		} else {
		fmt.Println("login failed")

	}
}

func  (r rect) area() int {
	return r.height * r.width
} 

func main() {
	//  sendMessage := messageToSend {
	// 	user: user{name: "nothing", number: 5},
	// 	message: "this is message",
	// 	sender: user{name: "Jo",number: 1},
	// 	recipient: user{name: "Jo",number: 1},
	// }


	// fmt.Println(sendMessage.name)

	//  ok := canSendMessage(sendMessage)

	//  if ok {
	// 	fmt.Println("message can be send")
	// 	} else {
	// 	fmt.Println("message cannot be send")

	//  }

	// r := rect{
	// 	width: 5,
	// 	height: 10,
	// }

	// fmt.Println(r.area())


	authData := authenticationInfo{
		username: "a",
		password: "12",
	}
	authData.authenticate()
}