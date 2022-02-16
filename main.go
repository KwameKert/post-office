package main

import (
	"postoffice/app"
	"postoffice/app/core"
)

func main() {
	config := core.NewConfig()
	app := &app.App{}
	app.Start(config)

}
