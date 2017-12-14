package PathDialog

import (
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/walk"
	"ModComposer/file"
	"log"
)
type Events struct {
	OnSelectDir func(string)
}

func Create(evt Events) Dialog {
	var ui struct {
		treeView *walk.TreeView
		LE *walk.LineEdit
		D *walk.Dialog
	}

	treeModel, err := file.NewDirectoryTreeModel()
	if err != nil {
		log.Fatal(err)
	}

	return Dialog{
		AssignTo: &ui.D,
		Visible: true,
		Title: "Choose game path",
		MinSize: Size{400, 647},
		Layout:  VBox{},
		Children: []Widget{
			LineEdit{
				AssignTo: &ui.LE,
			},
			TreeView{
				AssignTo: &ui.treeView,
				Model:    treeModel,
				ItemHeight: 22,
				OnCurrentItemChanged: func() {
					ui.LE.SetText(ui.treeView.CurrentItem().(*file.Directory).Path())
				},
			},
			PushButton{
				Text: "OK",
				Font: Font{
					Bold: true,
				},
				OnClicked: func() {
					evt.OnSelectDir(ui.LE.Text())
					ui.D.Close(0)
				},
			},
		},
	}
}
