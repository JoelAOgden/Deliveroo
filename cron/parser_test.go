package cron

import (
	"github.com/JoelAOgden/Deliveroo/cron/field"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseExpression_Number(t *testing.T) {

	given := "2"

	want := []int{2}

	got, err := parseExpression(given, field.Minute)
	assert.NoError(t, err)

	assert.Equal(t, want, got)

}

func TestParseExpression_Dash(t *testing.T) {

	given := "2-5"

	want := []int{2, 3, 4, 5}

	got, err := parseExpression(given, field.Minute)
	assert.NoError(t, err)

	assert.Equal(t, want, got)

}

func TestParseExpression_Asterisk(t *testing.T) {

	given := "*"

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	got, err := parseExpression(given, field.Month)
	assert.NoError(t, err)

	assert.Equal(t, want, got)

}

func TestParseExpression_Slash(t *testing.T) {

	given := "*/2"

	want := []int{2, 4, 6, 8, 10, 12}

	got, err := parseExpression(given, field.Month)
	assert.NoError(t, err)

	assert.Equal(t, want, got)

}

func TestParseExpression_Comma(t *testing.T) {

	given := "7,1,2"

	want := []int{7, 1, 2}

	got, err := parseExpression(given, field.Month)
	assert.NoError(t, err)

	assert.Equal(t, want, got)

}

func TestParseExpression_Complex(t *testing.T) {

	given := "17,*/15,3-5"

	want := []int{17, 0, 15, 30, 45, 3, 4, 5}

	got, err := parseExpression(given, field.Minute)
	assert.NoError(t, err)

	assert.Equal(t, want, got)

}
