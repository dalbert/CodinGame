package main

import (
	"fmt"
	"os"
)

func main() {
	// W: width of the building.
	// H: height of the building.
	var W, H int
	fmt.Scan(&W, &H)
	var TOP, BOTTOM, LEFT, RIGHT int = 0, H - 1, 0, W - 1

	// N: maximum number of turns before game over.
	var N int
	fmt.Scan(&N)

	var X0, Y0 int
	fmt.Scan(&X0, &Y0)
	for {
		// bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
		var bombDir string
		fmt.Scan(&bombDir)
		fmt.Fprintln(os.Stderr, bombDir)

		switch bombDir {
		case "U":
			BOTTOM = Y0
			Y0 = min((BOTTOM+TOP)/2, Y0+1)
		case "UR":
			BOTTOM = Y0
			Y0 = min((BOTTOM+TOP)/2, Y0+1)
			fallthrough
		case "R":
			LEFT = X0
			X0 = max((LEFT+RIGHT)/2, X0+1)
		case "DR":
			LEFT = X0
			X0 = max((LEFT+RIGHT)/2, X0+1)
			fallthrough
		case "D":
			TOP = Y0
			Y0 = max((BOTTOM+TOP)/2, Y0+1)
		case "DL":
			TOP = Y0
			Y0 = max((BOTTOM+TOP)/2, Y0+1)
			fallthrough
		case "L":
			RIGHT = X0
			X0 = min((LEFT+RIGHT)/2, X0+1)
		case "UL":
			BOTTOM = Y0
			Y0 = min((BOTTOM+TOP)/2, Y0+1)
			RIGHT = X0
			X0 = min((LEFT+RIGHT)/2, X0+1)
		}
		fmt.Println(X0, Y0)
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
