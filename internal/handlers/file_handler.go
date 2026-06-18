package handlers

import (
	"fmt"
	"go-qfs/internal/config"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
	path := c.Query("path")
	entries, err := os.ReadDir(config.Cfg.BaseDir + path)
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

func (h *FileHandler) DownloadFile(c *gin.Context) {
	fmt.Printf("Received download request for: %s\n", c.Param("filepath"))
	relativePath := strings.TrimPrefix(c.Param("filepath"), "/")
	baseDir := config.Cfg.BaseDir
	filename := filepath.Base(relativePath)

	absBase, _ := filepath.Abs(baseDir)
	absFile, err := filepath.Abs(filepath.Join(baseDir, relativePath))

	if err != nil || !strings.HasPrefix(absFile, absBase+string(filepath.Separator)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid path"})
		return
	}
	if _, err := os.Stat(absFile); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}

	c.FileAttachment(absFile, filename)
}
