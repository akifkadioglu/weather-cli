package main

import (
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	ENV_KEY = "KEY"
)

func initEnv() {
	godotenv.Load(filepath.Join(".env"))
}
