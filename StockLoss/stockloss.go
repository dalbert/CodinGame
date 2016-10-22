package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)
	var values []int64
	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	//	fmt.Fprintln(os.Stderr, inputs, n)
	for i := 0; i < n; i++ {
		v, _ := strconv.ParseInt(inputs[i], 10, 32)
		values = append(values, v)
	}

	var diff, maxDrop int64
	for i := 0; i < len(values); i++ {
		if i > 0 && (values[i] < values[i-1] || values[i] < maxDrop) {
			continue
		}
		for j := i + 1; j < len(values); j++ {
			diff = values[i] - values[j]
			//			fmt.Fprintln(os.Stderr, values[i].time, values[i].value, "\t", values[j].time, values[j].value, "\t", maxDrop, diff)
			if maxDrop < diff { // looking for the most negative difference
				maxDrop = diff
			}
		}
	}
	fmt.Println(maxDrop * -1)
}
