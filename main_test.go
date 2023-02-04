package main

import (
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestDateFormat(t *testing.T) {
	testtime := time.Date(2021, 8, 15, 14, 30, 7, 100, time.UTC)
	tests := []struct {
		timezone string
		format   string
		expected string
	}{
		{
			"Europe/Berlin",
			"2006-1-2 15:04:05",
			"2021-8-15 16:30:07",
		}, {
			"Europe/Berlin",
			"2.1.2006 15:04",
			"15.8.2021 16:30",
		}, {
			"Europe/Berlin",
			"2.1.2006 15:04:05",
			"15.8.2021 16:30:07",
		}, {
			"Europe/Berlin",
			"Mon 2.1.2006 15:04:05",
			"Sun 15.8.2021 16:30:07",
		}, {
			"",
			"Mon 2.1.2006 15:04:05",
			"Sun 15.8.2021 14:30:07",
		}, {
			"UTC",
			"Mon 2.1.2006 15:04:05",
			"Sun 15.8.2021 14:30:07",
		},
	}

	for _, test := range tests {
		t.Run(test.timezone+" "+test.format, func(t *testing.T) {
			tz, _ := time.LoadLocation(test.timezone)
			result := GetFormatedTimeString(testtime.In(tz), test.format)
			assert.Equal(t, test.expected, result)
		})
	}
}
