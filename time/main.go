package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Convert this string to time.Time:
	timeStr := "Thu May 21 2020 16:11:30 GMT+0000 (Coordinated Universal Time)"

	// We can work with (Coordinated Universal Time), however, the +0000 after the
	// GMT seems to be throwing parser off giving an error,
	// `parsing time "Thu May 21 2020 16:11:30 GMT+0000" as "Mon Jan 2 2006 15:04:05
	// MST-0700": cannot parse "" as "-0700"` Might as well just get rid of the text
	// in parenthesis as well since I want to remove the +0000, which is redundant to
	// GMT anyway. I had to adjust the layout to not include the -0700 also.
	timeStr = strings.TrimSuffix(timeStr, "+0000 (Coordinated Universal Time)")

	layout := "Mon Jan 2 2006 15:04:05 MST"

	t, err := time.Parse(layout, timeStr)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("time: %v, type: %[1]T", t)
}

// REFERENCE: https://forum.golangbridge.org/t/convert-string-to-time-time/8024/2
// What the reference above doesn't show is how to make extra words/numbers work in parsing.
// Here's a Go Playground demonstrating how GMT+0000 isn't understood:
// https://play.golang.org/p/U27dufnJRhN
// Also, the second case did not work the same in my GO env as Playground.
