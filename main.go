package main

import (
	"context"
	"go-qfs/internal/config"
	"go-qfs/internal/server"
	"go-qfs/static"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")
	done <- true
}

func main() {
	config.Load()

	srv := server.NewServer()
	serverErr := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		log.Fatalf("http server failed to start: %s", err)
		os.Exit(0)
	case <-time.After(100 * time.Millisecond):
		// Server started OK
	}

	if len(os.Args) < 2 {
		app := NewApp()
		err := wails.Run(&options.App{
			Title:  "wails-app",
			Width:  1024,
			Height: 768,
			AssetServer: &assetserver.Options{
				Assets: static.Files,
			},
			BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
			OnStartup:        app.startup,
			Bind: []interface{}{
				app,
			},
		})
		if err != nil {
			log.Fatal("Error:", err)
		}
		return // Wails exited, app is done
	}

	// Server-only mode: wait for graceful shutdown
	done := make(chan bool, 1)
	go gracefulShutdown(srv, done)
	<-done
	log.Println("Graceful shutdown complete.")
}
