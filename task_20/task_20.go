package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
20.	Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	line1 := "Then we copy a three-element string into it"
	line2 := "result := bytes.Index(values, []byte(\"g\"))"
	line3 := "Here-we-import-the-\"bytes\"-package-at-the-top-(in-the-import-statement)"

	fmt.Printf("[Was] : %s, [Now] : %s\n", line1, ReverseLine(line1, " "))
	fmt.Printf("[Was] : %s, [Now] : %s\n", line2, ReverseLine(line2, " "))
	fmt.Printf("[Was] : %s, [Now] : %s\n", line3, ReverseLine(line3, "-"))

	fmt.Printf("[Was] : %s, [Now] : %s\n", line1, ReverseLineRegexp(line1, " "))
	fmt.Printf("[Was] : %s, [Now] : %s\n", line2, ReverseLineRegexp(line2, " "))
	fmt.Printf("[Was] : %s, [Now] : %s\n", line3, ReverseLineRegexp(line3, "-"))

}

/*
	Reverse based on split and concat
*/
func ReverseLine(line string, delimiter string) string {
	words := strings.Split(line, delimiter)
	l, r := 0, len(words)-1
	for i := 0; i < len(words)/2; i++ {
		words[l], words[r] = words[r], words[l]
		l += 1
		r -= 1
	}
	return strings.Join(words, delimiter)
}

/*
	Reverse based on Regexp
*/
func ReverseLineRegexp(line string, delimiter string) string {
	exp := regexp.MustCompile(delimiter)
	splitByDelim := exp.Split(line, -1)
	l, r := 0, len(splitByDelim)-1
	for i := 0; i < len(splitByDelim)/2; i++ {
		splitByDelim[l], splitByDelim[r] = splitByDelim[r], splitByDelim[l]
		l += 1
		r -= 1
	}
	return strings.Join(splitByDelim, delimiter)
}
