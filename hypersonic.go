package main

import "fmt"
//import "os"
import "bytes"
import "strconv"

// These indicate the contents of each cell of the floor in the input
const WALL   = -2
const CELL   = -1
const BOX    = 0
const ITEM_RANGE  = 1
const ITEM_BOMB   = 2
const EXPLOSION = 32

// These are the values given to each type of cell during scoring of the floor
const WALL_SCORE    = -20
const DANGER_SCORE  = -30
const CELL_SCORE    = -1
const TOO_FAR       = 100

// These are entity types
const PLAYER = 0
const BOMB   = 1
const ITEM   = 2

// floor constraints
const WIDTH  = 13
const HEIGHT = 11

// time it takes for a bomb to explode, in turns (useless?)
const MAX_BOMB_TIME = 8
// turns I'm willing to spend searching for a path to a particular cell
// I had timeouts when this was > 8, which of course would also depend upon the ineffficiency of my code at the time
const SEARCH_DEPTH_LIMIT = 8

type Bomb struct {
    x, y int    // coordinates of the bomb
    time int    // number of turns until explosion
    reach int   // number of cells it reaches in each direction
}
type Cell struct {
    score int
    distance int
}

func main() {
    var width, height, myId int
    fmt.Scan(&width, &height, &myId)
    
    turnCounter := 0
    myX, myY := 2, 2
    var myReach int       // explosion range of my bombs
    var myBombCount int
    var floor [WIDTH][HEIGHT]int
    for {
        bombsOnTheFloor := []Bomb{}
        floor = [WIDTH][HEIGHT]int{}
        for i := 0; i < height; i++ {   // the grid of boxes and empty cells
            var row string
            fmt.Scan(&row)
//            fmt.Fprintln(os.Stderr, row)
            floor = buildTheFloor(row, i, floor)
        }
//        fmt.Fprintln(os.Stderr, floorToString(floor))
        
        var entities int
        fmt.Scan(&entities)             // how many players & bombs are on the grid
        for i := 0; i < entities; i++ { // info about all the players & bombs on the grid
            var entityType, owner, x, y, param1, param2 int
            fmt.Scan(&entityType, &owner, &x, &y, &param1, &param2)
            if entityType == PLAYER && owner == myId {
                myBombCount = param1
                myReach = param2 // may have changed due to power ups
                myX = x
                myY = y
            }
            if entityType == BOMB { // don't bother going here (get x,y and affect their score somehow)
                bombsOnTheFloor = append(bombsOnTheFloor, 
                                  Bomb{x: x, y: y, time: param1, reach: param2})
            }
        }
//        fmt.Fprintln(os.Stderr, bombsOnTheFloor)
        
        xT, yT := myX, myY
        maxScore := Cell{score: WALL_SCORE, distance: TOO_FAR}
        scoreFloor := scoreTheFloor(myX, myY, bombsOnTheFloor, myReach, floor) 
        scoreFloor = markBombs(bombsOnTheFloor, scoreFloor)
        floor = transferExplosions(scoreFloor, floor)
        
        for i := 0; i < WIDTH; i++ {
            for j := 0; j < HEIGHT; j++ {
                if scoreFloor[i][j].score > maxScore.score || (scoreFloor[i][j].score == maxScore.score && scoreFloor[i][j].distance < maxScore.distance) {
                    maxScore = scoreFloor[i][j]
                    xT = i
                    yT = j
                }
            }
        }
//        fmt.Fprintln(os.Stderr, fmt.Sprintf("MAX: %d, TARGET: %d, x: %d, y: %d", maxScore, scoreFloor[xT][yT], xT, yT))
//        fmt.Fprintln(os.Stderr, scoreFloorToString(scoreFloor))
//        fmt.Fprintln(os.Stderr, floorToString(floor))
//        if canIBeHere(myX, myY, 0, bombsOnTheFloor, floor) && !canIBeHere(x, y, 1, bombsOnTheFloor, floor)
        if  myBombCount > 0 && myX == xT && myY == yT && canIEscapeThisBomb(myX, myY, Bomb{x: myX, y: myY, reach: myReach, time: MAX_BOMB_TIME}, MAX_BOMB_TIME, 0, myReach, bombsOnTheFloor, floor) { // drop bomb on current cell while moving toward target cell (could be equivalent)
            fmt.Println(fmt.Sprintf("BOMB %d %d BUTT SOUP", xT, yT))
        } else {
            fmt.Println(fmt.Sprintf("MOVE %d %d (%d, %d) = %d", xT, yT, xT, yT, scoreFloor[xT][yT]))
        }
        
        turnCounter++
    }
}

