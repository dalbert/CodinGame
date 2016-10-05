package main

import "fmt"
import "math"
import "os"

//import "os"

/**
 * Shoot enemies before they collect all the incriminating data!
 * The closer you are to an enemy, the more damage you do but don't get too close or you'll get killed.
 **/
type Enemy struct {
	id, x, y, life int
}
type Data struct {
	id, x, y int
}

func main() {
	for {
		var x, y int
		fmt.Scan(&x, &y)

		var dataCount int
		fmt.Scan(&dataCount)

		var data []Data
		for i := 0; i < dataCount; i++ {
			var dataId, dataX, dataY int
			fmt.Scan(&dataId, &dataX, &dataY)
			data = append(data, Data{id: dataId, x: dataX, y: dataY})
		}

		var enemyCount int
		fmt.Scan(&enemyCount)
		var enemies []Enemy
		for i := 0; i < enemyCount; i++ {
			var enemyId, enemyX, enemyY, enemyLife int
			fmt.Scan(&enemyId, &enemyX, &enemyY, &enemyLife)
			enemies = append(enemies, Enemy{id: enemyId, x: enemyX, y: enemyY, life: enemyLife})
		}
		fmt.Fprintln(os.Stderr, enemies)
		advanceEnemies(enemies, data)
		fmt.Fprintln(os.Stderr, enemies)
		for _, enemy := range enemies {
			if enemy.life > 0 {
				fmt.Println(fmt.Sprintf("SHOOT %d", enemy.id)) // MOVE x y or SHOOT id
				break
			}
		}
	}
}

func advanceEnemies(enemies []Enemy, data []Data) {
	var datumId int
	for i, enemy := range enemies {
		datumId = findNearestDatum(enemy, data)
		enemy.x, enemy.y = calcMovement(enemy, data[datumId], 500)
		fmt.Fprintln(os.Stderr, fmt.Sprintf("eId: %d   eX: %d   eY: %d", enemy.id, enemy.x, enemy.y))
		enemies[i] = enemy // neccessary?
	}
}

func findNearestDatum(enemy Enemy, data []Data) int {
	var minId, minDistance, distance int = 0, 20000, 20000
	for id, datum := range data {
		distance = int(calcDistance(enemy.x, enemy.y, datum.x, datum.y))
		if distance < minDistance {
			minDistance = distance
			minId = id
		}
	}
	return minId
}

func calcDistance(x1 int, y1 int, x2 int, y2 int) float64 {
	horizontal := math.Pow(float64(x1-x2), 2)
	vertical := math.Pow(float64(y1-y2), 2)
	return math.Sqrt(horizontal + vertical)
}

func calcMovement(enemy Enemy, datum Data, distance int) (x int, y int) {
	slope := enemy.x - datum.x/enemy.y - datum.y
	angle := math.Atan(float64(slope))
	dX := math.Cos(angle) * float64(distance)
	dY := math.Sin(angle) * float64(distance)
	fmt.Fprintln(os.Stderr, fmt.Sprintf("m: %v angle: %v   dX: %v   dY: %v  eX: %v  dX: %v  cos: %v  sin: %v", slope, angle, dX, dY, enemy.x, datum.x, math.Cos(angle), math.Sin(angle)))
	return enemy.x - int(math.Ceil(dX)), enemy.y - int(math.Ceil(dY))
}
