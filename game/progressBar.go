package game

import (
	"sync"
)

type AtomicProgressBarSegments struct {
	AtomicProgressBar
	MultiThreshold
}

type AtomicProgressBar struct {
	percentage float32;
	mutex sync.Mutex
}

type MultiThreshold struct {
	thresholds []float32;
}

func (bar *AtomicProgressBarSegments) Segments() []float32 {
	return bar.thresholds
}

func (bar *AtomicProgressBarSegments) Segment() uint8 {
	for i := 0; i < len(bar.thresholds); i++ {
		if bar.percentage <= bar.thresholds[i] {
			return uint8(i)
		}
	}
	return uint8(bar.percentage * float32(len(bar.thresholds)))
}

func (bar *AtomicProgressBarSegments) Value() float32 {
	return bar.percentage
}

func (bar *AtomicProgressBarSegments) Reset() float32{
	
	bar.mutex.Lock()
	var last = bar.percentage
	bar.percentage = 0
	bar.mutex.Unlock()
	return last
}

func (bar *AtomicProgressBarSegments) IncrementByPercentage(percentage float32) {	
	bar.mutex.Lock()
	bar.percentage += percentage
	if bar.percentage > 1 {
		bar.percentage = 1
	}
	bar.mutex.Unlock()
}

func (bar *AtomicProgressBarSegments) SetValue(value float32){
	value = max(value, 0)
	value = min(value, 1)
	bar.mutex.Lock()
	bar.percentage = value
	bar.mutex.Unlock()
}



func CreateGame() *Game {
	var builder = GameBuilder{}
	builder.Add(0.25)
	builder.Add(0.5)
	builder.Add(0.75)
	builder.SetSpeed(-0.01)
	return builder.Build()
}