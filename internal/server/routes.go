package server

import (
	"go-qfs/internal/handlers"
	"go-qfs/static"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-Filename"},
		AllowCredentials: true,
	}))

	fileHandler := handlers.NewFileHandler()
	api := r.Group("/api")
	{
		api.POST("/file/upload", fileHandler.UploadFile)
		api.GET("/file/download/*filepath", fileHandler.DownloadFile)
		api.GET("/files", fileHandler.GetFiles)
	}

	sub, _ := fs.Sub(static.Files, "dist")
	r.Use(serveEmbedded(sub))

	return r
}

func serveEmbedded(fsys fs.FS) gin.HandlerFunc {
	fileServer := http.FileServer(http.FS(fsys))

	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// Strip leading slash, default to index.html for root
		filePath := strings.TrimPrefix(path, "/")
		if filePath == "" {
			filePath = "index.html"
		}

		// Try to open the file in the embedded FS
		f, err := fsys.Open(filePath)
		if err != nil {
			// SPA fallback — serve index.html for unknown paths
			c.FileFromFS("index.html", http.FS(fsys))
			c.Abort()
			return
		}

		stat, err := f.Stat()
		f.Close()

		if err != nil {
			c.FileFromFS("index.html", http.FS(fsys))
			c.Abort()
			return
		}

		// If it's a directory, serve index.html instead of letting
		// http.FileServer redirect to a trailing slash (301)
		if stat.IsDir() {
			c.FileFromFS("index.html", http.FS(fsys))
			c.Abort()
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
