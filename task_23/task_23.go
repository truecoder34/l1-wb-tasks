package main

import "fmt"

/*
23. Удалить i-ый элемент из слайса.
*/

func main() {
	slice_example := []string{"one", "two", "three", "four", "five", "six", "seven", "eight"}
	fmt.Printf("Slice to process : %v\n", slice_example)
	idx := 2                                        // remove element with index 2
	for i := 0; i < len(slice_example)-idx-1; i++ { // start swap from target index:
		// swap next elem with current index element
		fmt.Printf("[%d]: swap of %s to %s\n", i, slice_example[i+idx+1], slice_example[i+idx])
		slice_example[i+idx] = slice_example[i+idx+1]
		fmt.Printf("[%d] State: %v\n", i, slice_example)
	}
	slice_example = slice_example[:len(slice_example)-1] // drop last element
	fmt.Printf("[RESULT]: %v\n", slice_example)
}
