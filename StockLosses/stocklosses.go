package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "sort"

type pair struct {
  time, value int
}
type byValue []pair

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var n int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&n)
    var values []pair
    fmt.Fprintln(os.Stderr, values)
    scanner.Scan()
    inputs := strings.Split(scanner.Text()," ")
    fmt.Fprintln(os.Stderr, inputs, n)
    for i := 0; i < n; i++ {
        v,_ := strconv.ParseInt(inputs[i],10,32)
        values = append(values, pair{time: i, value: int(v)})
    }
    sort.Sort(byValue(values))
    fmt.Fprintln(os.Stderr, values)

    var maxDrop int
    for i, j := 0, len(values)-1; i < len(values) && j > 0; i, j = i+1, j-1 {
      fmt.Fprintln(os.Stderr, values[i].time, values[i].value, "\t", values[j].time, values[j].value)
      if maxDrop < values[i].value 
    }
    fmt.Println("answer")
}

func (v byValue) Len() int {
  return len(v)
}
func (v byValue) Swap(i, j int) {
  v[i], v[j] = v[j], v[i]
}
func (v byValue) Less(i, j int) bool {
  return v[i].value < v[j].value
}
