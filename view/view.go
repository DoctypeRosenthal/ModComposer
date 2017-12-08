package view

import (
	. "ModComposer/types"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strings"
	"fmt"
)

var ui UiControls

func Update(s AppState) {
	if ui.OutTE.Text() != s.GamePath {
		ui.OutTE.SetText(strings.ToUpper(s.GamePath))
	}
}

func Render(e EventHandler) {
	var inTE *walk.TextEdit
	var treeView *walk.TreeView

	defer MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &ui.OutTE, ReadOnly: true},
				},
			},
			TreeView{
				AssignTo: &treeView,
				//Model:    treeModel,
				OnCurrentItemChanged: func() {
					dir := treeView.CurrentItem()
					fmt.Println(dir)
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					e.SelectGame(inTE.Text())
				},
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Kopiere Datei",
						OnClicked: func() {
							e.CopyFile("./gist.txt", "./gist2.txt")
						},
					},
					PushButton{
						Text: "Lösche Datei",
						OnClicked: func() {
							e.DeleteFile("./gist2.txt")
						},
					},
				},
			},
		},
	}.Run()
}
