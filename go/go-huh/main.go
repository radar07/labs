package main

import (
	"path/filepath"

	"github.com/charmbracelet/huh"
)

func main() {
	path, err := filepath.Abs(".")

	if err != nil {
		panic(err)
	}

	// New Form
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewFilePicker().
				Picking(true).
				CurrentDirectory(path).
				DirAllowed(false).
				FileAllowed(true).
				Cursor("->").
				Height(2).
				ShowHidden(true).
				Value(&path),
		),
	).Run()

	// New Confirm
	err = huh.NewConfirm().
		Description("Do you want to proceed?").
		Run()
}
