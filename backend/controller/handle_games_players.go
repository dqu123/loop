package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/dqu123/loop/logger"
)

func handleGamesPlayers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		httpStatus, externalMessage, err := postPlayer(w, r)
		if err != nil {
			w.WriteHeader(httpStatus)
			io.WriteString(w, externalMessage)
			logger.LogError("handleGamesPlayers", err)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.LogError("handleGamesPlayers", errors.New("unimplemented method"))
	}

}

type PostPlayerRequest struct {
	GameHash   string `json:"gameHash"`
	PlayerUUID string `json:"playerUuid"`
}

func postPlayer(w http.ResponseWriter, r *http.Request) (int, string, error) {
	decoder := json.NewDecoder(r.Body)
	var req PostPlayerRequest
	err := decoder.Decode(&req)
	switch err {
	case nil:
		break
	case io.EOF:
		return http.StatusBadRequest, "Error: Request body required", errors.New("missing request body")
	default:
		return http.StatusInternalServerError, "", err
	}

	game, ok := dataGamesMap[req.GameHash]
	if !ok {
		return http.StatusInternalServerError, "", fmt.Errorf("gameHash:%s not found", req.GameHash)
	}
	game.PlayerMap[req.PlayerUUID] = NewPlayerData()

	logger.LogInfo(fmt.Sprintf("DATA GAMES MAP: %#v", dataGamesMap))
	return http.StatusOK, "", nil
}
