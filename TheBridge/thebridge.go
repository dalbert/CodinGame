package main

import "fmt"

//import "os"

type bike struct {
	x, y     int
	isActive bool
}

func main() {
	// M: the amount of motorbikes to control
	var M int
	fmt.Scan(&M)

	// V: the minimum amount of motorbikes that must survive
	var V int
	fmt.Scan(&V)

	// L0: L0 to L3 are lanes of the road. A dot character . represents a safe space, a zero 0 represents a hole in the road.
	var L0 string
	fmt.Scan(&L0)

	var L1 string
	fmt.Scan(&L1)

	var L2 string
	fmt.Scan(&L2)

	var L3 string
	fmt.Scan(&L3)

	var theBikes []bike
	for {
		// S: the motorbikes' speed
		var S int
		fmt.Scan(&S)

		for i := 0; i < M; i++ {
			// X: x coordinate of the motorbike
			// Y: y coordinate of the motorbike
			// A: indicates whether the motorbike is activated "1" or detroyed "0"
			var X, Y, A int
			fmt.Scan(&X, &Y, &A)
			act := false
			if A == 1 {
				act = true
			}
			theBikes = append(theBikes, bike{x: X, y: Y, isActive: act})
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")

		// A single line containing one of 6 keywords: SPEED, SLOW, JUMP, WAIT, UP, DOWN.
		fmt.Println("SPEED")
	}
}
