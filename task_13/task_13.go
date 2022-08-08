package main

import (
	"fmt"
	"reflect"
)

/*
13.
	Поменять местами два числа без создания временной переменной.
*/

func main() {
	// 0 - BASE SWAP
	val1 := 100
	val2 := 200
	fmt.Printf("[BEFORE] : val1=%d , val2=%d \n", val1, val2)
	val1, val2 = val2, val1
	fmt.Printf("[AFTER-1] : val1=%d , val2=%d \n", val1, val2)

	val1, val2, val1 = val2, val1, val2
	fmt.Printf("[AFTER-2] : val1=%d , val2=%d \n", val1, val2)

	// 1 - SWAP IN SLICE with reflect
	fmt.Println("SWAP IN SLICE ")
	slize := []int{100, 200, 300}
	fmt.Printf("[BEFORE] : %v\n", slize)
	swapFunc := reflect.Swapper(slize)
	swapFunc(1, 2) // swap 1 2 index elements
	fmt.Printf("[AFTER] : %v\n", slize)

	// 3 - MATH SWAP
	fmt.Println("MATH SWAP")
	num1, num2 := 600, 3500
	fmt.Printf("[BEFORE] : num1=%v, num2=%v \n", num1, num2)
	num1 = num1 * num2 // 600 * 3 500 = 2 100 000
	num2 = num1 / num2 // 2 100 000 / 3 500 = 600
	num1 = num1 / num2 // 2 100 000 / 600 = 3 500
	fmt.Printf("[BEFORE] : num1=%v, num2=%v \n", num1, num2)

}
