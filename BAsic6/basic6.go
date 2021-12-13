package main

import "fmt"

func main() {
	//type assertions and type switch
	players := []footballer{goal{striker: "Torres"}, defend{defender: "Terry"}}
	for _, player := range players {
		player.pass()
		if working, ok := player.(defend); ok {
			working.tackle() //type assertions ----> player interface type to defend struct type
		}
	}
	fmt.Println()
	for _, play := range players {
		switch play.(type) { //type switch ---> player interface type to defend struct type
		case defend:
			play.(defend).tackle()
		case goal:
			play.pass()
		}
		play.pass()
	}
}

type goal struct {
	striker string
}

type defend struct {
	defender string
}

type footballer interface {
	pass()
}

func (defend) pass() {
	fmt.Println("Pass the ball to defend")
}

func (goal) pass() {
	fmt.Println("Pass the ball to score")
}

func (defend) tackle() {
	fmt.Println("tackle the man with ball")
}
