package cron

import (
	"fmt"
	"github.com/JoelAOgden/Deliveroo/cron/field"
	"math"
)

func DashOperation(min int, max int, cronField field.Field) ([]int, error) {

	if min > max {
		return nil, fmt.Errorf("min cannot be larger than max")
	}

	fieldMin, fieldMax, err := field.ToRange(cronField)
	if err != nil {
		return nil, err
	}
	max = int(math.Min(float64(max), float64(fieldMax)))
	min = int(math.Max(float64(min), float64(fieldMin)))

	out := make([]int, 0)
	for i := min; i <= max; i++ {
		out = append(out, i)
	}

	return out, nil

}

func AsteriskOperation(cronField field.Field) ([]int, error) {

	min, max, err := field.ToRange(cronField)
	if err != nil {
		return nil, err
	}
	out := make([]int, 0)
	for i := min; i <= max; i++ {
		out = append(out, i)
	}

	return out, nil

}

func SlashOperation(a []int, divisor int, cronField field.Field) ([]int, error) {

	min, max, err := field.ToRange(cronField)
	if err != nil {
		return nil, err
	}

	out := make([]int, 0)
	for _, v := range a {
		if v < min || v > max {
			continue
		}

		if v%divisor != 0 {
			continue
		}

		out = append(out, v)

	}

	return out, nil
}

func CommaOperation(a []int, b []int, cronField field.Field) ([]int, error) {
	min, max, err := field.ToRange(cronField)
	if err != nil {
		return nil, err
	}

	a = append(a, b...)

	out := make([]int, 0)
	for _, v := range a {
		if v < min || v > max {
			continue
		}
		out = append(out, v)
	}

	return out, nil
}
