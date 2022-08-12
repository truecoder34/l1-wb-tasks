package main

import "fmt"

/*
10.	Дана последовательность температурных колебаний:
-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	temperaturesRow := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	temperaturesRow2 := []float64{-22.0, -15.0, -10.0, -9.0, -8.0, -7.0, -6.0, -5.0, -4.0, -3.0, 0.0, 2.0, 5.0, 8.0, 12.0}

	splitByDozens := split2Rows(temperaturesRow, 10)
	splitByDozens2 := split2Rows(temperaturesRow, 5)
	splitByDozens3 := split2Rows(temperaturesRow, 2)
	splitByDozens4 := split2Rows(temperaturesRow, 20)

	fmt.Printf("Result : %v\n", splitByDozens)
	fmt.Printf("Result : %v\n", splitByDozens2)
	fmt.Printf("Result : %v\n", splitByDozens3)
	fmt.Printf("Result : %v\n", splitByDozens4)
	fmt.Println("=================================")

	fmt.Printf("Result : %v\n", split2Rows(temperaturesRow2, 10))
	fmt.Printf("Result : %v\n", split2Rows(temperaturesRow2, 5))
	fmt.Printf("Result : %v\n", split2Rows(temperaturesRow2, 2))
	fmt.Printf("Result : %v\n", split2Rows(temperaturesRow2, 20))

}

/*
	[INPUT] get array of temperatures
	[INPUT] split step
	[OUTPUT] map of map arrays
*/
func split2Rows(temperatures []float64, step int) map[int][]float64 {
	result := make(map[int][]float64)

	for _, val := range temperatures {
		key := int(val) / step * step // частное от шага умнажаем на шаг чтобы получить ключ (группу) к которой отнесем число
		result[key] = append(result[key], val)
	}
	return result
}
