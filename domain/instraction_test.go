package domain

import (
	"testing"
)

func TestTurnRight(t *testing.T) {
	cleaner := NewCleaner(0, 0, 0)
	ins := TurnRight{}
	ins.Do(&cleaner)
	if cleaner.Location.Direction != 1 {
		t.Errorf("test TurnRight err, want 1, actual %d", cleaner.Location.Direction)
	}
	ins.Do(&cleaner)
	if cleaner.Location.Direction != 2 {
		t.Errorf("test TurnRight err, want 1, actual %d", cleaner.Location.Direction)
	}
	ins.Do(&cleaner)
	if cleaner.Location.Direction != 3 {
		t.Errorf("test TurnRight err, want 1, actual %d", cleaner.Location.Direction)
	}
	ins.Do(&cleaner)
	if cleaner.Location.Direction != 0 {
		t.Errorf("test TurnRight err, want 1, actual %d", cleaner.Location.Direction)
	}
}

func TestTurnLeft(t *testing.T) {
	cleaner := NewCleaner(0, 0, 0)
	ins := TurnLeft{}
	ins.Do(&cleaner)
	if cleaner.Location.Direction != 3 {
		t.Errorf("test TurnRight err, want 1, actual %d", cleaner.Location.Direction)
	}
	ins.Do(&cleaner)
	if cleaner.Location.Direction != 2 {
		t.Errorf("test TurnRight err, want 1, actual %d", cleaner.Location.Direction)
	}
	ins.Do(&cleaner)
	if cleaner.Location.Direction != 1 {
		t.Errorf("test TurnRight err, want 1, actual %d", cleaner.Location.Direction)
	}
	ins.Do(&cleaner)
	if cleaner.Location.Direction != 0 {
		t.Errorf("test TurnRight err, want 1, actual %d", cleaner.Location.Direction)
	}
}

func TestForwardN(t *testing.T) {
	cleaner := NewCleaner(0, 0, 0)
	ins := ForwardNStep{N: 10}
	ins.Do(&cleaner)
	if cleaner.X != 0 || cleaner.Y != 10 {
		t.Errorf("test ForwardNStep 10, want X: 0 Y: 10, actal X: %d Y: %d", cleaner.X, cleaner.Y)
	}
}

func TestBackwardN(t *testing.T) {
	cleaner := NewCleaner(0, 0, 0)
	ins := BackwardNStep{N: 10}
	ins.Do(&cleaner)
	if cleaner.X != 0 || cleaner.Y != -10 {
		t.Errorf("test ForwardNStep 10, want X: 0 Y: 10, actal X: %d Y: %d", cleaner.X, cleaner.Y)
	}
}

func TestRepeat(t *testing.T) {
	cleaner := NewCleaner(0, 0, 0)
	ins := Repeat{Ins: &Forward1Step{}, N: 10}
	ins.Do(&cleaner)
	if cleaner.X != 0 || cleaner.Y != 10 {
		t.Errorf("test Repeat err, want X: 0 Y: 10, actal X: %d Y: %d", cleaner.X, cleaner.Y)
	}
}

func TestSequence(t *testing.T) {
	cleaner := NewCleaner(0, 0, 0)
	inss := []InstractionInterface{
		&Backward1Step{},
		&Repeat{Ins: &Forward1Step{}, N: 10},
	}
	ins := Sequence{Inss: inss}
	ins.Do(&cleaner)
	if cleaner.X != 0 || cleaner.Y != 9 {
		t.Errorf("test Sequence err, want X: 0 Y: 9, actal X: %d Y: %d", cleaner.X, cleaner.Y)
	}
}
