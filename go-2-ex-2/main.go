package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	fibs[2] = fibs[0] + fibs[1]
	fibs[3] = fibs[1] + fibs[2]
	fibs[4] = fibs[2] + fibs[3]

	fibs = append(fibs, fibs[3]+fibs[4])                     // Index 5
	fibs = append(fibs, fibs[len(fibs)-1]+fibs[4])           // Index 6
	fibs = append(fibs, fibs[len(fibs)-1]+fibs[len(fibs)-2]) // Index 7
	fibs = append(fibs, fibs[len(fibs)-1]+fibs[len(fibs)-2]) // Index 8

	fmt.Println(fibs)
}
