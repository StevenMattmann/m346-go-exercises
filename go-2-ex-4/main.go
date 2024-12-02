package main

import "fmt"

// Student repräsentiert einen Schüler mit Vor- und Nachnamen.
type Student struct {
	FirstName string
	LastName  string
}

// Class repräsentiert eine Klasse, die aus mehreren Schülern besteht.
type Class struct {
	Name     string
	Students []Student
}

func main() {
	// Erstellung von zwei Klassen mit jeweils drei Schülern.
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

	// Erstellung von Modulen und Zuordnung der Klassen.
	modules := map[int][]Class{
		346: {class1, class2}, // Modul wird von beiden Klassen besucht.
		104: {class1},         // Modul wird nur von Klasse 1A besucht.
		117: {class2},         // Modul wird nur von Klasse 2B besucht.
	}

	// Ausgabe der Klassen und ihrer Schüler.
	fmt.Println("Klassen:")
	fmt.Printf("Klasse %s: %+v\n", class1.Name, class1.Students)
	fmt.Printf("Klasse %s: %+v\n", class2.Name, class2.Students)

	// Ausgabe der Module und der zugehörigen Klassen.
	fmt.Println("\nModule:")
	for moduleNumber, attendingClasses := range modules {
		fmt.Printf("Modul %d wird besucht von:\n", moduleNumber)
		for _, class := range attendingClasses {
			fmt.Printf("- %s\n", class.Name)
		}
	}
}
