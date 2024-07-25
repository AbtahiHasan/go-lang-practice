// package main

// import "fmt"

// func main() {
// 	var i int = 60

// 	switch i {
// 	case 10:
// 		fmt.Println("i is 10")
// 	case 20: 
// 		fmt.Println("i is 20")
// 	case 30: 
// 		fmt.Println("i is 30")
// 	case 40:
// 		fmt.Println("i is 40")
// 	case 50:
// 		fmt.Println("i is 50")
// 	default:
// 		fmt.Println("i is greater than 50")
// 	}


// 	var typeCheck = func(i interface{}) {
// 		switch i.(type) {
// 		case int: 
// 			fmt.Println("i is int")
// 		case string: 
// 			fmt.Println("i is string")
// 		default: 
// 			fmt.Println("i is other type")
// 		}
// 	}

// 	typeCheck(0.52)

// }