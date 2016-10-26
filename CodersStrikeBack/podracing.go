package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    boostSpent := false
    for {
        // nextCheckpointX: x position of the next check point
        // nextCheckpointY: y position of the next check point
        // nextCheckpointDist: distance to the next checkpoint
        // nextCheckpointAngle: angle between your pod orientation and the direction of the next checkpoint
        var x, y, nextCheckpointX, nextCheckpointY, nextCheckpointDist, nextCheckpointAngle int
        fmt.Scan(&x, &y, &nextCheckpointX, &nextCheckpointY, &nextCheckpointDist, &nextCheckpointAngle)
        
        var opponentX, opponentY int
        fmt.Scan(&opponentX, &opponentY)
        
        
//        fmt.Fprintln(os.Stderr, nextCheckpointAngle, nextCheckpointDist)
        
        // You have to output the target position
        // followed by the power (0 <= thrust <= 100)
        // i.e.: "x y thrust"
        var thrust int
        if nextCheckpointAngle > 90 || nextCheckpointAngle < -90 {
            thrust = 0
        } else {
            thrust = 100
        }
        if !boostSpent && nextCheckpointAngle == 0 && nextCheckpointDist > 7000 {
            fmt.Printf("%d %d BOOST\n", nextCheckpointX, nextCheckpointY)
        } else {
            fmt.Printf("%d %d %d\n", nextCheckpointX, nextCheckpointY, thrust)
        }
    }
}