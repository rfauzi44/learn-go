// Given a time in -hour AM/PM format, convert it to military (24-hour) time.

// Note: - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
// - 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.

package main

import (
	"strconv"
)

func timeConversion(s string) string {

	amorpm := s[len(s)-2 : len(s)]

	var frontInt int
	frontInt, _ = strconv.Atoi(s[:2])
	var frontResult string

	if amorpm == "PM" {
		if frontInt != 12 {
			frontResult = strconv.Itoa(frontInt + 12)
		} else {
			frontResult = s[:2]

		}
	} else {
		if frontInt == 12 {
			frontResult = "00"
		} else {
			frontResult = s[:2]

		}

	}

	return frontResult + s[2:(len(s)-2)]
}
