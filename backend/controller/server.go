package controller

import "net/http"

func NewServer() error {
	http.HandleFunc("/games", handleGames)
	http.HandleFunc("/games/players", handleGamesPlayers)
	return http.ListenAndServe(":8000", nil)
}
