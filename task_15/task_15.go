package main

import (
	"fmt"
	"math/rand"
)

/*
15.
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

1 - размер слайса не контролируется справа. можно вывалиться в аут оф индекс и панику поймать
2 - не реализована функиця создания строки
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // 1024
	fmt.Printf("[Generated string] : %s\n", v)
	fmt.Printf("[Length of string] : %d\n", len(v))
	//v := createHugeString(1 << 11) // 2048
	//v := createHugeString(1 << 9) // 512
	//v := createHugeString(1 << 8) // 256
	rightBorder := 10
	switch {
	case rightBorder <= len(v)-1:
		justString = v[:rightBorder] // take slice from start to 100
	default:
		fmt.Println("Border Error. Reduce left index. ")
	}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

/*
	Implementation of function to extend string size and randomly create
*/
func createHugeString(size int) string {
	bigByteStr := make([]byte, size)
	for idx := range bigByteStr {
		bigByteStr[idx] = byte(charset[rand.Int63n(int64(len(charset)))]) // generate string from alphabet
	}
	bigStr := string(bigByteStr)
	return bigStr
}

func main() {
	someFunc()
	println(justString)
}
