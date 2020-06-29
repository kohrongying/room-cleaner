package main

import (
    "errors"
    "robot/robot"
)

func main() {  
    // cleaner := robot.Robot{0, 0, N}
    // room := Room{6, 6, cleaner}
}

type Room struct {
    Width int 
    Length int
    Cleaner robot.Robot
}

func (r *Room) MoveRobot(command string) error {
    switch command {
    case "M":
        r.Cleaner.Move(1)
        if !r.IsValid() {
            r.Cleaner.Move(-1)
        }
    case "L":
        r.Cleaner.Turn(-1)
    case "R":
        r.Cleaner.Turn(1)
    default:
        return errors.New("Invalid command")
    }
    return nil
}

func (r *Room) IsValid() bool {
    if r.Cleaner.X >= 0 && r.Cleaner.X < r.Width && r.Cleaner.Y >= 0 && r.Cleaner.Y < r.Length {
        return true
    }
    return false
}

func (r *Room) GetRobotPosition() string {
    return r.Cleaner.GetPosition()
}

func (r *Room) RobotInstructions(commands []string) {
    for _, command := range commands {
        r.MoveRobot(command)
    }
}