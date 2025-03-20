package main

import (
	"git.sindadsec.ir/asm/backend/app"
	"git.sindadsec.ir/asm/backend/config"
)

func main() {
	config := config.Init()
	app := app.Init(config)
	app.Run()
}
