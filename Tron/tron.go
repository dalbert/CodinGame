package main

import "fmt"
import "os"

// WIDTH of the arena
const WIDTH int = 30

// HEIGHT of the arena
const HEIGHT int = 20

// arbitrary integer representations of the 4 available directions
const LEFT, RIGHT, UP, DOWN int = 0, 1, 2, 3

// Player the light cycle
type Player struct {
	tailX, tailY, headX, headY int
	direction                  int
}

func main() {
	var me Player
	for {
		// N: total number of players (2 to 4).
		// P: your player number (0 to 3).
		var N, P int
		fmt.Scan(&N, &P)

		players := [4]Player{}
		for i := 0; i < N; i++ {
			// X0: starting X coordinate of lightcycle (or -1)
			// Y0: starting Y coordinate of lightcycle (or -1)
			// X1: starting X coordinate of lightcycle (can be the same as X0 if you play before this player)
			// Y1: starting Y coordinate of lightcycle (can be the same as Y0 if you play before this player)
			var X0, Y0, X1, Y1 int
			fmt.Scan(&X0, &Y0, &X1, &Y1)
			players[i].tailX = X0
			players[i].tailY = Y0
			players[i].headX = X1
			players[i].headY = Y1
		}
		me = players[P]
		me.direction = nearestEdge(me)
		fmt.Fprintln(os.Stderr, me)
		if me.headX > 0 && me.direction != RIGHT {
			me.direction = LEFT
			fmt.Println("LEFT")
		} else if me.headX < 29 && me.direction != LEFT {
			me.direction = RIGHT
			fmt.Println("RIGHT")
		} else if me.headY < 19 && me.direction != UP {
			me.direction = DOWN
			fmt.Println("DOWN")
		} else if me.direction != DOWN {
			me.direction = UP
			fmt.Println("UP")
		} else {
			me.direction = LEFT
			fmt.Println("LEFT")
		}
	}
}

func nearestEdge(player Player) int {
	distances := [4]int{LEFT: player.headX, RIGHT: WIDTH - player.headX, UP: player.headY, DOWN: HEIGHT - player.headY}
	direction := LEFT
	if distances[RIGHT] < distances[direction] {
		direction = RIGHT
	}
	if distances[UP] < distances[direction] {
		direction = UP
	}
	if distances[DOWN] < distances[direction] {
		direction = DOWN
	}
	fmt.Fprintln(os.Stderr, player.headX, player.headY, direction)
	return direction
}
