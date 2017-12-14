package store

import (
	. "ModComposer/types"
)

var state AppState

var subscribers []func(AppState)

func GetState() AppState {
	return state
}

func Initialise(initState AppState) {
	state = initState
}

func Subscribe(fn func(AppState)) {
	subscribers = append(subscribers, fn)
}

func notify() {
	for _, sub := range subscribers {
		sub(state)
	}
}
