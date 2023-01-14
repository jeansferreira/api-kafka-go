package initializers

import (
	"collect-server/env"

	dotEnv "github.com/joho/godotenv"
)

func InitializeEnvironment() {
	dotEnv.Load()
	env.SetupEnv()
}
