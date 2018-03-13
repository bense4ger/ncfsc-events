package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	validDate   = "20180102T18:45"
	invalidDate = "not a valid date"
)

func TestDateTimeValidIsValid(t *testing.T) {
	sut := &Event{DateTime: validDate}
	result, err := sut.DateTimeValid()

	assert.True(t, result, "It should be true")
	assert.Nil(t, err)
}

func TestDateTimeValidIsInvalid(t *testing.T) {
	sut := &Event{DateTime: invalidDate}
	result, err := sut.DateTimeValid()

	assert.False(t, result, "It should be false")
	assert.Nil(t, err)
}

func TestDateStringReturnsFormattedDate(t *testing.T) {
	sut := &Event{DateTime: validDate}
	result, err := sut.DateString()

	assert.EqualValues(t, "02/01/2018", result)
	assert.Nil(t, err)
}

func TestDateStringReturnsError(t *testing.T) {
	sut := &Event{DateTime: invalidDate}
	_, err := sut.DateString()

	assert.NotNil(t, err)
}

func TestTimeStringReturnsFormattedTime(t *testing.T) {
	sut := &Event{DateTime: validDate}
	result, err := sut.TimeString()

	assert.EqualValues(t, "6:45PM", result)
	assert.Nil(t, err)
}

func TestTimeStringReturnsError(t *testing.T) {
	sut := &Event{DateTime: invalidDate}
	_, err := sut.TimeString()

	assert.NotNil(t, err)
}
