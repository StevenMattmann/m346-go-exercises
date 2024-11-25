package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	//Lösung Chat GPT

	rand.Seed(time.Now().UnixNano())

	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		fmt.Println("Error creating eyes.txt:", err)
		return
	}
	defer eyesFile.Close()

	logFile, err := os.Create("dice.log")
	if err != nil {
		fmt.Println("Error creating dice.log:", err)
		return
	}
	defer logFile.Close()

	fmt.Fprintln(eyesFile, "The dice shows", eyes, "eyes")
	fmt.Fprintln(logFile, "The dice shows", eyes, "eyes")
	fmt.Fprintln(logFile, "The dice was rolled at", when)

	// Also print to standard output
	fmt.Println("The dice shows", eyes, "eyes")
	fmt.Println("The dice was rolled at", when)
}
