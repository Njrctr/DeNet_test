package main

import "github.com/Njrctr/DeNet_test/internal/app"

const configsDir = "configs"

// @title DeNet API
// @version 1.0
// @description API Server for DeNet

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app.Run(configsDir)
}
