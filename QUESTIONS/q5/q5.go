package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := struct{}{}
	fmt.Printf("%v", unsafe.Sizeof(s))
}
