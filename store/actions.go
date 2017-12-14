package store

import . "ModComposer/types"

func SetGames(games []Game) {
	state.Games = games
	notify()
}

func AddGame(g Game) {
	state.Games = append(state.Games, g)
	notify()
}

func RemoveGame(i int) {
	s := state.Games
	state.Games = append(s[:i], s[i+1:]...)
	notify()
}

func ActivateGameFromList(i int) {
	state.SelectedGameID = i
	notify()
}

