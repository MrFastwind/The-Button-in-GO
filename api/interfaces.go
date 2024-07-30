package api

import (
	"github.com/mrfastwind/the-button-go/interfaces"
)

type ButtonService struct {
	Repository interfaces.IGameRepository[float32,uint8]
}
