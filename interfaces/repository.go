package interfaces


type IClickable interface {
	Click()
}

type IGameRepository[T Float,I Index] interface {
	Click() I
	Value() T
	Segments() []T
}


