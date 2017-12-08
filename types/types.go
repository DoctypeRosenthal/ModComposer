package types

import (
	"fmt"
	"github.com/lxn/walk"
)

type AppState struct {
	GamePath string
}

type EventHandler struct {
	SelectGame, DeleteFile func(string)
	CopyFile               func(string, string)
}

type UiControls struct {
	OutTE *walk.TextEdit
}

type Directory struct {
	Name     string
	Parent   *Directory
	Children []*Directory
}

type DirectoryTreeModel struct {
	walk.TreeModelBase
	Roots []*Directory
}

func (d *Directory) ChildAt(index int) walk.TreeItem {
	return d.Children[index]
}

func (d *Directory) ChildCount() int {
	if d.Children == nil {
		// It seems this is the first time our child count is checked, so we
		// use the opportunity to populate our direct children.
		if err := d.ResetChildren(); err != nil {
			fmt.Print(err)
		}
	}

	return len(d.Children)
}

func (d *Directory) ResetChildren() error {
	d.Children = nil

	dirPath := d.Path()

	if err := filepath.Walk(d.Path(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if info == nil {
				return filepath.SkipDir
			}
		}

		name := info.Name()

		if !info.IsDir() || path == dirPath || shouldExclude(name) {
			return nil
		}

		d.Children = append(d.Children, NewDirectory(name, d))

		return filepath.SkipDir
	}); err != nil {
		return err
	}

	return nil
}
