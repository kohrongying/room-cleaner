package main

import (
	"testing"
	"robot/robot"
	"errors"
)

func TestRoom(t *testing.T) {
	t.Run("Get Robot Position", func(t *testing.T) {
		cleaner := robot.Robot{X: 0, Y: 0, Direction: robot.N}
		room := Room{6, 6, cleaner}
		got := room.GetRobotPosition()
		want := "0, 0, N"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("TestCase 1", func(t *testing.T) {
		cleaner := robot.Robot{X: 1, Y: 2, Direction: robot.N}
		room := Room{6, 6, cleaner}
	
		instructions := []string{"L", "M", "L", "M", "L", "M", "L", "M", "M"}
		room.RobotInstructions(instructions)
	
		got := room.GetRobotPosition()
		want := "1, 3, N"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("TestCase 2", func(t *testing.T) {
		cleaner := robot.Robot{X: 3, Y: 5, Direction: robot.N}
		room := Room{6, 6, cleaner}
		
		instructions := []string{"M", "L", "M"}
		room.RobotInstructions(instructions)

		got := room.GetRobotPosition()
		want := "2, 5, W"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestRoomError(t *testing.T) {
	cleaner := robot.Robot{X: 3, Y: 5, Direction: robot.N}
	room := Room{6, 6, cleaner}
	
	err := room.MoveRobot("A")
	if err == nil {
		t.Errorf("Expected error but received no error")
	}
	var ErrInvalidCommand = errors.New("Invalid command")
	if errors.Is(err, ErrInvalidCommand) {
		t.Errorf("Unexpected error %v, but received %v", "Invalid command", err)
	}
}