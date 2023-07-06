package controller

import (
	"net/http"

	"github.com/dqu123/loop/logger"
)

func NewServer() error {
	http.HandleFunc("/games", handleGames)
	http.HandleFunc("/games/players", handleGamesPlayers)
	logger.LogInfo("Starting server...")
	return http.ListenAndServe(":8000", nil)
}
