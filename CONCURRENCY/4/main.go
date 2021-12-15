package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	defer func() {
		fmt.Println(time.Since(t))
	}()
	player1, player2 := make(chan string), make(chan string)

	random(player1, player2)

}

func random(ninja1 chan string, ninja2 chan string) {
	close(ninja1)
	close(ninja2)
	var count1, count2 int
	for i := 0; i < 100; i++ {
		select {
		case <-ninja1:
			count1++
		case <-ninja2:
			count2++
		}
	}

	fmt.Printf("player1--> %d\n player2--> %d\n", count1, count2)
}
