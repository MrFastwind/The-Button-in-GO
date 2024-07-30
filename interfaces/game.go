package interfaces

type IStartable interface{
	Start() error
	HasStarted() bool
	IsRunning() bool
}

type IGame[T any] interface{
	IStartable
	GetGameData() *T
}

type IGameBuilder[F Float, I Index] interface{
	Add(segment F)
	SetSpeed(speed F)
	Build() *IGame[ISegmentedProgressBar[F,I]]
}



type ISegmentedProgressBar[T Float, I Index] interface {
	IProgressBar[T]
	Segments() []T
	Segment() I
}

type IProgressBar[T Float] interface {
	Value() T
	IncrementByPercentage(percentage T) 
	Reset() T
	SetValue(value T)
}