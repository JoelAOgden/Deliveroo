package cron

import (
	"fmt"
	"github.com/JoelAOgden/Deliveroo/cron/field"
	"strconv"
	"strings"
)

func parseExpression(in string, f field.Field) ([]int, error) {

	out := make([]int, 0)

	// todo: neaten this up, it's a bit gross
	if strings.Contains(in, ",") {
		values, err := parseCommas(in, f)
		if err != nil {
			return nil, err
		}
		out = append(out, values...)
	} else if strings.Contains(in, "/") {
		values, err := parseSlash(in, f)
		if err != nil {
			return nil, err
		}
		out = append(out, values...)
	} else if strings.Contains(in, "-") {
		values, err := parseDash(in, f)
		if err != nil {
			return nil, err
		}
		out = append(out, values...)
	} else if isAsterisk(in) {
		values, err := parseAsterisk(f)
		if err != nil {
			return nil, err
		}
		out = append(out, values...)
	} else if isNumber(in) {
		value, err := parseNumber(in)
		if err != nil {
			return nil, err
		}
		out = append(out, value)
	}

	return out, nil
}

func parseCommas(in string, f field.Field) ([]int, error) {

	out := make([]int, 0)

	vs := strings.Split(in, ",")

	for _, v := range vs {
		newValues, err := parseExpression(v, f)
		if err != nil {
			return nil, err
		}

		out, err = CommaOperation(out, newValues, f)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func parseSlash(in string, f field.Field) ([]int, error) {

	vs := strings.Split(in, "/")

	var divisor *int

	// sanity checks
	if len(vs) > 2 {
		return nil, fmt.Errorf("'%s' is an uncertain expression", in)
	}

	if len(vs) == 2 {
		d, err := parseNumber(vs[1])
		if err != nil {
			return nil, err
		}
		divisor = &d
	}

	out := make([]int, 0)

	newValues, err := parseExpression(vs[0], f)
	if err != nil {
		return nil, err
	}

	out = append(out, newValues...)

	if divisor == nil {
		return out, nil
	}

	return SlashOperation(out, *divisor, f)
}

func parseDash(in string, f field.Field) ([]int, error) {

	vs := strings.Split(in, "-")

	if len(vs) != 2 {
		return nil, fmt.Errorf("invalid dash expression: '%s'", vs)
	}

	min, err := parseNumber(vs[0])
	if err != nil {
		return nil, fmt.Errorf("invalid dash expression: '%s' error: %s", vs, err)
	}

	max, err := parseNumber(vs[1])
	if err != nil {
		return nil, fmt.Errorf("invalid dash expression: '%s' error: %s", vs, err)
	}

	return DashOperation(min, max, f)
}

func parseAsterisk(f field.Field) ([]int, error) {
	return AsteriskOperation(f)
}

func parseNumber(in string) (int, error) {
	out, err := strconv.Atoi(in)

	if err != nil {
		return -1, fmt.Errorf("'%s' is an invalid numerical value", in)
	}

	return out, err
}
