package model

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	inDateFormat  = "20060102"
	outDateFormat = "02/01/2006"
	inTimeFormat  = "15:04"
	outTimeFormat = "3:04PM"
)

// Guest encapsulates information about guests
type Guest struct {
	Name string `json:"name"`
	Info string `json:"info"`
}

// Event encapsulates information about an event
type Event struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	DateTime string  `json:"dateTime"`
	Location string  `json:"location"`
	Guests   []Guest `json:"guests"`
	Info     string  `json:"info"`
}

// DateTimeValid indicates whether the receiver's DateTime property is in a valid format
func (e *Event) DateTimeValid() (bool, error) {
	return regexp.MatchString("^20[12]\\d([01]\\d){2}T([01]\\d|2[0-3]):[0-5]\\d$", e.DateTime)
}

func (e *Event) getDateTimePart(p int, intFmt string, outFmt string) (string, error) {
	v, err := e.DateTimeValid()
	if err != nil {
		return "", err
	}

	if !v {
		return "", fmt.Errorf("invalid DateTime string %s", e.DateTime)
	}

	s := strings.Split(e.DateTime, "T")

	if len(s) != 2 {
		return "", fmt.Errorf("splitting produced %d parts - expected 2", len(s))
	}

	parsed, err := time.Parse(intFmt, s[p])

	if err != nil {
		return "", fmt.Errorf("parsing %s with %s: %s", s[p], intFmt, err.Error())
	}

	return parsed.Format(outFmt), nil
}

// DateString gets the date part from the reciever's DateTime property
func (e *Event) DateString() (string, error) {
	return e.getDateTimePart(0, inDateFormat, outDateFormat)
}

// TimeString gets the time part from the receiver's DateTime property
func (e *Event) TimeString() (string, error) {
	return e.getDateTimePart(1, inTimeFormat, outTimeFormat)
}
