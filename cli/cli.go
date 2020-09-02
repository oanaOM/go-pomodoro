package cli

import (
	"flag"
	"time"

	"github.com/oanaOM/go-pomodoro/timer"
)

// NewPomodoro ....
var NewPomodoro = timer.Pomodoro{
	Start:         time.Now(),
	FocusDuration: 5 * time.Second,
	MaxIntervals:  4,
	ShortBreak:    5 * time.Second,
}

func init() {
	const (
		defaultTimer = 25
		usa
	)

	// subcommands
	duration := flag.Int("duration", 25, "Change the pomodoro time. Default is 25 minutes")
	shortbreak := flag.Int("shortbreak", 5, "Change the shortbreak time. Default is 5 minutes")
	intervals := flag.Int("intervals", 5, "Change the frequesce of the pomodoro timers. Default is 5")

	flag.Parse()

	if *duration != 0 {
		NewPomodoro.FocusDuration = time.Duration(*duration) * time.Second

	}
	if *shortbreak != 0 {
		NewPomodoro.ShortBreak = time.Duration(*shortbreak) * time.Second
	}
	if *intervals != 0 {
		NewPomodoro.MaxIntervals = *intervals
	}

}
