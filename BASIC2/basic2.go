package main

import "fmt"

func main() {
	//arrays
	var ar [3]int
	fmt.Println(ar)
	fmt.Println(ar[0])
	fmt.Println(len(ar))
	//arCopy gets the copy of elements not address
	arCopy := ar
	fmt.Println(arCopy)
	ar[0] = 2
	fmt.Println(ar == arCopy)
	//expliicit declaration
	br := [4]int{2, 4, 1}
	cr := [...]int{12, 12}
	fmt.Println(br, " and ", cr)
	fmt.Println(len(br), " and ", len(cr))
	//declaring array on specific index
	dr := [...]int{0: 23, 4: 34}
	fmt.Println(len(dr), " :and: ", dr)
	//multidimentional array
	ar3d := [3][3][3]int{{}, {}}
	fmt.Println(ar3d)
}
