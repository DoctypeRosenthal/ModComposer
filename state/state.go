package state

import (
	. "ModComposer/types"
)

var s = AppState{
	GamePath: "",
}

var subscribers [](func(AppState))

func GamePath(path string) {
	s.GamePath = path
	notify()
}

func GetGamePath(path string) string {
	return s.GamePath
}

func Subscribe(fn func(AppState)) {
	subscribers = append(subscribers, fn)
}

func notify() {
	for _, sub := range subscribers {
		sub(s)
	}
}
