package game

import (
	"cmp"
	"slices"
)


type GameBuilder struct{
	segments []float32
}

func (builder *GameBuilder)Add(segment float32){
	builder.segments = append(builder.segments, segment)
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
	}
}