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

func main() {
	var W, H int
	fmt.Scan(&W, &H)
	left, top, right, bottom := 0, 0, W-1, H-1
	X, Y := false, true // which dimension are we bifurcating

	// N: maximum number of turns before game over.
	var N int
	fmt.Scan(&N)

	var X0, Y0, midline int
	var A, B Point
	fmt.Scan(&X0, &Y0)
	A = Point{x: X0, y: Y0}

	for {
		var bombDistance string
		fmt.Scan(&bombDistance)
		fmt.Fprintln(os.Stderr, bombDistance)
		switch bombDistance {
		case warmer:
			if A.x != B.x { // prior turn was X, so update X bounds
				midline = int(Round(float64(A.x + B.x) / 2, .5, 0))
				if A.x > B.x {
					left = midline
				} else {
					right = midline
				}
			}
			if A.y != B.y { // prior turn was Y, so update Y bounds
				midline = int(Round(float64(A.y + B.y) / 2, .5, 0))
				fmt.Fprintln(os.Stderr, midline)
				if A.y > B.y {
					top = midline
				} else {
					bottom = midline-1
				}
			}
		case colder:
			if A.x != B.x { // prior turn was X, so update X bounds
				midline = int(Round(float64(A.x + B.x) / 2, .5, 0))
				if A.x > B.x {
					right = midline
				} else {
					left = midline
				}
			}
			if A.y != B.y { // prior turn was Y, so update Y bounds
				midline = int(Round(float64(A.y + B.y) / 2, .5, 0))
				if A.y > B.y {
					bottom = midline
				} else {
					top = midline
				}
			}
		case same:
		case unknown:

		}
		if B.x == A.x && B.y == A.y { // made no change, skip this turn and try the other dimension
			X, Y = swapDimension(X, Y)
			fmt.Fprintln(os.Stderr, X, Y)
			continue
		}
		B.x, B.y = A.x, A.y
		if Y {
			A.y = int(Round(float64(bottom + top) / 2, .5, 0))
		} else {
			A.x = int(Round(float64(left + right) / 2, .5, 0))
		}
		fmt.Fprintln(os.Stderr, left, right, top, bottom, X, Y, Round(float64(left + right) / 2, .5, 0), Round(float64(bottom + top) / 2, .5, 0))
		if left == right {
			X, Y = false, true
		} else if top == bottom {
			X, Y = true, false
		} else {
			X, Y = swapDimension(X, Y)
		}
		fmt.Println(A.x, A.y) // Write action to stdout
	}
}

func swapDimension(X bool, Y bool) (bool, bool) {
	if X {
		return false, true
	}
	return true, false
}

func calcDistance(a Point, b Point) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2))
}

func calcSlope(a Point, b Point) float64 {
	m := float64(a.y - b.y)/float64(a.x - b.x)
	if math.IsInf(m, -1) {return 0}
	return m
}

func calcPerpSlope(a Point, b Point) float64{
	return 1 / (float64(a.y - b.y)/float64(a.x - b.x)) * -1
}

func calcYOffset(midPoint Point, slope float64) (b float64) {
	return float64(midPoint.y) - (slope * float64(midPoint.x))
}

func calcDeltas(slope float64, distance float64) (dX float64, dY float64) {
	angle := math.Atan(slope)
	dX = math.Cos(angle) * distance
	dY = math.Sin(angle) * distance
	return dX, dY
}

func Round(val float64, roundOn float64, places int ) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	_div := math.Copysign(div, val)
	_roundOn := math.Copysign(roundOn, val)
	if _div >= _roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

type Point struct {
	x, y int
}
