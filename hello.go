package main

import "fmt"

const const1 = 53

type apple int8

func main() {
	//var msg = "Hello World!"
	msg := "golang"
	fmt.Println(msg)
	fmt.Println(const1)
	var bar apple
	//	var bar int8
	fmt.Println(bar)
	p := func1()
	fmt.Println(*p)
	*p++
	fmt.Println(*p)
}

func func1() *int {
	var a int = 5
	return &a
}
