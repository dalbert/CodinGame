package main

import "fmt"
import "math"
import "os"

// Enemy butts
type Enemy struct {
	id, x, y, life int
	distance       float64 // distance from player
}

// X butts
func (enemy Enemy) X() int { return enemy.x }

// Y butts
func (enemy Enemy) Y() int { return enemy.y }

// Data butts
type Data struct {
	id, x, y int
}

// X butts
func (data Data) X() int { return data.x }

// Y butts
func (data Data) Y() int { return data.y }

//Player butts
type Player struct {
	x, y int
}

// X butts
func (player Player) X() int { return player.x }

// Y butts
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
		var newEnemy Enemy
		for i := 0; i < enemyCount; i++ {
			var enemyID, enemyX, enemyY, enemyLife int
			fmt.Scan(&enemyID, &enemyX, &enemyY, &enemyLife)
			newEnemy = Enemy{id: enemyID, x: enemyX, y: enemyY, life: enemyLife}
			newEnemy.distance = calcDistance(player, newEnemy)
			enemies = append(enemies, newEnemy)
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
		datumID = findNearest(enemy, dataCoords)
		enemy.x, enemy.y = calcMovement(enemy, data[datumID], 500)
		enemies[i] = enemy
	}
}

func findNearest(subject Coord, targets []Coord) int {
	var minID, minDistance, distance int = 0, 20000, 20000
	for id, target := range targets {
		distance = int(calcDistance(subject, target))
		if distance < minDistance {
			minDistance = distance
			minID = id
		}
	}
	return minID
}

func calcDistance(a Coord, b Coord) float64 {
	horizontal := math.Pow(float64(a.X()-b.X()), 2)
	vertical := math.Pow(float64(a.Y()-b.Y()), 2)
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
