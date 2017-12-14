package main

import (
	"ModComposer/file"
	"ModComposer/store"
	. "ModComposer/types"
	"ModComposer/view"
	"fmt"
)

const (
	PATH_TO_STATE = "./state.json"
)

func check(e error, success string) {
	if e != nil {
		panic(e)
	} else {
		fmt.Println(success)
	}
}

/*
 * BUILD COMMANDS
 * without console: go build -ldflags="-H windowsgui -linkmode internal"
 * without gcc installed AND without console: go build -ldflags="-H windowsgui -linkmode internal"
 */
func main() {
	var initState AppState
	var err = file.JsonToObject(PATH_TO_STATE, &initState)
	check(err, "State loaded from "+PATH_TO_STATE)

	store.Initialise(initState)
	state := store.GetState()
	var model = view.Model{state.SelectedGameID, state.Games}
	var evts = view.Events{
		OnSelectGameFromList: func(i int) {
			store.ActivateGameFromList(i)
		},
		OnAddGame: func(path string) {
			store.AddGame(Game{ Name: "", Path: path, ModsDir: "" })
			state := store.GetState()
			store.ActivateGameFromList(len(state.Games)-1)
			defer file.ObjectToJson(state, PATH_TO_STATE)
		},
		OnRemoveSelectedGameFromList: func() {
			state := store.GetState()
			i := state.SelectedGameID

			if i == len(state.Games)-1 {
				// select 2nd last game before deleting last
				store.ActivateGameFromList(i-1)
			}
			store.RemoveGame(i)
			defer file.ObjectToJson(state, PATH_TO_STATE)
		},
		OnCopyFile: func(src, dst string) {
			fmt.Printf("Copying %s to %s\n", src, dst)
			err := file.Copy(src, dst)
			check(err, "CopyFile succeeded!")
		},
		OnDeleteFile: func(path string) {
			fmt.Printf("Deleting %s\n", path)
			err := file.Delete(path)
			check(err, "File successfully deleted!")
		},
	}

	store.Subscribe(func(state AppState) {
		view.Update(view.Model{state.SelectedGameID, state.Games})
	})
	view.Create(model, evts).Run()
}
