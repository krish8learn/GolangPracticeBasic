package main

import "fmt"

func main() {
	//interface as generic, only used in custom package ---> https://pkg.go.dev/golang.org/dl/gotip
	show(1, 2, 3)
	//show("k", "b")
	
	//runes
	
}

// func show [T interface{}](output ...T){
// 	fmt.Println(output)
// }

func show(integer ...int) {
	fmt.Println(integer)
}
