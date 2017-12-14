package types

type AppState struct {
	SelectedGameID int
	Games          GamesList
}

type GamesList []Game

type Game struct {
	Name    string
	Path    string
	ModsDir string
	ModsEnabled []int
}

func (games GamesList) GameNames() []string {
	var out []string
	for _, game := range games {
		out = append(out, game.Name)
	}
	return out
}