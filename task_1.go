package main

import "fmt"

// parent class
type Human struct {
	age  int
	name string
	sex  string
}

// interface for Human class
type Presenter interface {
	Present()
}

// method of Human class
func (h Human) Present() {
	fmt.Printf("Hello, my name is %s. I am %d years old.\n", h.name, h.age)
}

// child class
type Action struct {
	Human
	hobbies []string
}

func (a Action) ShowHobbies() {
	fmt.Printf("My hobbies are:  \n")
	for i, hobby := range a.hobbies {
		fmt.Printf("[%d] - %s \n", i, hobby)
	}
}

func main() {
	human := Human{
		age:  24,
		name: "Vlad",
		sex:  "male",
	}
	human.Present() // call from parent instance

	act := Action{
		Human: Human{
			age:  55,
			name: "VIN DIESEL",
			sex:  "male",
		},
		hobbies: []string{"race", "latinas"},
	}

	act.Present() // call from CHILD instance
	act.ShowHobbies()
	act.Human.Present() // call from CHILD instance
	//act.Human.ShowHobbies()
}
