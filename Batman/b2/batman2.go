package main

import "fmt"
import "os"

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

	var X0, Y0, X1, Y1 int
	fmt.Scan(&X0, &Y0)

	for {
		var bombDistance string
		fmt.Scan(&bombDistance)
		fmt.Fprintln(os.Stderr, bombDistance)
		if bombDistance == unknown {
			X0, Y0 = W/2, H/2
		} else if bombDistance == warmer {
			fmt.Fprintln(os.Stderr, X1, Y1)
		}

		fmt.Println(X0, Y0) // Write action to stdout
		X1, Y1 = X0, Y0
	}
}
