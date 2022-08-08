package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

/*
9.
https://go.dev/blog/pipelines
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	x, err := inputInt("Input array length :")
	if err != nil {
		fmt.Printf("Failed to convert input to int64: %s \n", err)
	}
	arrayInts := InitArray(x) // get initial array of ints

	firstChan := WriteInt2Channel(arrayInts) // return output channel  1st channel
	powChan := Number2Square(firstChan)

	for i := range powChan {
		fmt.Println(i)
	}

}

/*
1 part - sender :
преобразует список целых чисел в канал, который выдает целые числа из списка.
Функция запускает горутину,
которая отправляет целые числа по каналу и закрывает канал, когда все значения отправлены:

[IN] : array
[OUT] : channel
*/
func WriteInt2Channel(arr []int) <-chan int {
	out := make(chan int) // channel where to send
	go func() {
		for _, val := range arr {
			out <- val
		}
		close(out)
	}()
	return out
}

/*
2 part - receiver :
получает целые числа из канала и возвращает канал,
который выдает квадрат каждого полученного целого числа. После того,
как входящий канал закрыт и этот этап отправил все значения вниз по течению,
 он закрывает исходящий канал:

 [IN] : channel base
 [OUT] : channel square
*/
func Number2Square(valIn <-chan int) <-chan int {
	outSq := make(chan int)
	go func() {
		for num := range valIn {
			outSq <- num * num
		}
		close(outSq)
	}()
	return outSq
}

/*
	Create random array of integers of specified length
*/
func InitArray(arrayLen int) []int {
	array := make([]int, arrayLen)
	for i := 0; i < arrayLen; i++ {
		array[i] = rand.Intn(150)
	}
	fmt.Printf("Generated array : %v \n", array)
	return array
}

/*
Process the input
*/
func inputInt(line string) (int, error) {
	fmt.Println(line)
	reader := bufio.NewReader(os.Stdin)
	number, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	number = strings.TrimSuffix(number, "\n") // remove symbols like \t \r
	number = strings.TrimSuffix(number, "\r") // remove symbols like \t \r

	n, err := strconv.Atoi(number) // strconv.ParseInt(number, 10, 32)
	if err != nil {
		fmt.Printf("Error on convert: %d of type %T \n", n, n)

	}
	return n, nil
}
