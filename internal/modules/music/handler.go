package music

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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

	// Security: Strict path validation
	cleanName := filepath.Base(filepath.Clean(filename))
	if cleanName == "." || cleanName == ".." || cleanName == "/" {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid filename"})
	}

	targetPath := filepath.Join("./music", cleanName)

	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		return c.Status(404).JSON(fiber.Map{"error": "Music file not found"})
	}

	return c.SendFile(targetPath)
}

func Stream(c *fiber.Ctx) error {
	filename := c.Params("name")

	// Security: Strict path validation
	cleanName := filepath.Base(filepath.Clean(filename))
	if cleanName == "." || cleanName == ".." || cleanName == "/" {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid filename"})
	}

	sourceFile := filepath.Join("./music", cleanName)

	// Check if source file exists and is actually in the music dir
	absMusicDir, _ := filepath.Abs("./music")
	absSourceFile, _ := filepath.Abs(sourceFile)
	if !strings.HasPrefix(absSourceFile, absMusicDir) {
		return c.Status(403).JSON(fiber.Map{"error": "Access denied"})
	}

	if _, err := os.Stat(sourceFile); os.IsNotExist(err) {
		return c.Status(404).JSON(fiber.Map{"error": "Music file not found"})
	}

	hlsDir := filepath.Join("./storage/hls/music", cleanName)
	masterPlaylist := filepath.Join(hlsDir, "master.m3u8")

	// Generate HLS if not exists
	if _, err := os.Stat(masterPlaylist); os.IsNotExist(err) {
		if err := os.MkdirAll(hlsDir, 0755); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create HLS directory"})
		}

		// Fail-fast: Use context with timeout for FFmpeg
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		cmd := exec.CommandContext(ctx, "ffmpeg", "-i", sourceFile,
			"-map", "0:a", "-b:a:0", "64k",
			"-map", "0:a", "-b:a:1", "128k",
			"-map", "0:a", "-b:a:2", "320k",
			"-f", "hls",
			"-hls_time", "6",
			"-hls_playlist_type", "vod",
			"-master_pl_name", "master.m3u8",
			"-var_stream_map", "a:0 a:1 a:2",
			filepath.Join(hlsDir, "audio_%v.m3u8"))

		if err := cmd.Run(); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error":   "Failed to generate HLS segments",
				"details": "FFmpeg execution failed or timed out",
			})
		}
	}

	// Serve the requested file
	requestedFile := c.Params("*")
	if requestedFile == "" || requestedFile == "*" {
		requestedFile = "master.m3u8"
	}
	// Clean requested file to prevent traversal (only allow top-level files in hlsDir)
	cleanRequested := filepath.Base(filepath.Clean(requestedFile))

	return c.SendFile(filepath.Join(hlsDir, cleanRequested))
}
