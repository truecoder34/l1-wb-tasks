package main

import (
	"fmt"
	"reflect"
)

/*
14.
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

/*
Package reflect implements run-time reflection,
allowing a program to manipulate objects with arbitrary types.
The typical use is to take a value with static type interface{}
and extract its dynamic type information by calling TypeOf, which returns a Type.
*/
func main() {
	channel1 := make(chan string)
	channel2 := make(chan interface{})
	boo := true
	num1 := 400
	num2 := 400.50
	text := "text"
	var line string
	var inter interface{}

	fmt.Printf("Type of var: %s\n", detectType(channel1))
	fmt.Printf("Type of var: %s\n", detectType(channel2))
	fmt.Printf("Type of var: %s\n", detectType(boo))
	fmt.Printf("Type of var: %s\n", detectType(num1))
	fmt.Printf("Type of var: %s\n", detectType(num2))
	fmt.Printf("Type of var: %s\n", detectType(text))
	fmt.Printf("Type of var: %s\n", detectType(line))
	fmt.Printf("Type of var: %s\n", detectType(inter))
}

func detectType(inVar interface{}) reflect.Type {
	return reflect.TypeOf(inVar)
}
