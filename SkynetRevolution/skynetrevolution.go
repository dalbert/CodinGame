package main

import (
	"fmt"
	"os"
)

// Node butts
type Node struct {
	id             int
	distance       int
	isGateway      bool
	connectedNodes map[int]*Node
	gatewayCount   int
	beenVisited    bool
}

// NewNode butts
func NewNode(id int) *Node {
	n := Node{id: id, connectedNodes: make(map[int]*Node, 0)}
	return &n
}
func (node *Node) addConnection(connectedNode *Node) {
	node.connectedNodes[connectedNode.id] = connectedNode
}
func (node *Node) removeConnection(connectedNode *Node) {
	if connectedNode.isGateway {
		node.gatewayCount--
	}
	delete(node.connectedNodes, connectedNode.id)
}
func (node *Node) setIsGateway(is bool) {
	node.isGateway = is
	for _, connectedNode := range node.connectedNodes {
		connectedNode.gatewayCount++
	}
}
func (node Node) String() string {
	return fmt.Sprintf("{id: %d, gCount: %d, distance: %d, beenVisited: %v, isGateway: %v}\n", node.id, node.gatewayCount, node.distance, node.beenVisited, node.isGateway)
}

func main() {
	// N: the total number of nodes in the level, including the gateways
	// L: the number of links
	// E: the number of exit gateways
	var N, L, E int
	fmt.Scan(&N, &L, &E)

	topo := make(map[int]*Node, L)
	for i := 0; i < L; i++ {
		// N1: N1 and N2 defines a link between these nodes
		var N1, N2 int
		fmt.Scan(&N1, &N2)
		if _, present := topo[N1]; !present {
			topo[N1] = NewNode(N1)
		}
		if _, present := topo[N2]; !present {
			topo[N2] = NewNode(N2)
		}
		topo[N1].addConnection(topo[N2])
		topo[N2].addConnection(topo[N1])
	}

	for i := 0; i < E; i++ {
		// EI: the index of a gateway node
		var EI int
		fmt.Scan(&EI)
		node := topo[EI]
		node.setIsGateway(true)
	}

	for {
		// SI: The index of the node on which the Skynet agent is positioned this turn
		var SI int
		fmt.Scan(&SI)
		clearVisits(topo)
		N1, N2 := breadthSearch(topo[SI], topo)
		fmt.Fprintln(os.Stderr, topo)
		N1.removeConnection(N2)
		N2.removeConnection(N1)
		if len(topo[N1.id].connectedNodes) < 1 {
			delete(topo, N1.id)
		}
		if len(topo[N2.id].connectedNodes) < 1 {
			delete(topo, N2.id)
		}

		// Example: 0 1 are the indices of the nodes you wish to sever the link between
		fmt.Println(fmt.Sprintf("%d %d", N1.id, N2.id))
	}
}

func clearVisits(topo map[int]*Node) {
	for _, node := range topo {
		node.beenVisited = false
		node.distance = 0
	}
}

func breadthSearch(start *Node, topo map[int]*Node) (*Node, *Node) {
	// Skynet is one step away from a gateway, must cut that link immediately
	if start.gatewayCount > 0 {
		return selectEdgeToCut(start)
	}

	var nodeCount int
	queue := make(chan *Node, 100)
	start.beenVisited = true
	topo[start.id] = start
	queue <- topo[start.id]
	for current := range queue {
		nodeCount++
		for _, next := range topo[current.id].connectedNodes {
			if next.beenVisited {
				continue
			} else {
				fmt.Fprint(os.Stderr, fmt.Sprintf("c.id:%d\tc.gateCnt: %d\tc.dist: %d\tn.id: %d \t", current.id, current.gatewayCount, current.distance, next.id))
				next.beenVisited = true
				if current.gatewayCount > 0 {
					next.distance = current.distance
					fmt.Fprint(os.Stderr, next.distance)
				} else {
					next.distance = current.distance + 1
					fmt.Fprint(os.Stderr, next.distance)
				}
				fmt.Fprintln(os.Stderr, "")
				queue <- next
			}
		}
		if nodeCount >= len(topo) {
			close(queue)
		}
	}
	var selectedNode = start
	for _, node := range topo {
		//		fmt.Fprintln(os.Stderr, fmt.Sprintf("maxID %d, distance %d, maxGates: %d", maxGateID, minDistance, maxGateways))
		if node.gatewayCount < 1 {
			continue
		}
		if node.distance-node.gatewayCount < 0 {
			selectedNode = node
			break
		} else if node.gatewayCount > selectedNode.gatewayCount {
			selectedNode = node
		}
	}
	return selectEdgeToCut(selectedNode)
}

func selectEdgeToCut(selectedNode *Node) (*Node, *Node) {
	for _, node := range selectedNode.connectedNodes {
		if node.isGateway {
			return selectedNode, node
		}
	}
	return &Node{}, &Node{}
}
