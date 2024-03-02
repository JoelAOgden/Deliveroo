package field

import "fmt"

type Field = int

const (
	Minute Field = iota
	Hour
	DayOfTheMonth
	Month
	DayOfTheWeek
)

func ToRange(in Field) (min, max int, err error) {
	switch in {
	case Minute:
		return 0, 59, nil
	case Hour:
		return 0, 59, nil
	case DayOfTheMonth:
		return 1, 30, nil // todo: improve this
	case Month:
		return 1, 12, nil
	case DayOfTheWeek:
		return 1, 7, nil
	default:
		return -1, -1, fmt.Errorf("unknown field - cannot give a range")
	}
}