/**
 * 
 **/
func canIEscapeThisBomb(myX int, myY int, bomb Bomb, turnLimit int, numTurns int, reach int, bombs []Bomb, floor [WIDTH][HEIGHT]int) bool {
//    fmt.Fprintln(os.Stderr, fmt.Sprintf("MX: %d, MY: %d, BX: %d, BY: %d, t: %d, r: %d", myX, myY, bombX, bombY, numTurns, reach))
    // Already safe on a diagonal from the bomb's cell, don't need to move
    if myX != bomb.x && myY != bomb.y {return true}
    // I'm lined up with the bomb, but out of its reach
//    fmt.Fprintln(os.Stderr, fmt.Sprintf("myX: %d, myY: %d, BX: %d, BY: %d, turns: %d", myX, myY, bombX, bombY, numTurns))
    if myX > bomb.x + reach || myX < bomb.x - reach || myY > bomb.y + reach || myY < bomb.y - reach {return true}
//    fmt.Fprintln(os.Stderr, "MARKER 1")
    // In danger, need to move, but there is no time left
    if turnLimit - numTurns < 1 {return false}
    // In danger, need to move, have some time left
    if canIBeHere(myX+1, myY, numTurns+1, bombs, floor) && canIEscapeThisBomb(myX+1, myY, bomb, turnLimit, numTurns+1, reach, bombs, floor) {return true}
    if canIBeHere(myX-1, myY, numTurns+1, bombs, floor) && canIEscapeThisBomb(myX-1, myY, bomb, turnLimit, numTurns+1, reach, bombs, floor) {return true}
    if canIBeHere(myX, myY+1, numTurns+1, bombs, floor) && canIEscapeThisBomb(myX, myY+1, bomb, turnLimit, numTurns+1, reach, bombs, floor) {return true}
    if canIBeHere(myX, myY-1, numTurns+1, bombs, floor) && canIEscapeThisBomb(myX, myY-1, bomb, turnLimit, numTurns+1, reach, bombs, floor) {return true}
    // in danger, no where to go
    return false
}

func willIDieHere(x int, y int, bombs []Bomb, floor [WIDTH][HEIGHT]int) bool {
    for _, bomb := range bombs {
        if canIEscapeThisBomb(x, y, bomb, bomb.time, 0, bomb.reach, bombs, floor) {
            return false
        }
    }
    return true
}

/**
 * How many boxes are within bombing range of the given cell, are there items in those boxes, and can I get there?
 **/
func scoreACell(x int, y int, myX int, myY int, bombsOnTheFloor []Bomb, myReach int, floor [WIDTH][HEIGHT]int) Cell {
    if (myX != x || myY != y) { // I'm not already standing here
        if !canIBeHere(x, y, 1, bombsOnTheFloor, floor) {return Cell{score: WALL_SCORE, distance: TOO_FAR}} // cannot move to here next turn
    }
    moves, maybe := canIGoToThere(myX, myY, myX, myY, x, y, SEARCH_DEPTH_LIMIT, bombsOnTheFloor, floor)
    if !maybe {return Cell{score: WALL_SCORE, distance: TOO_FAR}} // cannot get here, even after multiple turns
    if willIDieHere(x, y, bombsOnTheFloor, floor) {return Cell{score: DANGER_SCORE, distance: TOO_FAR}} // does not account for time left on the bomb, could optimize here rather than walling it off
    score := 0
    for i := 0; i < myReach; i++ {
        if x+i < WIDTH && floor[x+i][y] >= BOX {score++}
        if x-i > 0 && floor[x-i][y] >= BOX {score++}
        if y+i < HEIGHT && floor[x][y+i] >= BOX {score++}
        if y-i > 0 && floor[x][y-i] >= BOX {score++}
    }
    if floor[x][y] > BOX {score++} // there's an item in the box
    return Cell{score: score, distance: moves}
}

