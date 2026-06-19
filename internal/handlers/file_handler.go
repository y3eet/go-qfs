package handlers

import (
	"fmt"
	"go-qfs/internal/config"
	"io"
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

func (h *FileHandler) UploadFile(c *gin.Context) {
	filename := c.GetHeader("X-Filename")
	uploadDir := config.Cfg.BaseDir + c.Query("path")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Filename header"})
		return
	}

	dstPath, err := safeDestPath(uploadDir, filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dst, err := os.Create(dstPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create file"})
		return
	}
	defer dst.Close()

	written, err := io.Copy(dst, c.Request.Body)
	if err != nil {
		os.Remove(dstPath)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to write file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"filename": filename, "size": written})
}

func safeDestPath(baseDir, filename string) (string, error) {
	clean := filepath.Base(filename) // strips any directory components
	if clean == "." || clean == ".." || clean == "" {
		return "", fmt.Errorf("invalid filename")
	}

	dst := filepath.Join(baseDir, clean)

	absBase, err := filepath.Abs(baseDir)
	if err != nil {
		return "", err
	}
	absDst, err := filepath.Abs(dst)
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(absDst, absBase+string(os.PathSeparator)) && absDst != absBase {
		return "", fmt.Errorf("invalid path")
	}

	return dst, nil
}
