package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Name     string
	Students []Student
}

func main() {

	class1 := Class{
		Name: "INA23bL",
		Students: []Student{
			{FirstName: "Steven", LastName: "Mattmann"},
			{FirstName: "Nico", LastName: "Marcuard"},
			{FirstName: "Bob", LastName: "Fischer"},
		},
	}

	class2 := Class{
		Name: "INA23aL",
		Students: []Student{
			{FirstName: "Nils", LastName: "Utiger"},
			{FirstName: "Marvin", LastName: "Hegi"},
			{FirstName: "Domenic", LastName: "Troxler"},
		},
	}

	modules := map[int][]Class{
		346: {class1, class2},
		104: {class1},
		117: {class2},
	}

	fmt.Println("Klassen:")
	fmt.Printf("Klasse %s: %+v\n", class1.Name, class1.Students)
	fmt.Printf("Klasse %s: %+v\n", class2.Name, class2.Students)

	fmt.Println("\nModule:")
	for moduleNumber, attendingClasses := range modules {
		fmt.Printf("Modul %d wird besucht von:\n", moduleNumber)
		for _, class := range attendingClasses {
			fmt.Printf("- %s\n", class.Name)
		}
	}
}
