package main

import "fmt"

//import "os"

// Pair - pair of colors dropping into the game
type Pair struct {
	a, b int
}

func main() {
	for {
		var upcoming [8]Pair
		for i := 0; i < 8; i++ {
			// colorA: color of the first block
			// colorB: color of the attached block
			var colorA, colorB int
			fmt.Scan(&colorA, &colorB)
			upcoming[i] = Pair{a: colorA, b: colorB}
		}

		var score1 int
		fmt.Scan(&score1)

		for i := 0; i < 12; i++ {
			var row string
			fmt.Scan(&row)
		}
		var score2 int
		fmt.Scan(&score2)

		for i := 0; i < 12; i++ {
			// row: One line of the map ('.' = empty, '0' = skull block, '1' to '5' = colored block)
			var row string
			fmt.Scan(&row)
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println(upcoming[0].a, "BUTT SOUP") // "x": the column in which to drop your blocks
	}
}
