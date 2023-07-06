package controller

import (
	"github.com/dqu123/loop/constants"
	"github.com/google/uuid"
)

// In-memory data structures are used to avoid needing a database.
var (
	dataGamesMap   = map[string]GameData{}
	dataNumPlayers = 0
)

type GameData struct {
	GameHash  string
	PlayerMap map[string]PlayerData
}

type PlayerData struct {
	PlayerID int
}

func NewGameData() GameData {
	gameHash := uuid.NewString()[:constants.GameHashLength]
	return GameData{
		GameHash: gameHash,
	}
}

func NewPlayerData() PlayerData {
	dataNumPlayers++
	return PlayerData{
		PlayerID: dataNumPlayers - 1,
	}
}
