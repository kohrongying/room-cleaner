package robot

import (
	"testing"
)
func TestRobotGetPosition(t *testing.T) {
	got := Robot{0, 0, N}.GetPosition()
	want := "0, 0, N"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestRobotTurn(t *testing.T) {
	t.Run("Turn Right", func(t *testing.T) {
		t.Run("when facing north", func(t *testing.T) {
			r := Robot{0, 0, N}
			r.Turn(1)
			got := r.GetPosition()
			want := "0, 0, E"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	
		t.Run("when facing west", func(t *testing.T) {
			r := Robot{0, 0, W}
			r.Turn(1)
			got := r.GetPosition()
			want := "0, 0, N"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	})

	t.Run("Turn Left", func(t *testing.T) {
		t.Run("when facing west", func(t *testing.T) {
			r := Robot{0, 0, W}
			r.Turn(-1)
			got := r.GetPosition()
			want := "0, 0, S"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})

		t.Run("when facing north", func(t *testing.T) {
			r := Robot{0, 0, N}
			r.Turn(-1)
			got := r.GetPosition()
			want := "0, 0, W"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	})
	
}

func TestRobotMove(t *testing.T) {
	t.Run("Forward", func(t *testing.T) {
		t.Run("Facing North", func(t *testing.T) {
			r := Robot{0, 0, N}
			r.Move(1)
			got := r.GetPosition()
			want := "0, 1, N"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})

		t.Run("Facing East", func(t *testing.T) {
			r := Robot{0, 0, E}
			r.Move(1)
			got := r.GetPosition()
			want := "1, 0, E"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})

		t.Run("Facing South", func(t *testing.T) {
			r := Robot{0, 0, S}
			r.Move(1)
			got := r.GetPosition()
			want := "0, -1, S"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})

		t.Run("Facing West", func(t *testing.T) {
			r := Robot{0, 0, W}
			r.Move(1)
			got := r.GetPosition()
			want := "-1, 0, W"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	})

	t.Run("Backward", func(t *testing.T) {
		t.Run("Facing North", func(t *testing.T) {
			r := Robot{0, 0, N}
			r.Move(-1)
			got := r.GetPosition()
			want := "0, -1, N"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})

		t.Run("Facing East", func(t *testing.T) {
			r := Robot{0, 0, E}
			r.Move(-1)
			got := r.GetPosition()
			want := "-1, 0, E"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})

		t.Run("Facing South", func(t *testing.T) {
			r := Robot{0, 0, S}
			r.Move(-1)
			got := r.GetPosition()
			want := "0, 1, S"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})

		t.Run("Facing West", func(t *testing.T) {
			r := Robot{0, 0, W}
			r.Move(-1)
			got := r.GetPosition()
			want := "1, 0, W"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	})
}

func TestRobot(t *testing.T) {
	got := Robot{0, 0, N}
	want := got
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
