package domain

type RobotCleaner struct {
	Location
}

const (
	Direction_EAST = iota
	Direction_SOUTH
	Direction_WEST
	Direction_NORTH
)

type Location struct {
	X         int
	Y         int
	Direction int
}

func NewCleaner(x, y, d int) RobotCleaner {
	cleaner := RobotCleaner{
		Location: Location{
			X:         x,
			Y:         y,
			Direction: d,
		},
	}
	return cleaner
}
