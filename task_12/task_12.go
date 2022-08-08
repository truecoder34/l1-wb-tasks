package main

import "fmt"

/*
12.
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	set1 := InitializeSet([]string{"30", "70", "10", "25", "5"})
	fmt.Printf("Result : %v \n", set1)
	fmt.Printf("Result : %v \n", InitializeSet([]string{"cat", "cat", "dog", "cat", "tree"}))
}

func InitializeSet(values []string) map[string]bool {
	set := make(map[string]bool)
	fmt.Printf("[SOURCE VALUES]: %v; Len: %d\n", values, len(values))
	for _, val := range values {
		set[val] = true
	}
	fmt.Printf("[SET INITIALIZED]: %v; Len: %d \n", set, len(set))
	return set
}
