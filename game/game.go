package game

import (
	"sync/atomic"
	"time"

	"log"
)


type Game struct{
	running atomic.Bool
	progress AtomicProgressBarSegments
}


func (game *Game) start(){
	game.running.Store(true)
	go game.runnable()
}
func (game *Game)runnable(){

	for game.IsRunning(){
		game.progress.IncrementByPercentage(-0.01)
		<-time.After(time.Second)
	}
	log.Println("Game ended!")
}

func (game *Game) Start() {
	game.start()
}

func (game *Game) HasStarted() bool {
	return game.running.Load()
}

func (game *Game) IsRunning() bool {
	return game.running.Load()
}


func (game *Game) GetGameData() *AtomicProgressBarSegments {
	return &game.progress
}
