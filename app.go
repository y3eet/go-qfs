package main

import (
	"context"
	"fmt"
	"go-qfs/internal/config"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Test() {
	fmt.Println("Test")
}

type ServerInfo struct {
	Port string `json:"port"`
	IP   string `json:"ip"`
	Env  string `json:"env"`
	Host string `json:"host"`
}

func (a *App) GetServerInfo() ServerInfo {
	return ServerInfo{
		Port: config.Cfg.Port,
		IP:   config.Cfg.Port,
		Env:  config.Cfg.AppEnv,
		Host: fmt.Sprintf("%s:%s", config.Cfg.IP, config.Cfg.Port),
	}

}
