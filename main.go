package main

import (
	"./application"
)

func main() {
	app := application.NewApplication()
	app.Run()
}