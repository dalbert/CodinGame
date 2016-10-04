package main

import "fmt"
import "os"
import "bufio"
import "strconv"

//import "strings"
//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	scanner.Split(bufio.ScanWords)

	// n: the number of temperatures to analyse
	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)
	fmt.Fprintln(os.Stderr, n)
	var m, md, td int = 0, 5527, 0
	for scanner.Scan() {
		if t, err := strconv.Atoi(scanner.Text()); err == nil {
			if t < 0 {
				td = t * -1
			} else {
				td = t
			}
			if td < md || (td == md && t > 0) {
				m = t
				md = td
			}
		}
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(m) // Write answer to stdout
}
