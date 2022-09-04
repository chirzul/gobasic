package main

import (
	"fmt"
	"gobasic/task"
)

func main() {
	fmt.Printf("%g0\n", task.RoundNumber(3.46))

	sequence := task.Sequence{Num: 40}
	fmt.Println(sequence.OddSequence())
	fmt.Println(sequence.EvenSequence())
	fmt.Println(sequence.PrimeSequence())
	fmt.Println(sequence.FiboSequence())

	var cube task.Calculate = &task.Cube{Side: 5}
	fmt.Println("Area :", cube.Area())
	fmt.Println("Perimeter :", cube.Perimeter())
	fmt.Println("Volume :", cube.Volume())
}