func scoreTheFloor(myX int, myY int, bombsOnTheFloor []Bomb, myReach int, floor [WIDTH][HEIGHT]int) [WIDTH][HEIGHT]Cell{
    scoreFloor := [WIDTH][HEIGHT]Cell{}
    for i := 0; i < WIDTH; i++ {
        for j := 0; j < HEIGHT; j++ {
            scoreFloor[i][j] = scoreACell(i, j, myX, myY, bombsOnTheFloor, myReach, floor)
        }
    }
    return scoreFloor
}

func canIGoToThere(x int, y int, myX int, myY int, xT int, yT int, moveLimit int, bombs []Bomb, floor [WIDTH][HEIGHT]int) (distance int, maybe bool) {
//    fmt.Fprintln(os.Stderr, fmt.Sprintf("GO - x: %d, y: %d, m: %d", myX, myY, moves))
    moves, minMoves := 0, TOO_FAR
    yes, isPathFound := false, false
    if moveLimit < 1 {return TOO_FAR, false}
    // if it's not the cell that I'm already standing on, then ensure that I can stand on it when I get there
    if (x != myX || y != myY) && !canIBeHere(myX, myY, 0, bombs, floor) {return TOO_FAR, false}
    if myX == xT && myY == yT {return moves, true}

    // try moving Right
    moves, yes = canIGoToThere(x, y, myX+1, myY, xT, yT, moveLimit-1, bombs, floor)
    if yes {
        moves++
        if moves < minMoves {minMoves = moves}
        isPathFound = true
    }
    // try moving Left
    moves, yes = canIGoToThere(x, y, myX-1, myY, xT, yT, moveLimit-1, bombs, floor) 
    if yes {
        moves++
        if moves < minMoves {minMoves = moves}
        isPathFound = true
    }   
    // try moving Down
    moves, yes = canIGoToThere(x, y, myX, myY+1, xT, yT, moveLimit-1, bombs, floor) 
    if yes {
        moves++
        if moves < minMoves {minMoves = moves}
        isPathFound = true
    }   
    // try moving Up
    moves, yes = canIGoToThere(x, y, myX, myY-1, xT, yT, moveLimit-1, bombs, floor) 
    if yes {
        moves++
        if moves < minMoves {minMoves = moves}
        isPathFound = true
    }
    
    // all possibilities exhausted
    return minMoves, isPathFound
}

func markBombs(bombs []Bomb, scoreFloor [WIDTH][HEIGHT]Cell) [WIDTH][HEIGHT]Cell {
    var dangerScore int
    for _, bomb := range bombs {
        if bomb.time < 2 {
            dangerScore = DANGER_SCORE
            // on the bomb and it's exploding
            scoreFloor[bomb.x][bomb.y].score = dangerScore
        } else {
            dangerScore = DANGER_SCORE / (bomb.time - 1) // treat the bomb's timer as if we'd already advanced to the next turn
        }
        // on the bomb
        if scoreFloor[bomb.x][bomb.y].score > dangerScore {scoreFloor[bomb.x][bomb.y].score = WALL_SCORE}
        // left of the bomb
        for i := bomb.x; i >= bomb.x - bomb.reach; i-- {
            if amIWithinTheBoundaries(i, 0) { // prevent array index out of bounds
                if scoreFloor[i][bomb.y].score > dangerScore { // do not overwrite a score that's already even lower
                    scoreFloor[i][bomb.y].score = dangerScore
                    if scoreFloor[i][bomb.y].score == WALL_SCORE {break} // stop propagating the explosion in this direction, there is a blocker (wall or box)
                }
            }
        }
        // right of the bomb
        for i := bomb.x; i <= bomb.x+bomb.reach; i++ {
            if amIWithinTheBoundaries(i, 0) { // prevent array index out of bounds
                if scoreFloor[i][bomb.y].score > dangerScore { // do not overwrite a score that's already even lower
                    scoreFloor[i][bomb.y].score = dangerScore
                    if scoreFloor[i][bomb.y].score == WALL_SCORE {break} // stop propagating the explosion in this direction, there is a blocker (wall or box)
                }
            }
        }
        // below the bomb
        for i := bomb.y; i >= bomb.y - bomb.reach; i-- {
            if amIWithinTheBoundaries(0, i) {
                if scoreFloor[bomb.x][i].score > dangerScore {
                    scoreFloor[bomb.x][i].score = dangerScore
                    if scoreFloor[bomb.x][i].score == WALL_SCORE {break} // stop propagating the explosion in this direction, there is a blocker (wall or box)
                }
            }
        }
        // above the bomb
        for i := bomb.y; i <= bomb.y+bomb.reach; i++ {
            if amIWithinTheBoundaries(0, i) {
                if scoreFloor[bomb.x][i].score > dangerScore {
                    scoreFloor[bomb.x][i].score = dangerScore
                    if scoreFloor[bomb.x][i].score == WALL_SCORE {break} // stop propagating the explosion in this direction, there is a blocker (wall or box)
                }
            }
        }
    }
    return scoreFloor
}

