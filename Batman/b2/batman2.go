package main

import (
	"fmt"
	"os"
	"math"
)

const warmer = "WARMER"
const colder = "COLDER"
const same = "SAME"
const unknown = "UNKNOWN"
const left = -1
const right = 1
const up = -1
const down = 1

func main() {
	var W, H int
	fmt.Scan(&W, &H)

	// N: maximum number of turns before game over.
	var N int
	fmt.Scan(&N)

	var X0, Y0 int
	var A, B Point
	fmt.Scan(&X0, &Y0)
	A = Point{x: X0, y: Y0}
	B = Point{x: X0, y: Y0}

	for {
		var bombDistance string
		fmt.Scan(&bombDistance)
		fmt.Fprintln(os.Stderr, bombDistance)
		if bombDistance == unknown {
			B.x, B.y = A.x, A.y
			A.x, A.y = W/2, H/2
			A.x, A.y = 1, 3
		} else if bombDistance != same {
			m, d := calcSlope(A, B), calcDistance(A, B)
			dX, dY := calcDeltas(m, d/2)
			fmt.Fprintln(os.Stderr, d)
			fmt.Fprintln(os.Stderr, m)
			fmt.Fprintln(os.Stderr, dX, dY)
			fmt.Fprintln(os.Stderr, fmt.Sprintf("%v --> %v", B, A))
		}

		fmt.Println(A.x, A.y) // Write action to stdout
	}
}

func calcDistance(a Point, b Point) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2))
}

func calcSlope(a Point, b Point) float64 {
	return float64(a.y - b.y)/float64(a.x - b.x)
}

func calcDeltas(slope float64, distance float64) (dX float64, dY float64) {
	angle := math.Atan(slope)
	dX = math.Cos(angle) * distance
	dY = math.Sin(angle) * distance
	return dX, dY
}

type Point struct {
	x, y int
}
