package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
17.
Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	array := GenerateRandomArray(1)
	fmt.Printf("[Initial Array]: %v\n", array)
	array = QuickSort(array)
	fmt.Printf("[Sorted Array]: %v\n", array)

	what2Find, err := inputInt64("What number we are going to search for?")
	if err != nil {
		fmt.Printf("Failed to convert input to int64: %s \n", err)
	}

	res := BinarySearch(array, what2Find)
	if res == -1 {
		fmt.Println("Element was not found in array =( ")
	} else {
		fmt.Printf("Element has index %d \n", res)
	}
}

func BinarySearch(arr []int, searchElem int) int {
	length := len(arr)
	l, r := 0, length-1

	for l <= r {
		m := (l + r) / 2
		if arr[m] == searchElem { // тыкаем в центр
			return m
		} else if arr[m] < searchElem { // двигаем левую границу центрее
			l = m + 1
		} else if arr[m] > searchElem { // двигаем праавую левее
			r = m - 1
		}
	}
	return -1
}

// Generates a slice of size, size filled with random numbers
func GenerateRandomArray(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func inputInt64(line string) (int, error) {
	fmt.Println(line)
	reader := bufio.NewReader(os.Stdin)
	number, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	number = strings.TrimSuffix(number, "\n") // remove symbols like \t \r
	number = strings.TrimSuffix(number, "\r") // remove symbols like \t \r

	n, err := strconv.Atoi(number)
	if err != nil {
		fmt.Printf("Error on convert: %d of type %T \n", n, n)

	}

	return n, nil
}

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

	//fmt.Printf("[Array state WIP] : %v\n", array)

	// рекурсивно запускаем длЯ обеих частей левой и правой
	QuickSort(array[:l])
	QuickSort(array[l+1:])

	return array
}
