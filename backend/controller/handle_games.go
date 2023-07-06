package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/dqu123/loop/constants"
	"github.com/dqu123/loop/logger"
)

func handleGames(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		postGame(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.LogError("handleGamesPlayers", errors.New("unimplemented method"))
	}
}

type PostGameResponse struct {
	URL string
}

func postGame(w http.ResponseWriter, r *http.Request) {
	game := NewGameData()
	dataGamesMap[game.GameHash] = game

	resBytes, err := json.Marshal(PostGameResponse{
		URL: fmt.Sprintf("%s/games/%s", constants.GetHost(), game.GameHash),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(resBytes))
}
