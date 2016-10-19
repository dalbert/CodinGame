package main

import "fmt"
import "os"
import "bufio"
import "strconv"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	scanner.Split(bufio.ScanWords)

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)
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
	fmt.Println(m)
}
