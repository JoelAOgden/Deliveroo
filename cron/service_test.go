package cron

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Parse(t *testing.T) {

	given := []string{"*/15", "0", "1,15", "*", "1-5"}
	want := Info{
		Minute:        []int{0, 15, 30, 45},
		Hour:          []int{0},
		DayOfTheMonth: []int{1, 15},
		Month:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		DayOfTheWeek:  []int{1, 2, 3, 4, 5},
	}

	s := Service{}

	got, err := s.Parse(given)

	assert.NoError(t, err)
	assert.Equal(t, want, *got)

}
