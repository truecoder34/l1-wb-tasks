package main

import "fmt"

/*
11.
Реализовать пересечение двух неупорядоченных множеств.
*/

/*
	Реализация set через map, где ключ = значение добавляемое в сет, а значение = bool true
	можно заменить bool на пустую структуру , тк пусатя структура не используем память
	--------------------------------------------------
	type void struct{}
	var member void
	set := make(map[string]void)
	--------------------------------------------------
	set["Foo"] = true            // Добавить
	for k := range set {         // Пройти в цикле
		fmt.Println(k)
	}
	delete(set, "Foo")    // Удалить
	size := len(set)      // Размер
	exists := set["Foo"]  // Наличие элемента
*/

func main() {

	set1 := InitializeSet([]string{"30", "70", "10", "25", "5"})
	set2 := InitializeSet([]string{"100", "40", "5", "30", "10"})
	set3 := InitializeSet([]string{"30", "0", "30", "1", "2", "0"})

	intersection := IntersectSets(set1, set2)
	fmt.Printf("Intersected set1 and set2 %v \n", intersection)
	fmt.Printf("Intersected set3 and set2 %v \n", IntersectSets(set3, set2))

}

func InitializeSet(values []string) map[string]bool {
	set := make(map[string]bool)
	for _, val := range values {
		set[val] = true
	}
	fmt.Printf("[SET INITIALIZED] : %v \n", set)
	return set
}

/*
	Intersect __unordered__ sets ; O(n2) complexity
*/
func IntersectSets(set1 map[string]bool, set2 map[string]bool) map[string]bool {
	counterBuffer := make(map[string]int)
	intersection := make(map[string]bool, 0)
	for v := range set1 { // 1 - mark each elem in set 1  .
		counterBuffer[v] = 1
	}
	for value := range set2 { // 2 - iter through set 2 and find pair in set 1
		times := counterBuffer[value]
		if times == 1 {
			intersection[value] = true
		}
	}
	return intersection
}

/*
TODO
Set Difference
*/
