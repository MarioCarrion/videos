package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	newYork, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("time.LoadLocation %v\n", err)
	}

	now := time.Date(2022, time.February, 3, 8, 0, 0, 0, newYork)

	future := now.Add(time.Hour * 2)
	past := now.Add(time.Hour * -2)

	fmt.Printf("now(future) %v (-1 = future)\n", now.Compare(future))
	fmt.Printf("now(past) %v (+1 = past)\n", now.Compare(past))
	fmt.Printf("now(now) %v (0 = same)\n", now.Compare(now))

	//-

	fmt.Println(now.Format(time.DateTime))
	fmt.Println(now.Format(time.DateOnly))
	fmt.Println(now.Format(time.TimeOnly))
}
