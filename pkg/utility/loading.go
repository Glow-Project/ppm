package utility

import (
	"fmt"
	"time"
)

type LoadingVisualization struct {
	Active          bool
	AnimationFrames []string
	finished        chan any
}

// create and start a new LoadingVisualziation instance
func StartLoading() *LoadingVisualization {
	l := LoadingVisualization{}
	l.Start()
	return &l
}

// prepare and run the loading animation in a new goroutine
func (l *LoadingVisualization) Start() {
	// check for custom animation frames
	if l.AnimationFrames == nil {
		l.AnimationFrames = []string{"|", "/", "-", "\\", "|"}
	}
	l.finished = make(chan any)

	// run the loading animation in a seperate goroutine
	go l.Run(10)
}

// run the loading animation
func (l *LoadingVisualization) Run(speed time.Duration) {
	// index of the current loading symbol
	index := 0
	for {
		select {
		// check if something has been sent over the `finished` channel, and exit if it has
		case <-l.finished:
			return
		default:
			// overwrite the current loading symbol with the new one
			fmt.Printf("\r%s", l.AnimationFrames[index])

			if index == len(l.AnimationFrames)-1 {
				index = 0
			} else {
				index++
			}
			time.Sleep(time.Second / speed)
		}
	}
}

// stop the loading animation
func (l *LoadingVisualization) Stop() {
	l.finished <- nil
}
