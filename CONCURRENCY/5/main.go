package main

import (
	"fmt"
	"sync"
)

func main() {
	players := []string{"tommy", "Johnny", "torres"}
	var signal sync.WaitGroup
	signal.Add(len(players))
	for _, player := range players {
		go attack(player, &signal)
	}

	signal.Wait()
}

func attack(players string, signal *sync.WaitGroup) {
	fmt.Println("attack ", players)
	signal.Done()
}

//done method call must be equal to the input of add
