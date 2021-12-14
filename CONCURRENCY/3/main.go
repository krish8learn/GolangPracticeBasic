package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//buffered channel
	track := time.Now()
	defer func() {
		fmt.Println(time.Since(track))
	}()
	channel := make(chan string, 5)
	go goal(channel)
	for message := range channel {
		fmt.Println(message)
	}

}
func goal(channel chan string) {
	number := 4
	for i := 0; i < number; i++ {
		rand.Seed(time.Now().UnixNano())
		score := rand.Intn(10)
		channel <- fmt.Sprint("scored--", score)
	}
	close(channel)
}
