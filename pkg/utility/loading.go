package utility

import (
	"fmt"
	"time"
)

type LoadingVisualization struct {
	Active          bool
	AnimationFrames []string
	finished        chan interface{}
}

// create and start a new LoadingVisualziation instance
func StartLoading() *LoadingVisualization {
	l := LoadingVisualization{}
	l.Start()
	return &l
}

// prepare and run the loading animation in a new goroutine
func (l *LoadingVisualization) Start() {
	if l.AnimationFrames == nil {
		l.AnimationFrames = []string{"|", "/", "-", "\\", "|"}
	}
	l.finished = make(chan interface{})

	go l.Run(10)
}

// run the loading animation
func (l *LoadingVisualization) Run(speed time.Duration) {
	index := 0
	for {
		select {
		case <-l.finished:
			return
		default:
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
