package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"slices"

	api "github.com/mrfastwind/the-button-go/api"
	"github.com/mrfastwind/the-button-go/interfaces"

	game "github.com/mrfastwind/the-button-go/game"
)


type GameRepository struct{
	progress *game.Game
}

func (game *GameRepository) Click() uint8{
	var value = game.progress.GetGameData().Reset()
	var segments = game.progress.GetGameData().Segments()
	slices.Sort(segments)

	for index, threashold:=range segments {
		if value < threashold{
			return uint8(index)
		}
	}
	return uint8(math.NaN())
}

func (game *GameRepository) Value() float32{
	return game.progress.GetGameData().Value()
}

func (game *GameRepository) Segments() []float32{
	return game.progress.GetGameData().Segments()
}

func main() {


	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var gameptr *game.Game = game.CreateGame()

	var gamerepo interfaces.IGameRepository[float32,uint8] = &GameRepository{
		progress: gameptr,
	}

	var buttonService = api.ButtonService{
	Repository : gamerepo,
	}

	fs := http.FileServer(http.Dir("static"))


	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "templates/index.html")})
	api.AddRoutes("/api", buttonService)
	fmt.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}