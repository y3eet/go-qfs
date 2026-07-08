package main

import (
	"context"
	"fmt"
	"go-qfs/internal/config"
	"os/exec"
	"runtime"
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
	OS   string `json:"os"`
}

func (a *App) GetServerInfo() ServerInfo {
	return ServerInfo{
		Port: config.Cfg.Port,
		IP:   config.Cfg.Port,
		Env:  config.Cfg.AppEnv,
		Host: fmt.Sprintf("%s:%s", config.Cfg.IP, config.Cfg.Port),
		OS:   runtime.GOOS,
	}

}

func (a *App) ExecCommand(command string) error {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Running on macOS.")
	case "linux":
		terminals := [][]string{
			{"gnome-terminal", "--", "bash", "-c", command + "; exec bash"},
			{"konsole", "-e", "bash", "-c", command + "; exec bash"},
			{"xterm", "-e", "bash", "-c", command + "; exec bash"},
		}
		for _, t := range terminals {
			if _, err := exec.LookPath(t[0]); err == nil {
				return exec.Command(t[0], t[1:]...).Start()
			}
		}
		return fmt.Errorf("no terminal emulator found")
	case "windows":
		return exec.Command("cmd", "/C", "start", "powershell", "-NoExit", "-Command", command).Start()
	default:
		fmt.Printf("Running on unsupported OS: %s\n", os)
	}
	return fmt.Errorf("unsupported platform")
}
