package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	Name             FullName
	Birth            BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		Name: FullName{
			FirstName: "Steven",
			LastName:  "Mattmann",
		},
		Birth: BirthDate{
			Day:   30,
			Month: 10,
			Year:  2007,
		},
		NumberOfSiblings: 3,
		ZodiacSign:       '♏',
	}

	fmt.Println("Steckbrief:", me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
