package main

import "fmt"
import "os"
import "bufio"

//import "strings"
//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	MESSAGE := scanner.Text()
	fmt.Fprintln(os.Stderr, toBinary(MESSAGE))
	binMess := toBinary(MESSAGE)

	var priorC, nextC rune = rune(binMess[0]), rune(binMess[0])
	var sequenceCounter int
	for _, nextC = range binMess {
		if nextC == priorC {
			sequenceCounter++
		} else {
			if priorC == '1' { // 1's
				fmt.Print("0")
			} else { // 0's
				fmt.Print("00")
			}
			fmt.Print(" ")
			for i := 0; i < sequenceCounter; i++ {
				fmt.Print("0")
			}
			sequenceCounter = 1
			priorC = nextC
			fmt.Print(" ")
		}
	}
	if priorC == '1' { // 1's
		fmt.Print("0")
	} else { // 0's
		fmt.Print("00")
	}
	fmt.Print(" ")
	for i := 0; i < sequenceCounter; i++ {
		fmt.Print("0")
	}
}

func toBinary(s string) string {
	var binStr string
	for _, c := range s {
		binStr = fmt.Sprintf("%s%.7b", binStr, c)
	}
	return binStr
}
