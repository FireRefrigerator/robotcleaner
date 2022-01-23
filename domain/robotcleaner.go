package domain

type RobotCleaner struct {
	Location
}

type Location struct {
	X         int
	Y         int
	Direction int
}

func NewCleaner(x, y int) RobotCleaner {
	cleaner := RobotCleaner{
		Location: Location{
			X:         x,
			Y:         y,
			Direction: 0,
		},
	}
	return cleaner
}
