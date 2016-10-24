package main

import (
	"fmt"
	"strings"
)

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
		fmt.Scan(&myPlatinum)

		var neutralZoneCount int
		for i := 0; i < zoneCount; i++ {
			z := zones[i]
			var myPods int
			var enemyPods [3]int
			// zId: this zone's ID
			// ownerId: the player who owns this zone (-1 otherwise)
			// podsP0: player 0's PODs on this zone
			// podsP1: player 1's PODs on this zone
			// podsP2: player 2's PODs on this zone (always 0 for a two player game)
			// podsP3: player 3's PODs on this zone (always 0 for a two or three player game)
			var zID, ownerID, podsP0, podsP1, podsP2, podsP3 int
			fmt.Scan(&zID, &ownerID, &podsP0, &podsP1, &podsP2, &podsP3)
			switch myID {
			case -1:
				myPods, enemyPods = 0, [3]int{0, 0, 0}
				neutralZoneCount++
			case 0:
				myPods, enemyPods = podsP0, [3]int{podsP1, podsP2, podsP3}
			case 1:
				myPods, enemyPods = podsP1, [3]int{podsP0, podsP2, podsP3}
			case 2:
				myPods, enemyPods = podsP2, [3]int{podsP0, podsP1, podsP3}
			case 3:
				myPods, enemyPods = podsP3, [3]int{podsP0, podsP1, podsP2}
			}
			z.ownerID, z.myPods, z.enemyPods = ownerID, myPods, enemyPods
		}

		//		fmt.Fprintln(os.Stderr, zones)
		//		fmt.Fprintln(os.Stderr, myPlatinum)
		moves, buys := []string{}, []string{}
		podBudget := myPlatinum / podCost
		buys = claimNeutralZones(podBudget, zones, zonesByPlatinum, myID)
		buys = append(buys, reinforceOwnedZones(podBudget, zones, zonesByPlatinum, myID)...)

		moveCommand := strings.Join(moves, " ")
		fmt.Println(moveCommand)
		buyCommand := strings.Join(buys, " ")
		fmt.Println(buyCommand)

	}
}
func reinforceOwnedZones(podBudget int, zones map[int]*zone, zonesByPlatinum map[int][]int, myID int) (buys []string) {
	return
}
func claimNeutralZones(podBudget int, zones map[int]*zone, zonesByPlatinum map[int][]int, myID int) (buys []string) {
	for i := 6; i > -1; i-- {
		if podBudget < 1 {
			return
		}
		for _, zID := range zonesByPlatinum[i] {
			z := zones[zID]
			need := 1
			if z.ownerID == myID || z.ownerID != -1 {
				continue
			}
			if z.EnemyCount() > z.myPods { // TODO: if there are enemyPods, it's probably not neutral and you can't place pods directly on it
				need = z.EnemyCount() - z.myPods + 1
			}
			if podBudget > need { // can I afford to take this zone?
				buys = append(buys, fmt.Sprintf("%d %d", need, zID))
				podBudget = podBudget - need
			}
			if podBudget < 1 {
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

const podCost = 20
