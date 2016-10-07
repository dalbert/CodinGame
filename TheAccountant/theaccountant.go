package main

import "fmt"
import "math"
import "os"

//import "os"

// Enemy butts
type Enemy struct {
	id, x, y, life int
}

func (enemy Enemy) X() int { return enemy.x }
func (enemy Enemy) Y() int { return enemy.y }

// Data butts
type Data struct {
	id, x, y int
}

func (data Data) X() int { return data.x }
func (data Data) Y() int { return data.y }

//Player butts
type Player struct {
	x, y int
}

func (player Player) X() int { return player.x }
func (player Player) Y() int { return player.y }

// Coord butts
type Coord interface {
	X() int
	Y() int
}

func main() {
	for {
		var x, y int
		fmt.Scan(&x, &y)
		player := Player{x: x, y: y}

		var dataCount int
		fmt.Scan(&dataCount)

		var data []Data
		for i := 0; i < dataCount; i++ {
			var dataID, dataX, dataY int
			fmt.Scan(&dataID, &dataX, &dataY)
			data = append(data, Data{id: dataID, x: dataX, y: dataY})
		}

		var enemyCount int
		fmt.Scan(&enemyCount)
		var enemies []Enemy
		for i := 0; i < enemyCount; i++ {
			var enemyID, enemyX, enemyY, enemyLife int
			fmt.Scan(&enemyID, &enemyX, &enemyY, &enemyLife)
			enemies = append(enemies, Enemy{id: enemyID, x: enemyX, y: enemyY, life: enemyLife})
		}
		fmt.Fprintln(os.Stderr, enemies)
		advanceEnemies(enemies, data)
		fmt.Fprintln(os.Stderr, enemies)
		for _, enemy := range enemies {
			if enemy.life > 0 && canIKillHimBeforeHeGetsANode(enemy, data, player) {
				fmt.Println(fmt.Sprintf("SHOOT %d", enemy.id)) // MOVE x y or SHOOT id
				break
			}
		}
	}
}

func canIKillHimBeforeHeGetsANode(enemy Enemy, data []Data, player Player) bool {
	return true
}

func advanceEnemies(enemies []Enemy, data []Data) {
	dataCoords := make([]Coord, len(data))
	for i := range data {
		dataCoords[i] = data[i]
	}
	var datumID int
	for i, enemy := range enemies {
		datumID = findNearestDatum(enemy, dataCoords)
		enemy.x, enemy.y = calcMovement(enemy, data[datumID], 500)
		enemies[i] = enemy
	}
}

func findNearestDatum(enemy Coord, data []Coord) int {
	var minID, minDistance, distance int = 0, 20000, 20000
	for id, datum := range data {
		distance = int(calcDistance(enemy.X(), enemy.Y(), datum.X(), datum.Y()))
		if distance < minDistance {
			minDistance = distance
			minID = id
		}
	}
	return minID
}

func calcDistance(x1 int, y1 int, x2 int, y2 int) float64 {
	horizontal := math.Pow(float64(x1-x2), 2)
	vertical := math.Pow(float64(y1-y2), 2)
	return math.Sqrt(horizontal + vertical)
}

func calcMovement(enemy Enemy, datum Data, distance int) (x int, y int) {
	var angle float64
	if enemy.y == datum.y {
		angle = 0.0
	} else {
		slope := (enemy.x - datum.x) / (enemy.y - datum.y)
		angle = math.Atan(float64(slope))
	}
	dX := math.Sin(angle) * float64(distance)
	dY := math.Cos(angle) * float64(distance)
	if (enemy.x > datum.x && dX > 0) || (enemy.x < datum.x && dX < 0) {
		dX = dX * -1
	}
	if (enemy.y > datum.y && dY > 0) || (enemy.y < datum.y && dY < 0) {
		dY = dY * -1
	}
	return enemy.x + int(math.Ceil(dX)), enemy.y + int(math.Ceil(dY))
}
