package main

import (
	"fmt"
	"time"
)

func main() {
	track := time.Now()
	defer func() {
		fmt.Println(time.Since(track))
	}()
	signal := make(chan bool)
	player := "casillas"
	go score(player, signal) //main go routine will be executed before score go routine finished
	fmt.Println(<-signal)
}
func score(goalKeeper string, commited chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Shoot at ", goalKeeper)
	commited <- true
}
