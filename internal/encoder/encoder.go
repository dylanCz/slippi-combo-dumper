package encoder

import (
	"fmt"

	"github.com/dylanCz/slippi-combo-dumper/internal/ffmpeg"
)

var qualityPresets = map[string]QualitySettings{
	"Lossless": {"18", "18", "1080", "64k", "1"},
	"Best":     {"20", "20", "960", "32k", "2"},
	"Good":     {"25", "23", "720", "32k", "3"},
	"Average":  {"30", "26", "540", "32k", "4"},
	"Low":      {"50", "30", "480", "16k", "5"},
}

var encoderPresets = map[string]EncoderPresets{
	"webm": {".webm", "libvpx-vp9"},
	"mp4":  {".mp4", "libx264"},
}

func Encode(videoDump, audioDump, output, quality, codec string, discord bool) error {
	preset, ok := qualityPresets[quality]
	if !ok {
		return fmt.Errorf("invalid quality: %s", quality)
	}
	encoder, ok := encoderPresets[codec]
	if !ok {
		return fmt.Errorf("invalid codec: %s", codec)
	}

	args := buildArgs(preset, encoder, videoDump, audioDump, output, quality)
	return ffmpeg.Run(args...)
}

func buildArgs(preset QualitySettings, encoder EncoderPresets, videoDump string, audioDump string, output string, chosenQuality string) []string {
	args := []string{
		"-i", videoDump,
		"-i", audioDump,
		"-vf", "scale=-2:" + preset.Resolution + ":flags=lanczos",
		"-c:v", encoder.Encoder,
		"-crf", preset.CRF(encoder.Encoder),
		"-b:v", "0",
		"-c:a", "libopus",
		"-b:a", preset.AudioBitrate,
		"-threads", "8",
	}

	switch encoder.Encoder {
	case "libvpx-vp9":
		args = append(args, "-row-mt", "1", "-cpu-used", preset.CPUUsed)
	case "libx264":
		args = append(args, "-preset", "fast")
	}

	args = append(args, output+"-"+chosenQuality+encoder.Filetype)

	return args
}
