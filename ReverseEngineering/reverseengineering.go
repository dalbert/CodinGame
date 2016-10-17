package main

import (
	"fmt"
	"os"
)

//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var firstInitInput int
	fmt.Scan(&firstInitInput)

	var secondInitInput int
	fmt.Scan(&secondInitInput)

	var thirdInitInput int
	fmt.Scan(&thirdInitInput)
	fmt.Fprintln(os.Stderr, fmt.Sprintf("INIT:\t%v %v %v", firstInitInput, secondInitInput, thirdInitInput))
	for {
		var firstInput string
		fmt.Scan(&firstInput)

		var secondInput string
		fmt.Scan(&secondInput)

		var thirdInput string
		fmt.Scan(&thirdInput)

		var fourthInput string
		fmt.Scan(&fourthInput)
		fmt.Fprintln(os.Stderr, fmt.Sprintf("1-4:\t%v %v %v %v", firstInput, secondInput, thirdInput, fourthInput))
		for i := 0; i < thirdInitInput; i++ {
			var fifthInput, sixthInput int
			fmt.Scan(&fifthInput, &sixthInput)
			fmt.Fprintln(os.Stderr, fmt.Sprintf("5-6.%d:\t%v %v", i, fifthInput, sixthInput))
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		//fmt.Println("A, B, C, D or E")// Write action to stdout
		fmt.Println("A") // Write action to stdout
	}
}
