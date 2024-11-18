package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!

	var firstName string = "Max"
	var lastName string = "Mustermann"
	var dayOfBirth int = 30
	var monthOfBirth int = 10
	var yearOfBirth int = 1999
	var numberOfSiblings int = 3
	var heightInMeters float64 = 1.82
	var zodiacSign rune = '\u264F'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
