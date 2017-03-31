package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat     int
	progress int
}
type WAVPlayer struct {
	stat     int
	progress int
}

func (p *MP3Player) Play(source string) {

	fmt.Println("playing mp3 music", source)
	p.progress = 0

	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(".")
		p.progress += 10
	}
	fmt.Println("\n finsished playing", source)
}

func (p *WAVPlayer) Play(source string) {
	fmt.Println("Playing WAV music", source)

	p.progress = 0
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\n Finished playing", source)
}
