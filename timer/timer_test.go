package timer

import (
	"testing"
	"time"
)

// var tests = []struct{
// 	description string
// 	time time.Duration
// 	want string
// }{
// 	{
// 		description: "Format seconds",
// 		time: 4,
// 		want: "4.0",
// 	},
// }

// func TestFormatSeconds(t *testing.T){
// 	for _,tt:=range tests{
// 		t.Run(tt.description, func(t *testing.T){
// 			if got:=FormatSeconds(tt.time); got!=tt.want{
// 				t.Errorf("formatSeconds(%v) = %v; want %v", tt.time, got, tt.want)
// 			}
// 		})
// 	}
// }

var testCases = []struct {
	description string
	start       string
	duration    time.Duration
	want        string
}{
	{
		description: "20 seconds has passed",
		start:       "Thu Aug 27 03:00:00 PST 2020",
		duration:    20 * time.Second,
		want:        "Thu Aug 27 03:00:20 PST 2020",
	},
	{
		description: "20 minutes has passed",
		start:       "Thu Aug 27 03:00:00 PST 2020",
		duration:    20 * time.Minute,
		want:        "Thu Aug 27 03:20:00 PST 2020",
	},
}

var startT, endT time.Time

func TestElapsedDuration(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {

			startT, _ = time.Parse(time.UnixDate, tt.start)
			endT, _ = time.Parse(time.UnixDate, tt.want)
			got := ElapsedDuration(startT, tt.duration)

			if got.Sub(startT) != tt.duration {
				t.Errorf("error calculating elapsed time using: ElapsedDuration(%v, %v) = %v; want %v", startT, tt.duration, got, endT)
			}
		})
	}
}
