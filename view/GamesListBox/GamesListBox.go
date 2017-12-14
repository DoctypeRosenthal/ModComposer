package GamesListBox

import (
	. "ModComposer/types"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var ui struct {
	gamesList *walk.ListBox
}

type Config struct {
	Games []string
	OnSelectGameFromList func(int)
}

func Update(state AppState) {
	ui.gamesList.SetModel(state.Games)
	ui.gamesList.SetCurrentIndex(state.SelectedGameID)
}

func Create(cnf Config) ListBox {
	return ListBox{
		AssignTo: &ui.gamesList,
		Model: cnf.Games,
		OnSelectedIndexesChanged: func() {
			var i = ui.gamesList.CurrentIndex()
			if i > -1 {
				cnf.OnSelectGameFromList(i)
			}
		},
	}
}