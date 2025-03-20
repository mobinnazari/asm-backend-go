package main

import (
	"git.sindadsec.ir/asm/backend/app"
	"git.sindadsec.ir/asm/backend/config"
)

// @title			ASM Backend
// @version		1.0
// @license.name	Private Company License
// @BasePath		/
func main() {
	config := config.Init()
	app := app.Init(config)
	app.Run()
}
