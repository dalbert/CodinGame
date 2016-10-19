package main

import (
	"fmt"
	"sort"
)

//import "os"

func main() {
	var N int
	fmt.Scan(&N)

	var strengths sort.IntSlice
	var minDiff, priorStrength = 100000000, -1000000
	var newDiff int

	for i := 0; i < N; i++ {
		var Pi int
		fmt.Scan(&Pi)
		strengths = append(strengths, Pi)
	}
	strengths.Sort()
	for _, s := range strengths {
		newDiff = diff(s, priorStrength)
		if newDiff < minDiff {
			minDiff = newDiff
		}
		priorStrength = s
	}
	fmt.Println(minDiff)
}

// always returns a non-negative value
func diff(a int, b int) (diff int) {
	if a > b {
		return a - b
	}
	return b - a
}
