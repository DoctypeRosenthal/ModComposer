package GamesListBox

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var ui struct {
	gamesList *walk.ListBox
}

type Events struct {
	OnSelectGame func(int)
}

type Model struct {
	Games []string
	SelectedIndex int
}

func Update(m Model) {
	ui.gamesList.SetModel(m.Games)
	ui.gamesList.SetCurrentIndex(m.SelectedIndex)
}

func Create(model Model, evts Events) ListBox {
	return ListBox{
		AssignTo: &ui.gamesList,
		Model:    model.Games,
		CurrentIndex: model.SelectedIndex,
		OnSelectedIndexesChanged: func() {
			var i = ui.gamesList.CurrentIndex()
			if i > -1 {
				evts.OnSelectGame(i)
			}
		},
	}
}