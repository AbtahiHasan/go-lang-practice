package main

import "fmt"

// type employee interface {
// 	getName() string
// 	getSalary() int
// }

// type contractor struct {
// 	name         string
// 	hourlyPay    int
// 	hoursPerYear int
// }

// func (data contractor) getName() string {
// 	return data.name
// }

// func (data contractor) getSalary() int {
// 	return data.hourlyPay * data.hoursPerYear
// }


func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	} else {
		return 0.01 * float64(len(e.body))
	}
}


func (e email) print() {
	fmt.Println(e.body)
}

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool 
	body string
}

func test(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n",e.cost())
	p.print()
	
	fmt.Println("====================")
}

func main() {
	// em := contractor{
	// 	name:         "jo",
	// 	hourlyPay:    40,
	// 	hoursPerYear: 200,
	// }

	// fmt.Println(em.getSalary())
	// fmt.Println(em.getName())

	e := email{
		isSubscribed: true,
		body: "hello there",
	}

	test(e,e)

	e = email{
		isSubscribed: false,
		body: "i want my money back",
	}
	test(e,e)
}