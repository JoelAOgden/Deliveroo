package cron

import (
	"fmt"
	"github.com/JoelAOgden/Deliveroo/cron/field"
)

type Info struct {
	Minute        []int
	Hour          []int
	DayOfTheMonth []int
	Month         []int
	DayOfTheWeek  []int
}

type Service struct {
}

func (s Service) Parse(in []string) (*Info, error) {

	// sanity check
	if len(in) != 5 {
		return nil, fmt.Errorf("not a valid expression '%s'", in)
	}

	minute, err := parseExpression(in[0], field.Minute)
	if err != nil {
		return nil, err
	}
	hour, err := parseExpression(in[1], field.Hour)
	if err != nil {
		return nil, err
	}
	dayOfTheMonth, err := parseExpression(in[2], field.DayOfTheMonth)
	if err != nil {
		return nil, err
	}
	month, err := parseExpression(in[3], field.Month)
	if err != nil {
		return nil, err
	}
	dayOfTheWeek, err := parseExpression(in[4], field.DayOfTheWeek)
	if err != nil {
		return nil, err
	}

	out := Info{
		Minute:        minute,
		Hour:          hour,
		DayOfTheMonth: dayOfTheMonth,
		Month:         month,
		DayOfTheWeek:  dayOfTheWeek,
	}

	return &out, nil

}
