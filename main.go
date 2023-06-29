package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

const (
	hashLength = 10
)

type GameData struct {
	GameHash  string
	PlayerMap map[string]struct{}
}

func NewGameData() GameData {
	gameHash := uuid.NewString()[:hashLength]
	return GameData{
		GameHash: gameHash,
	}
}

func getHost() string {
	return "localhost:8000"
}

var (
	gamesMap = map[string]GameData{}
)

func handleGames(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		postGame(w, r)
	}
}

type PostGameResponse struct {
	URL string
}

func postGame(w http.ResponseWriter, r *http.Request) {
	game := NewGameData()
	gamesMap[game.GameHash] = game

	resBytes, err := json.Marshal(PostGameResponse{
		URL: fmt.Sprintf("%s/games/%s", getHost(), game.GameHash),
	})
	if err != nil {
		w.WriteHeader(500)
		return
	}
	io.WriteString(w, string(resBytes))
}

func postPlayer(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/games", handleGames)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("ERROR in http.ListenAndServe: ", err)
	}
}
