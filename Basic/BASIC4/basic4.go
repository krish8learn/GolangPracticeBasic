package main

import "fmt"

func main() {
	//anonymous function
	anFunction := func() {
		fmt.Println("Anything")
	}
	anFunction()
	fmt.Printf("%T\n", anFunction)
	// or
	func() {
		fmt.Println("anything 2")
	}()

	fmt.Printf("%T\n", retInt)
	fmt.Printf("%T\n", retInt())

	crazy := func(f func()) func() {
		return f
	}

	crazyFun(crazy)(first)()

	//interfaces need
	stars := []ninjaStar{{owner: "torres"}, {owner: "fernando"}}
	for _, star := range stars {
		star.attack()
	}

	swords := []ninjaSword{{owner: "kyle"}, {owner: "walker"}}
	for _, sword := range swords {
		sword.attack()
	}

	weapons := []ninjaWeapon{ninjaStar{owner: "wallace"}, ninjaSword{owner: "tores"}}
	for _, weapon := range weapons {
		attackWeapon(weapon)
	}
}

type ninjaWeapon interface {
	attack()
}

func attackWeapon(n ninjaWeapon) {
	n.attack()
}

type ninjaStar struct {
	owner string
}

type ninjaSword struct {
	owner string
}

func (ninjaStar) attack() {
	fmt.Println("Throwing stars")
}

func (ninjaSword) attack() {
	fmt.Println("Throwing swords")
}
func retInt() int {
	return 3
}

func crazyFun(f func(func()) func()) func(func()) func() {
	return f
}

func first() {
	fmt.Println("first")
}
