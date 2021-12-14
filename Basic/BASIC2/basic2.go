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
	cr[0] = 11
	fmt.Println(br, " and ", cr)
	fmt.Println(len(br), " and ", len(cr))
	//declaring array on specific index
	dr := [...]int{0: 23, 4: 34}
	fmt.Println(len(dr), " :and: ", dr)
	//multidimentional array
	ar3d := [3][3][3]int{{}, {}}
	fmt.Println(ar3d)

	//dynamic arrays ---> slices
	fixed := [...]int{2, 4}
	//fixed = append(fixed, 3) --throw compile error not unable to add element
	fmt.Println(fixed)

	dynamic := []int{5, 6}
	dynamic = append(dynamic, 4)
	fmt.Println(len(dynamic), " and ", cap(dynamic))
	dynamic[0] = 7
	fmt.Println(dynamic)
	//slice allocation
	dynamic2 := make([]int, 8)
	fmt.Println(dynamic2, " and ", len(dynamic2), " and ", cap(dynamic2))
}
