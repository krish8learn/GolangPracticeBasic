package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	defer safeExit()
	//standard error  and custom error
	value, err := retGolangError(true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	value2, err := retCustomError(true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value2)
	}

	//panic , recover, defer
	val := 3
	if val != 1 {
		panic("something wrong")
	}

}

func safeExit() {
	if r := recover(); r != nil {
		fmt.Println("Panic recovered")
	}
}

type specialError struct {
	Message string
	Code    int
}

func retGolangError(boolError bool) (string, error) {
	if boolError {
		return "", errors.New("standard Golang Error Use")
	}

	return "fine", nil
}

func (s specialError) Error() string {
	return s.Message + "," + strconv.Itoa(s.Code)
}

func retCustomError(boolError bool) (string, error) {
	if boolError {
		return "", specialError{"special Error", 123}
	}

	return "fine", nil
}
