package timer

import (
	"fmt"
	"math"
	"os"
	"time"
)

// Pomodoro defines the structure of a pomodoro timer
type Pomodoro struct {
	Start         time.Time
	FocusDuration time.Duration
	MaxIntervals  int
	ShortBreak    time.Duration
	LongBreak     time.Duration
}

// SetPomodoro initializes a pomodoro timer
func SetPomodoro(timer Pomodoro) {

	// open a buffer for the sound that marks the end of 1 pomodoro timer
	buffer := StartPlayer("dog2.mp3")

	for i := 1; i < timer.MaxIntervals; i++ {
		// initialize the starting time
		start := time.Now()

		// calculate the finish time
		finish := ElapsedDuration(start, timer.FocusDuration)

		fmt.Printf("--Interval %v focus time %v. \n", i, start)

		// count until the target focus time is reached
		Countdown(finish)

		// trigger an ending signal
		PlaySound(*buffer)

		fmt.Printf("Ending interval %v focus time %v. \n", i, finish)
		fmt.Println("-Time for a 5 min break.")

		// count until the target break time is reached
		Countdown(ElapsedDuration(finish, timer.ShortBreak))

	}

}

// FormatSeconds formats the time durations in seconds
func FormatSeconds(t time.Duration) time.Duration {
	//return time.Duration(t.Seconds()) * 1000
	return time.Duration(int64(time.Millisecond) * int64(time.Duration(t.Seconds())))
}

// formatMinutes formats the time durations in minutes
func formatMinutes(t time.Duration) string {
	return fmt.Sprintf("%02.1f", math.Abs(t.Minutes()))
}

// ElapsedDuration calculates the finish time
func ElapsedDuration(t time.Time, d time.Duration) time.Time {
	return t.Add(d)
}

// Countdown indicates timeleft since the timer started
func Countdown(target time.Time) {
	fmt.Println("Starting counting ...")
	for range time.Tick(1000 * time.Millisecond) {
		timeLeft := -time.Since(target)
		if timeLeft < 0 {
			fmt.Printf("Finished counting %v\n", 0)
			return
		}
		timeLeft = time.Duration(timeLeft)
		fmt.Fprint(os.Stdout, "#: ", FormatSeconds(timeLeft), "\n")
	}
	return
}
