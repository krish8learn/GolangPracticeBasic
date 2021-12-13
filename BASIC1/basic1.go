package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

//first letter matters , decide the scope of function/variables
var p = "package level, only files inside the package can access this"
var G = "Global level, exported, anyone can use it importing the location"

//constant
const (
	_ = 1 << (10 * iota) //<< binary shift to left side
	kb
	mb
	gb
	tb
	pb
	eb
	zb
	yb //yb cannot be print but operation can be performed
)

func main() {
	//variable declaration
	B := "block level, only accesible inside the function(short hand declaration)"
	fmt.Println(B)

	//type declaration
	type farenheit int
	type celcius int
	var far farenheit = 20
	var cel celcius = -23
	fmt.Println(far, cel)

	//type checking
	fmt.Printf("This show type upper level--->%T and %T\n", far, cel)
	fmt.Println("This show type upper level-->", reflect.TypeOf(far), " and shows underlying type-->", reflect.ValueOf(cel).Kind())

	//declaring two diffeent types of variable in one line
	var integer1, string1 = 1, "krish"
	fmt.Println(integer1, string1, " and types are ", reflect.TypeOf(integer1), reflect.ValueOf(string1).Kind())

	//pointers
	x := 14
	pointer := &x
	//x and *pointer are same
	if x == *pointer {
		fmt.Println("same")
	}
	fmt.Println("this will print address and value respectively", pointer, *pointer)

	//math package
	fmt.Println(math.Ceil(.00001))  //1
	fmt.Println(math.Floor(.00001)) //0
	fmt.Println(math.Min(.00001, 18))
	fmt.Println(math.Max(.00001, 23))
	fmt.Println(math.Abs(-1))   //1
	fmt.Println(math.Sqrt(100)) //10
	fmt.Println(math.Pow(2, 3)) //8

	//string(slice of bytes)
	example := "Hello World"
	fmt.Println(example[2:5])
	fmt.Printf("%c\n", example[2])
	fmt.Println("Hello \bWorld")
	by := []byte(example)
	fmt.Printf("%s %s\n", example, by)
	bystr := string(by)
	fmt.Println(bystr)

	var sb strings.Builder
	fmt.Println(sb.String(), " and ", sb, " and ", sb.Cap())
	sb.WriteString("Krish")
	fmt.Println("cap-->", sb.Cap(), "length-->", sb.Len())
	sb.Grow(40) //double the initial capacity then add
	fmt.Println("Len-->", sb.Len(), "Cap-->", sb.Cap())
	sb.Write([]byte{65, 56, 89})
	fmt.Println(sb, " and ", sb.String())

	//constant
	fmt.Println(kb, mb, gb, tb, pb)

	//label
iForLoop:
	for i := 0; i < 10; i++ {
		for j := 10; j > 0; j-- {
			if j == 3 {
				break iForLoop
			}
			fmt.Println(i, j)
		}
	}

	

}
