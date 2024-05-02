package main

import (
	"log"

	"github.com/alexshd/total-coder-w17/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	timeStringS := "2024-03-15"
// 	timeStringF := "2024-02-15"
// 	theTime1, err := time.Parse("2006-01-02", timeStringS)
// 	if err != nil {
// 		fmt.Println("Could not parse time:", err)
// 	}
// 	theTime2, err := time.Parse("2006-01-02", timeStringF)
// 	if err != nil {
// 		fmt.Println("Could not parse time:", err)
// 	}
// 	fmt.Println("The time is", theTime1.Before(theTime2))
//
// 	fmt.Println(theTime1.Format(time.RFC3339))
// }
