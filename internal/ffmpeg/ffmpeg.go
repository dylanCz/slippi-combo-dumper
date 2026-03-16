// internal/ffmpeg/ffmpeg.go
package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
)

func Run(args ...string) error {
	path, err := exec.LookPath("ffmpeg")
	if err != nil {
		return fmt.Errorf("ffmpeg not found on PATH: %w", err)
	}

	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
