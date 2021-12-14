package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	secret1 := secret{Name: "krish", Secret: "Coder"}
	res, err := template.New("rand.txt").ParseFiles("/home/krishnendu/Go/src/github.com/krish8learn/GolangPracticeBasic/Basic3/rand.txt")
	if err != nil {
		fmt.Println("Error:-->", err)
	}

	err = res.Execute(os.Stdout, secret1)
	if err != nil{
		fmt.Println(err)
	}
}

type secret struct {
	Name   string
	Secret string
}
