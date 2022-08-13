package main

import "fmt"

func update(p *int) {
	b := 2
	p = &b // в указатель записываем новую переменную
	fmt.Printf("update(). p это %T = %v\n", p, *p)
}

/*
https://golangify.com/pointers

(&) адрес в памяти
(*) обращание к значению в памяти

Указатели хранят адреса памяти.
*/

func main() {
	var (
		a = 1
		p = &a // p это указатель на a . на  адресс памаяти где лежит a равное 1
		// p - указатель типа *int
	)
	fmt.Println(*p) // выводим  значение по аддресу a 1
	fmt.Printf("main(). p это %T = %v \n", p, *p)
	update(p)
	fmt.Println(*p) // выведет 1 тк в этой области видимости p не менялось
}
