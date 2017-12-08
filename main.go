package main

import (
	"ModComposer/file"
	"ModComposer/state"
	. "ModComposer/types"
	"ModComposer/view"
	"fmt"
)

/*
 * BUILD COMMANDS
 * without console: go build -ldflags="-H windowsgui -linkmode internal"
 * without gcc installed AND without console: go build -ldflags="-H windowsgui -linkmode internal"
 */
func main() {
	var evt = EventHandler{
		SelectGame: func(path string) {
			state.GamePath(path)
		},
		CopyFile: func(src, dst string) {
			fmt.Printf("Copying %s to %s\n", src, dst)
			err := file.Copy(src, dst)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("CopyFile succeeded!")
			}
		},
		DeleteFile: func(path string) {
			fmt.Printf("Deleting %s\n", path)
			err := file.Delete(path)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("File successfully deleted!")
			}
		},
	}

	state.Subscribe(view.Update)
	view.Render(evt)

}
