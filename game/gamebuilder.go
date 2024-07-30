package game

import (
	"cmp"
	"slices"
)


type GameBuilder struct{
	segments []float32
	speed float32 `default:"0.01"`
}

func (builder *GameBuilder)Add(segment float32){
	builder.segments = append(builder.segments, segment)
}

func (builder *GameBuilder)SetSpeed(speed float32){
	builder.speed = speed
}

func (builder *GameBuilder)Build() *Game{
	slices.SortFunc(builder.segments, func(a, b float32) int {
		return cmp.Compare(a, b)
	})
	return &Game{
		progress: AtomicProgressBarSegments{
			AtomicProgressBar: AtomicProgressBar{},
			MultiThreshold: MultiThreshold{
				thresholds: builder.segments,
			},
		},
		config: GameConfig{
			percentageBySecond: builder.speed,
		},
	}
}