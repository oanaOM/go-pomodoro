package timer

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// StartPlayer initializes a new player returning a buffer
func StartPlayer(song string) *beep.Buffer {
	f, err := os.Open(song)
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	defer streamer.Close()

	return buffer
}

// PlaySound plays the sound using an open buffer
func PlaySound(b beep.Buffer) {
	sound := b.Streamer(0, b.Len())
	speaker.Play(sound)
}
