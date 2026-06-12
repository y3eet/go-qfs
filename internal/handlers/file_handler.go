package handlers

import (
	"go-qfs/internal/config"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type FileHandler struct{}
type FileInfo struct {
	Name  string `json:"name"`
	IsDir bool   `json:"is_dir"`
	Size  int64  `json:"size"`
	Ext   string `json:"ext,omitempty"`
}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

func (h *FileHandler) GetFiles(c *gin.Context) {
	entries, err := os.ReadDir(config.Cfg.BaseDir)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read directory"})
		return
	}

	files := make([]FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		files = append(files, FileInfo{
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
			Size:  info.Size(),
			Ext:   filepath.Ext(entry.Name()),
		})
	}
	c.JSON(200, files)
}
