package main

import (
	"fmt"
	"time"
)

func main() {
	inst := instance{
		month: 0,
	}
	fmt.Println("Value of month 1--->", inst.month)
	fmt.Printf("Type of month---> %T\n", inst.month)
	fmt.Println("Value of inst.month.String()--->", inst.month.String())
	fmt.Printf("Type of inst.month.String()---> %T\n", inst.month.String())
	fmt.Println(inst.month - 2)
	fmt.Printf("Type of (inst.month - 2)--->%T\n", (inst.month - 2))
	fmt.Println("date, time with location--->", time.Now())

	currentTime := time.Now()
	fmt.Println(currentTime)
	// Subtract one hour from current time
	fmt.Println(currentTime.Add(-time.Hour * 1))

}

type instance struct {
	month time.Month
}
