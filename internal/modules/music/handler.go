package music

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	musicDir := "./music"

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
			}
		}
	}

	return c.JSON(musicFiles)
}
