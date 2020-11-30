package main

import (
	"fmt"

	"github.com/muzfr7/polymorphism/living"
)

func main() {
	fmt.Println("Hello World...")

	animal1 := living.Animal{
		Name: "Dog",
		Age:  2,
	}
	animal2 := living.Animal{
		Name: "Cat",
		Age:  1,
	}

	person1 := living.Person{
		Name: "Muzafar Ali",
		Age:  32,
	}

	person2 := living.Person{
		Name: "Javed Ali",
		Age:  30,
	}

	linvingThings := []living.Speakable{person1, person2, animal1, animal2}
	nameOfLivingThings := findNamesOfLivng(linvingThings)
	fmt.Println(nameOfLivingThings)

	fmt.Println("End..")
}

func findNamesOfLivng(livingThings []living.Speakable) []string {
	var names []string
	for _, t := range livingThings {
		names = append(names, t.TellName())
	}

	return names
}
