package domain

type InstractionInterface interface {
	Do(cleaner *RobotCleaner)
}

type Forward1Step struct {
}

func (ins *Forward1Step) Do(cleaner *RobotCleaner) {
	if cleaner.Direction == 0 {
		cleaner.Location.Y++
	}
	if cleaner.Direction == 1 {
		cleaner.Location.X++
	}
	if cleaner.Direction == 2 {
		cleaner.Location.Y--
	}
	if cleaner.Direction == 3 {
		cleaner.Location.X--
	}
}

type ForwardNStep struct {
	N int
}

func (ins ForwardNStep) Do(cleaner *RobotCleaner) {
	forward1 := Forward1Step{}
	for i := 0; i < ins.N; i++ {
		forward1.Do(cleaner)
	}
}

type Backward1Step struct {
}

func (ins *Backward1Step) Do(cleaner *RobotCleaner) {
	if cleaner.Direction == 0 {
		cleaner.Location.Y--
	}
	if cleaner.Direction == 1 {
		cleaner.Location.X--
	}
	if cleaner.Direction == 2 {
		cleaner.Location.Y++
	}
	if cleaner.Direction == 3 {
		cleaner.Location.X++
	}
}

type BackwardNStep struct {
	N int
}

func (ins BackwardNStep) Do(cleaner *RobotCleaner) {
	backward1 := Backward1Step{}
	for i := 0; i < ins.N; i++ {
		backward1.Do(cleaner)
	}
}

type TurnRight struct {
}

func (ins *TurnRight) Do(cleaner *RobotCleaner) {
	if cleaner.Direction == Direction_NORTH {
		cleaner.Direction = Direction_EAST
		return
	}
	if cleaner.Direction == Direction_EAST {
		cleaner.Direction = Direction_SOUTH
		return
	}
	if cleaner.Direction == Direction_SOUTH {
		cleaner.Direction = Direction_WEST
		return
	}
	if cleaner.Direction == Direction_WEST {
		cleaner.Direction = Direction_NORTH
		return
	}
}

type TurnLeft struct {
}

func (ins *TurnLeft) Do(cleaner *RobotCleaner) {
	if cleaner.Direction == Direction_NORTH {
		cleaner.Direction = Direction_WEST
		return
	}
	if cleaner.Direction == Direction_EAST {
		cleaner.Direction = Direction_NORTH
		return
	}
	if cleaner.Direction == Direction_SOUTH {
		cleaner.Direction = Direction_EAST
		return
	}
	if cleaner.Direction == Direction_WEST {
		cleaner.Direction = Direction_SOUTH
		return
	}
}

type Repeat struct {
	Ins InstractionInterface
	N   int
}

func (ins *Repeat) Do(cleaner *RobotCleaner) {
	for i := 0; i < ins.N; i++ {
		ins.Ins.Do(cleaner)
	}
}

type Sequence struct {
	Inss []InstractionInterface
}

func (ins *Sequence) Do(cleaner *RobotCleaner) {
	for i := 0; i < len(ins.Inss); i++ {
		ins.Inss[i].Do(cleaner)
	}
}
