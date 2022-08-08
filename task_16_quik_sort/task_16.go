package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
16.
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	slice := GenerateRandomArray(16)
	fmt.Println("\n--- [BEFORE SORT] --- \n\n", slice)
	// call sort
	slice = QuickSort(slice)
	fmt.Println("\n--- [AFTER SORT] --- \n\n", slice)
}

/*
	Average : O(n*logn)
	Worst: O(n^2)
	Memory: In-place
*/
func QuickSort(array []int) []int {
	length := len(array)
	if length < 2 {
		return array
	}
	// fix borders
	l, r := 0, length-1
	// choose pivot elem to sort based on it . Random or simply middle. опорный элемент
	p := rand.Int() % length

	array[p], array[r] = array[r], array[p] // initial swap of pivot and right. крайний правый меняем с опорным

	for i, _ := range array {
		// все элементы меньше опорного
		if array[i] < array[r] {
			array[l], array[i] = array[i], array[l] // двигаем левее
			l++
		}
	}
	array[l], array[r] = array[r], array[l] // опорный двигаем после последнего элемента который меньше опорного

	fmt.Printf("[Array state WIP] : %v\n", array)

	// рекурсивно запускаем длЯ обеих частей левой и правой
	QuickSort(array[:l])
	QuickSort(array[l+1:])

	return array
}

// Generates a slice of size, size filled with random numbers
func GenerateRandomArray(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
