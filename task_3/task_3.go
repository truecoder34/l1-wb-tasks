package main

import "fmt"

/*
	Дана последовательность чисел: 2,4,6,8,10.
	Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

type Answer struct {
	idx int
	val int
}

func calculatePows(val int, idx int, ch chan Answer) {
	var pow int = val * val
	var answer Answer = Answer{
		idx: idx,
		val: pow,
	}
	ch <- answer
}

func calculateSum(s []Answer, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v.val
	}
	// writes the sum to the go routines.
	c <- sum // send sum to c
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	// # 1 - calculate POWS on goroutines
	// channel to get result from goroutines to main goroutine
	ch := make(chan Answer)
	for i, elem := range arr {
		go calculatePows(elem, i, ch)
	}

	// # 2 - calculate POWS
	// будут добавляться в порядке вычисления. от запуска к запуску порядок будет разный
	answer_arr := make([]Answer, len(arr))
	sum := 0
	for i := range answer_arr {
		answer_arr[i] = <-ch
		fmt.Printf("[%d] - %d \n", answer_arr[i].idx, answer_arr[i].val)
		sum += answer_arr[i].val
	}
	fmt.Printf("Pows Results: %+v\n", answer_arr)

	fmt.Printf("Sum Result: %d\n", sum)

}
