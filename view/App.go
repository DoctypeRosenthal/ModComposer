package view

import (
	. "ModComposer/types"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/walk"
	"ModComposer/view/PathDialog"
	"ModComposer/view/GamesListBox"
)

var ui struct {
	mw        *walk.MainWindow
	disabledBtn *walk.PushButton
}

type Config struct {
	SelectedGameID int
	Games          GamesList

	OnChangeCurrentGamePath, OnDeleteFile, OnAddGame func(string)
	OnCopyFile                                       func(string, string)
	OnSelectGameFromList                             func(int)
	OnRemoveSelectedGameFromList                     func()
}

func Update(state AppState) {
	currGame := state.Games[state.SelectedGameID]
	ui.mw.SetTitle(currGame.Path+" - ModComposer")
	ui.disabledBtn.SetEnabled(state.SelectedGameID > -1)
	GamesListBox.Update(state)
}

// creates the main window
func Create(cnf Config) MainWindow {
	var wf = new(walk.Form)

	return MainWindow{
		AssignTo: &ui.mw,
		Title:   "ModComposer",
		MinSize: Size{600, 400},
		Layout:  HBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{
					Columns: 1,
					MarginsZero: true,
				},
				Children: []Widget{
					Label {
						Text: "Games:",
					},
					GamesListBox.Create(GamesListBox.Config{
						cnf.Games.GameNames(),
						cnf.OnSelectGameFromList,
					}),
					Composite{
						Layout: Grid{
							Columns:2,
							MarginsZero: true,
						},
						Children: []Widget{
							PushButton{
								Text: "+ Add",
								MaxSize: Size{Width: 50},
								OnClicked: func() {
									PathDialog.Create(PathDialog.Events{
										OnSelectDir: cnf.OnAddGame,
									}).Run(*wf)
								},
							},
							PushButton{
								Text:      "- Delete",
								AssignTo:  &ui.disabledBtn,
								Enabled:   false,
								MaxSize:   Size{Width: 60},
								OnClicked: cnf.OnRemoveSelectedGameFromList,
							},
						},
					},

				},
			},
			HSeparator{
				MinSize: Size{Height: 1},
			},

			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Kopiere Datei",
						OnClicked: func() {
							cnf.OnCopyFile("./gist.txt", "./gist2.txt")
						},
					},
					PushButton{
						Text: "Lösche Datei",
						OnClicked: func() {
							cnf.OnDeleteFile("./gist2.txt")
						},
					},
				},
			},
		},
	}
}
