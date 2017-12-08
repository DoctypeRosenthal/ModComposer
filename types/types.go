package types

import (
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


