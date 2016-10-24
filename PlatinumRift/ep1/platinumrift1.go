package main

import "fmt"

func main() {
	// playerCount: the amount of players (2 to 4)
	// myId: my player ID (0, 1, 2 or 3)
	// zoneCount: the amount of zones on the map
	// linkCount: the amount of links between all zones
	var playerCount, myID, zoneCount, linkCount int
	fmt.Scan(&playerCount, &myID, &zoneCount, &linkCount)
	zones := make(map[int]*zone, zoneCount)
	zonesByPlatinum := make(map[int][]int, 7)
	for i := 0; i < zoneCount; i++ {
		// zoneId: this zone's ID (between 0 and zoneCount-1)
		// platinumSource: the amount of Platinum this zone can provide per game turn
		var zoneID, platinumSource int
		fmt.Scan(&zoneID, &platinumSource)
		zones[zoneID] = &zone{id: zoneID, platinum: platinumSource}
		zonesByPlatinum[platinumSource] = append(zonesByPlatinum[platinumSource], zoneID)
	}
	for i := 0; i < linkCount; i++ {
		var zone1, zone2 int
		fmt.Scan(&zone1, &zone2)
		zones[zone1].links = append(zones[zone1].links, zone2)
		zones[zone2].links = append(zones[zone2].links, zone1)
	}
	// game loop
	for {
		// platinum: my available Platinum
		var myPlatinum int
		zonesByPlayer := make(map[int][]int, 4)

		fmt.Scan(&myPlatinum)

		var neutralZoneCount int
		for i := 0; i < zoneCount; i++ {
			z := zones[i]
			var myPodCount int
			var enemyPods [3]int
			// zId: this zone's ID
			// ownerId: the player who owns this zone (-1 otherwise)
			// podsP0: player 0's PODs on this zone
			// podsP1: player 1's PODs on this zone
			// podsP2: player 2's PODs on this zone (always 0 for a two player game)
			// podsP3: player 3's PODs on this zone (always 0 for a two or three player game)
			var zID, ownerID, podsP0, podsP1, podsP2, podsP3 int
			fmt.Scan(&zID, &ownerID, &podsP0, &podsP1, &podsP2, &podsP3)
			if ownerID == -1 {
				neutralZoneCount++
			}
			zonesByPlayer[ownerID] = append(zonesByPlayer[ownerID], zID)
			switch myID {
			case -1:
				myPodCount, enemyPods = 0, [3]int{0, 0, 0}
			case 0:
				myPodCount, enemyPods = podsP0, [3]int{podsP1, podsP2, podsP3}
			case 1:
				myPodCount, enemyPods = podsP1, [3]int{podsP0, podsP2, podsP3}
			case 2:
				myPodCount, enemyPods = podsP2, [3]int{podsP0, podsP1, podsP3}
			case 3:
				myPodCount, enemyPods = podsP3, [3]int{podsP0, podsP1, podsP2}
			}
			z.ownerID, z.myPods, z.enemyPods = ownerID, myPodCount, enemyPods
		}

		//		fmt.Fprintln(os.Stderr, zones)
		//		fmt.Fprintln(os.Stderr, myPlatinum)
		moves, buys := moveList{}, []buy{}
		podBudget := myPlatinum / podCost
		if neutralZoneCount > 0 {
			buys = claimNeutralZones(&podBudget, zones, zonesByPlatinum, myID)
		}
		buys = append(buys, reinforceOwnedZones(&podBudget, zones, zonesByPlayer[myID])...)

		moveCommand := moves.String()
		fmt.Println(moveCommand)
		buyCommand := buys.String()
		fmt.Println(buyCommand)

	}
}
func reinforceOwnedZones(podBudget *int, zones map[int]*zone, myZones []int) (buys []buy) {
	for _, zID := range myZones {
		if *podBudget > 0 {
			buys = append(buys, buy{quantity: 1, zID: zID})
		} else {
			return
		}
	}
	return
}
func claimNeutralZones(podBudget *int, zones map[int]*zone, zonesByPlatinum map[int][]int, myID int) (buys []buy) {
	for i := 6; i > -1; i-- {
		if *podBudget < 1 {
			return
		}
		for _, zID := range zonesByPlatinum[i] {
			z := zones[zID]
			need := 1
			if z.ownerID == myID || z.ownerID != -1 {
				continue
			}
			if z.EnemyCount() > z.myPods { // fighting for contested neutral zones (maybe, not sure how this simultaneous biz works)
				need = z.EnemyCount() - z.myPods + 1
			}
			if *podBudget > need { // can I afford to take this zone?
				buys = append(buys, buy{quantity: need, zID: zID})
				*podBudget = *podBudget - need
			}
			if *podBudget < 1 {
				return
			}
		}
	}
	return
}

type zone struct {
	id        int
	platinum  int
	links     []int
	ownerID   int
	myPods    int
	enemyPods [3]int
}

func (z *zone) EnemyCount() int {
	return z.enemyPods[0] + z.enemyPods[1] + z.enemyPods[2]
}
func (z *zone) String() string {
	return fmt.Sprintf("%d{i: %d, p: %d, e: %d}", z.ownerID, z.id, z.platinum, z.enemyPods[0]+z.enemyPods[1]+z.enemyPods[2])
}

type buy struct {
	quantity, zID int
}
type buyList []buy

func (b buy) String() string {
	return fmt.Sprintf("%d %d", b.quantity, b.zID)
}
func (bl []buy) String() string {
	return fmt.Sprint(bl...)
}

type move struct {
	quantity, originZID, destinationZID int
}
type moveList []move

func (m move) String() string {
	return fmt.Sprintf("%d %d %d", m.quantity, m.originZID, m.destinationZID)
}
func (ml moveList) String() string {
	return fmt.Sprint(ml...)
}

const podCost = 20
