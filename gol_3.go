// gol_1.go

/*

- Application implementing a one-dimensional simplified version of Conway's Game of life:

- There are no generations, cells just check and change (at a pace just slowed down
by the timer) in complete independence of each other, just based on the
sum of the values of their neighbours at a given moment.

- Basically A "cell" is repented by a value (0 or 1) hold in the array Value[]
- Each cell (i) lives in its own goroutine RunCell(i) and has the value Value[i]
- Each cell (i) evolves independently of the others, its evolution being determined only
by the state of its two closest neighbours (see function Rucell for the detail of the rules)

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Imax = 10
)

// The "Value" array, whose elements (cells) will evolve in parallel via gotoutines "RunCell(i)":
var Value [10]int

// Get the neighbours of cell "i" (which holds value: "Value(i)")
func GetNeighbours(I int) []int {
	var Nb []int
	Nb = make([]int, 0, Imax)
	for u := -1; u < 2; u++ {
		if (I+u >= 0) && (I+u < Imax) && (u != 0) {
			Nb = append(Nb, (I + u))
		}
	}
	return Nb
}

// Calculate sum of values of neighbours of cell "i"
func sumNeighbours(i int) int {
	s := 0
	for _, v := range GetNeighbours(i) {
		s = s + Value[v]
	}
	return s
}

// Sum of all cell values:
func sumAll() int {
	s := 0
	for _, v := range Value {
		s = s + v
	}
	return s
}

// -- funcion implementing a one-dimensional simplified version of
// the rules of Conway's Game of life:
func RunCell(i int, C chan bool) {
	for {
		if sumNeighbours(i) == 0 || sumNeighbours(i) == 2 {
			// Rule 1 : Starvation and over-crowding:
			Value[i] = 0
		} else {
			//Rule 2 : "Reproduction"
			Value[i] = 1
		}
		// -- Dump the state of the array "Value" after one cell made a change
		// or not depending of the state of its neighbours at that point in time :
		//fmt.Printf("V = %v\n", Value)
		for _, v := range Value {
			if v == 0 {
				print("[ ]")
			} else {
				print("[*]")
			}
		}
		fmt.Println()

		if sumAll() == 0 {
			print("\nStarvation\n")
			C <- true
		}

		if sumAll() == 10 {
			print("\nOvercrowded\n")
			C <- true
		}

		// pause to slown down display

		t := 1e7 * RandomWait()
		time.Sleep(time.Duration(t))

	}
	C <- true
}

func RandomWait() int {
	return rand.Intn(100)
}

// -- Entry point of the application:
func main() {
	fmt.Println("Starting")

	// Randomizing initial state of Value elements
	for i, _ := range Value {
		if rand.Float64() < 0.1 {
			Value[i] = 1
		} else {
			Value[i] = 0
		}
	}

	// -- Starting 10 parallel goroutines
	// -- Channel "C" is used merely as a lock (same as join() in other programming languages)
	C := make(chan bool)
	for i := 0; i < 10; i++ {
		go RunCell(i, C)
	}
	<-C
}
