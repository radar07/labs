package main

import (
	"github.com/charmbracelet/huh"
)

func main() {
	// New Form
	allowedEnv := []string{"dev", "stag", "prod"}
	env := ""
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Select env").Options(huh.NewOptions(allowedEnv...)...).Value(&env),
		).WithHide(env != ""),
	).Run()

	if err != nil {
		panic(err)
	}

}
