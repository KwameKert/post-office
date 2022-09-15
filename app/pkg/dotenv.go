package pkg

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

func GodotEnv(key string) string {

	env := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		godotenv.Load(filepath.Join(".env"))
		env <- os.Getenv(key)
	} else {
		env <- os.Getenv(key)
	}
	return <-env
}