func buildTheFloor(row string, y int, floor [WIDTH][HEIGHT]int) [WIDTH][HEIGHT]int {
    width := len(row)
    for x := 0; x < width; x++ {
        if string(row[x]) == "." {
            floor[x][y] = CELL
        } else if string(row[x]) == "X" {
            floor[x][y] = WALL
        } else {
            floor[x][y] = int(row[x] - '0')
        }
    }
    return floor
}

func canIBeHere(x int, y int, timeElapsed int, bombs []Bomb, floor [WIDTH][HEIGHT]int) bool {
    if !amIWithinTheBoundaries(x, y) {return false}
    if floor[x][y] == WALL || floor[x][y] >= BOX || floor[x][y] == EXPLOSION {return false}
    for _, bomb := range bombs {
        if x == bomb.x && y == bomb.y {return false} // can't walk through bombs once they're placed        
    }
    return true
}

func amIWithinTheBoundaries(x int, y int) bool {
    if x < 0 || x >= WIDTH || y < 0 || y >= HEIGHT {return false}
    return true
}

// TODO: this does not account for walls & boxes, which block propagation of the explosion
func amIWithinTheBlastRadius(myX int, myY int, bomb Bomb) bool{
    if myX > bomb.x + bomb.reach || myX < bomb.x - bomb.reach || myY > bomb.y + bomb.reach || myY < bomb.y - bomb.reach {return true}
    return false
}

func transferExplosions(scoreFloor [WIDTH][HEIGHT]Cell, floor [WIDTH][HEIGHT]int) [WIDTH][HEIGHT]int {
    for i := 0; i < WIDTH; i++ {
        for j := 0; j < HEIGHT; j++ {
            if scoreFloor[i][j].score == DANGER_SCORE {floor[i][j] = EXPLOSION}
        }
    }
    return floor
}

func scoreFloorToString(floor [WIDTH][HEIGHT]Cell) string {
    var buffer bytes.Buffer
    var scoreStr, distanceStr string
    for i := 0; i < HEIGHT; i++ {
        for j := 0; j < WIDTH; j++ {
            scoreStr = strconv.Itoa(floor[j][i].score)
            distanceStr = strconv.Itoa(floor[j][i].distance)
            buffer.WriteString("[")
            for f := 0; f < 3 - len(scoreStr); f++ {buffer.WriteString(" ")}
            buffer.WriteString(scoreStr)
            buffer.WriteString(", ")
            for f := 0; f < 3 - len(distanceStr); f++ {buffer.WriteString(" ")}
            buffer.WriteString(distanceStr)
            buffer.WriteString("]")
        }
        buffer.WriteString("\n")
    }
    return buffer.String()
}

func floorToString(floor [WIDTH][HEIGHT]int) string {
    var buffer bytes.Buffer
    var cell int
    for i := 0; i < HEIGHT; i++ {
        for j := 0; j < WIDTH; j++ {
            cell = floor[j][i]
            buffer.WriteString(" ")
            if cell == BOX {buffer.WriteString("B")}
            if cell == WALL {buffer.WriteString("W")}
            if cell == CELL {buffer.WriteString("_")}
            if cell == EXPLOSION {buffer.WriteString("E")}
            if cell == ITEM_RANGE || cell == ITEM_BOMB {buffer.WriteString("I")}
            buffer.WriteString(" ")
        }
        buffer.WriteString("\n")
    }
    return buffer.String()
}

