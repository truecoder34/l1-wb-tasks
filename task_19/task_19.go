package main

import "fmt"

/*
19.
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func main() {
	str1 := "text"
	str2 := "✁ ✂ ✃ ✄"
	str3 := "✓-✔-✕-✖-✗-✘-✙"

	fmt.Printf("Was : %s, Now : %s\n", str1, ReverseString(str1))
	fmt.Printf("Was : %s, Now : %s\n", str2, ReverseString(str2))
	fmt.Printf("Was : %s, Now : %s\n", str3, ReverseString(str3))
}

/*
	Reverse string
	get string
	convert to  array of bytes
	return string
*/
func ReverseString(text string) string {
	//str2bytes := []byte(text)
	str2rune := []rune(text)
	// with 2 pointers solution
	l, r := 0, len(str2rune)-1
	for i := 0; i < len(str2rune)/2; i++ {
		str2rune[l], str2rune[r] = str2rune[r], str2rune[l]
		l += 1
		r -= 1
	}
	return string(str2rune)
}
