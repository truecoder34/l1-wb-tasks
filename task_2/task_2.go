package main

import "fmt"

/*
конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
и выведет их квадраты в stdout
*/

type Answer struct {
	idx int
	val int
}

// return value via channel
func calculatePows(val int, idx int, ch chan Answer) {
	var pow int = val * val
	var answer Answer = Answer{
		idx: idx,
		val: pow,
	}
	ch <- answer
}

func main_goroutines() {
	fmt.Printf("========================= \n")
	fmt.Printf("Solution # 1: \n")
	fmt.Printf("========================= \n")
	arr := []int{2, 4, 6, 8, 10}

	// channel to get result from goroutines to main goroutine
	ch := make(chan Answer)
	for i, elem := range arr {
		go calculatePows(elem, i, ch)
	}

	// будут добавляться в порядке вычисления. от запуска к запуску порядок будет разный
	answer_arr := make([]Answer, len(arr))
	for i := range answer_arr {
		answer_arr[i] = <-ch
		fmt.Printf("[%d] - %d \n", answer_arr[i].idx, answer_arr[i].val)
	}

	fmt.Printf("Results: %+v\n", answer_arr)
}

func main() {
	main_goroutines()
	main_wg()

	main_wg_ch()
}
