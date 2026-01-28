package music

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	musicDir := "./music"

	// Get limit from query parameter (default 25)
	limitStr := c.Query("limit", "25")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 25
	}
	if limit > 100 {
		limit = 100 // Max 100
	}

	files, err := os.ReadDir(musicDir)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to read music directory",
		})
	}

	var musicFiles []string
	for _, file := range files {
		if !file.IsDir() {
			// Only include audio files
			ext := filepath.Ext(file.Name())
			if ext == ".mp3" || ext == ".wav" || ext == ".flac" || ext == ".m4a" || ext == ".ogg" {
				musicFiles = append(musicFiles, file.Name())
				if len(musicFiles) >= limit {
					break
				}
			}
		}
	}

	return c.JSON(musicFiles)
}

func GetFile(c *fiber.Ctx) error {
	filename := c.Params("name")

	// Security: prevent path traversal
	if strings.Contains(filename, "..") || strings.Contains(filename, "/") {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid filename",
		})
	}

	filepath := "./music/" + filename

	// Check if file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return c.Status(404).JSON(fiber.Map{
			"error": "Music file not found",
		})
	}

	return c.SendFile(filepath)
}
