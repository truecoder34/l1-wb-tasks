package main

import (
	"fmt"
	"strings"
)

/*
26.	Разработать программу, которая проверяет,
что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	test1 := "abcd"
	test2 := "abCdefAaf"
	test3 := "aabcd"
	test4 := "qwertyuiop"

	fmt.Printf("Res : %d\n", checkIfStringUnique(test1))
	fmt.Printf("Res : %d\n", checkIfStringUnique(test2))
	fmt.Printf("Res : %d\n", checkIfStringUnique(test3))
	fmt.Printf("Res : %d\n", checkIfStringUnique(test4))
}

/*
	[INPUT] string
	[OUTPUT] true - unique , false - not unique
	capitalized independent
*/
func checkIfStringUnique(line2check string) bool {
	res := true
	// 0 - string to lower
	// 1 - string to rune
	runeString := []rune(strings.ToLower(line2check))

	// add runes to set
	setOfRunes := make(map[rune]bool)
	for _, r := range runeString {
		if _, ok := setOfRunes[r]; ok {
			fmt.Printf("Element [%s] is already in string \n", string(r))
			res = false
			return res
		} else {
			setOfRunes[r] = true
		}

	}

	return res
}
