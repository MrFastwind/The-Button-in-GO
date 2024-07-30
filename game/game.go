package game

import (
	"errors"
	"sync/atomic"
	"time"

	"log"
)

type GameConfig struct{
	percentageBySecond float32
}

type Game struct{
	running atomic.Bool
	progress AtomicProgressBarSegments
	config GameConfig
}


func (game *Game) start() (error){
	if game.HasStarted(){
		return errors.New("game already started")
	}
	game.running.Store(true)
	go game.runnable()
	return nil
}
func (game *Game)runnable(){
	game.progress.SetValue(1)
	game.running.Store(true)
	log.Println("Game started!")
	for game.IsRunning(){
		game.progress.IncrementByPercentage(game.config.percentageBySecond)
		<-time.After(time.Second)
		if(game.progress.Value()<=0){
			game.running.Store(false)
		}
	}
	log.Println("Game ended!")
}

func (game *Game) Start() (error){
	return game.start()
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
