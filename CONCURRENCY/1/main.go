package main

import (
	"fmt"
	"time"
)

func main() {
	//go routines
	start := time.Now()

	defer func() {
		fmt.Println("Duration: ", time.Since(start))
	}()

	teams := []string{"torres", "lampard", "puyol", "neuer"}

	for _, player := range teams {
		go pass(player)
	}

	time.Sleep(time.Second * 2)

}

func pass(player string) {
	fmt.Println("Pass the ball " + player)
	time.Sleep(time.Second)
}
