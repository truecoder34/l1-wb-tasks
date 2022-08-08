package main

import (
	"fmt"
	"sync"
)

func calculatePowsWG(wg *sync.WaitGroup, val int, idx int) int {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()
	pow := val * val
	fmt.Printf("[%d] - %d \n", idx, pow)
	return pow
}

func main_wg() {
	fmt.Printf("========================= \n")
	fmt.Printf("Solution # 2: \n")
	fmt.Printf("========================= \n")
	arr := []int{2, 4, 6, 8, 10}
	// WaitGroup is used to wait for the program to finish goroutines.
	var wg sync.WaitGroup
	wg.Add(len(arr))

	//answer_arr := make([]Answer, len(arr))
	for idx, elem := range arr {
		go calculatePowsWG(&wg, elem, idx)
	}

	wg.Wait()
}

func main_wg_ch() {
	fmt.Printf("========================= \n")
	fmt.Printf("Solution # 3: \n")
	fmt.Printf("========================= \n")
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(len(arr))
	ch := make(chan Answer)

	//answer_arr := make([]Answer, len(arr))
	for idx, elem := range arr {
		go calculatePows(elem, idx, ch)
	}

	answer_arr := make([]Answer, len(arr))
	for i := range answer_arr {
		answer_arr[i] = <-ch
		fmt.Printf("[%d] - %d \n", answer_arr[i].idx, answer_arr[i].val)
	}

	wg.Wait()
}
