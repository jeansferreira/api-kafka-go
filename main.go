package main

import (
	"collect-server/initializers"
)

func init() {
	initializers.InitializeEnvironment()
}

func main() {
	app := initializers.InitializeServer()
	app.Listen(":8080")
}
