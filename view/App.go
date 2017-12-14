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
	removeBtn *walk.PushButton
	currentGamePath *walk.Label
}

type Events struct {
	OnChangeCurrentGamePath, OnDeleteFile, OnAddGame func(string)
	OnCopyFile                                       func(string, string)
	OnSelectGameFromList                             func(int)
	OnRemoveSelectedGameFromList                     func()
}

type Model struct {
	SelectedGameID int
	Games          GamesList
}

func Update(model Model) {
	ui.mw.SetTitle(createWindowTitle(model))
	ui.removeBtn.SetEnabled(model.SelectedGameID > -1)
	ui.currentGamePath.SetText(createGamePath(model))

	GamesListBox.Update(GamesListBox.Model{model.Games.Names(), model.SelectedGameID})
}

func createWindowTitle(model Model) string {
	var title = "ModComposer"
	if len(model.Games) > 0 {
		title = model.Games[model.SelectedGameID].Name + " - ModComposer"
	}
	return title
}

// creates the main window
func Create(model Model, evts Events) MainWindow {
	var wf = new(walk.Form)

	return MainWindow{
		AssignTo: &ui.mw,
		Title:    createWindowTitle(model),
		MinSize:  Size{600, 400},
		Layout:   HBox{},
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
					GamesListBox.Create(
						GamesListBox.Model{
							SelectedIndex: model.SelectedGameID,
							Games: model.Games.Names(),
						},
						GamesListBox.Events{ evts.OnSelectGameFromList, },
					),
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
										OnSelectDir: evts.OnAddGame,
									}).Run(*wf)
								},
							},
							PushButton{
								Text:      "- Delete",
								AssignTo:  &ui.removeBtn,
								Enabled:   false,
								MaxSize:   Size{Width: 60},
								OnClicked: evts.OnRemoveSelectedGameFromList,
							},
						},
					},

				},
			},
/*
			HSeparator{
				MinSize: Size{Height: 1},
			},
*/
			Composite{
				Layout: Grid{
					Rows:     3,
					MarginsZero: true,
				},
				Children: []Widget{
					Label {
						MaxSize: Size{Height:12},
						AssignTo: &ui.currentGamePath,
						Text: createGamePath(model),

					},
					PushButton{
						Text: "+ Select ModFolder",
						OnClicked: func() {

						},
					},
				},
			},
			VSplitter{
				Children: []Widget{
					PushButton{
						Text: "Kopiere Datei",
						OnClicked: func() {
							evts.OnCopyFile("./gist.txt", "./gist2.txt")
						},
					},
					PushButton{
						Text: "Lösche Datei",
						OnClicked: func() {
							evts.OnDeleteFile("./gist2.txt")
						},
					},
				},
			},
		},
	}
}
func createGamePath(model Model) string {
	var out = "Game path: "
	if len(model.Games) > 0 {
		out += model.Games[model.SelectedGameID].Path
	}
	return out

}
