package main

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// N: the total number of nodes in the level, including the gateways
	// L: the number of links
	// E: the number of exit gateways
	var N, L, E int
	fmt.Scan(&N, &L, &E)

	topo := make([][]int, N)
	for i := 0; i < L; i++ {
		// N1: N1 and N2 defines a link between these nodes
		var N1, N2 int
		fmt.Scan(&N1, &N2)
		topo[N1] = append(topo[N1], N2)
		topo[N2] = append(topo[N2], N1)
	}

	gates := make([]bool, N)
	for i := 0; i < E; i++ {
		// EI: the index of a gateway node
		var EI int
		fmt.Scan(&EI)
		gates[EI] = true
	}

	for {
		// SI: The index of the node on which the Skynet agent is positioned this turn
		var SI int
		fmt.Scan(&SI)
		N1, N2 := breadthSearch(SI, topo, gates)
		topo = removeConnection(topo, N1, N2)

		// Example: 0 1 are the indices of the nodes you wish to sever the link between
		fmt.Println(fmt.Sprintf("%d %d", N1, N2))
	}
}

func breadthSearch(start int, topo [][]int, gates []bool) (int, int) {
	visited := map[int]bool{}
	queue := make(chan int, 100)
	queue <- start
	for next := range queue {
		visited[next] = true
		for i := 0; i < len(topo[next]); i++ {
			if gates[topo[next][i]] {
				return next, topo[next][i]
			}
			if !visited[topo[next][i]] {
				queue <- topo[next][i]
			}
		}
	}
	return 0, 0
}

func removeConnection(topo [][]int, N1 int, N2 int) [][]int {
	for i := 0; i < len(topo[N1]); i++ {
		if topo[N1][i] == N2 {
			topo[N1] = remove(topo[N1], i)
		}
	}
	for i := 0; i < len(topo[N2]); i++ {
		if topo[N2][i] == N1 {
			topo[N2] = remove(topo[N2], i)
		}
	}
	return topo
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
