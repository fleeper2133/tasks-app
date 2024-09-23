package main

import "github.com/fleeper2133/tasks-app/internal/app"

const (
	configPath = "configs"
)

// @title           TasksApp API
// @version         1.0
// @description     This is api for tasks

// @host      127.0.0.1:8000
// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization

func main() {
	app.Run(configPath)
}
