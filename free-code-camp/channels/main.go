// package main

// import (
// 	"fmt"
// 	"time"
// )

// func filterOldEmails(emails []email) {
// 	isOldChan := make(chan bool)

// 	go func ()  {
// 		for _, e := range emails {
// 			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
// 				isOldChan <- true
// 				continue
// 			}
// 			isOldChan <- false
// 		}
// 	}()

// 	isOld := <-isOldChan
// 	fmt.Println("email 1 is old:", isOld)
// 	isOld = <-isOldChan
// 	fmt.Println("email 2 is old:", isOld)
// 	isOld = <-isOldChan
// 	fmt.Println("email 3 is old:", isOld)
// }

// // TEST SUITE -- Don't touch below this line

// type email struct {
// 	body string
// 	date time.Time
// }

// func test(emails []email) {
// 	filterOldEmails(emails)
// 	fmt.Println("==========================================")
// }

// func main() {
// 	test([]email{
// 		{
// 			body: "Are you going to make it?",
// 			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "I need a break",
// 			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "What were you thinking?",
// 			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 	})
// 	test([]email{
// 		{
// 			body: "Yo are you okay?",
// 			date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "Have you heard of that website Boot.dev?",
// 			date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "It's awesome honestly.",
// 			date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 	})
// 	test([]email{
// 		{
// 			body: "Today is the day!",
// 			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "What do you want for lunch?",
// 			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "Why are you the way that you are?",
// 			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 	})
// 	test([]email{
// 		{
// 			body: "Did we do it?",
// 			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "Letsa Go!",
// 			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "Okay...?",
// 			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 	})
// }

package main

import (
	"fmt"
	"time"
)

func waitForDbs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

// don't touch below this line

func test(numDBs int) {
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
	fmt.Println("All databases are online!")
	fmt.Println("=====================================")
}

func main() {
	test(3)
	test(4)
	test(5)
}

func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()
	return ch
}