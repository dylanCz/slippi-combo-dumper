package encoder

import (
	"fmt"

	"github.com/dylanCz/slippi-combo-dumper/internal/config"
	"github.com/dylanCz/slippi-combo-dumper/internal/ffmpeg"
)

var qualityPresets = map[string]QualitySettings{
	"Lossless": {CRFvp9: "18", CRFx264: "18", Resolution: "1080", AudioBitrate: "64k", CPUUsed: "1"},
	"Best":     {CRFvp9: "20", CRFx264: "20", Resolution: "960", AudioBitrate: "32k", CPUUsed: "2"},
	"Good":     {CRFvp9: "25", CRFx264: "23", Resolution: "720", AudioBitrate: "32k", CPUUsed: "3"},
	"Average":  {CRFvp9: "30", CRFx264: "26", Resolution: "540", AudioBitrate: "32k", CPUUsed: "4"},
	"Low":      {CRFvp9: "50", CRFx264: "30", Resolution: "480", AudioBitrate: "16k", CPUUsed: "5"},
}

var encoderPresets = map[string]EncoderPresets{
	"webm": {Filetype: ".webm", Encoder: "libvpx-vp9"},
	"mp4":  {Filetype: ".mp4", Encoder: "libx264"},
}

func Encode(config config.Config) error {
	preset, ok := qualityPresets[config.Quality]
	if !ok {
		return fmt.Errorf("invalid quality: %s", config.Quality)
	}
	encoder, ok := encoderPresets[config.Codec]
	if !ok {
		return fmt.Errorf("invalid codec: %s", config.Codec)
	}

	args := buildArgs(preset, encoder, config)
	return ffmpeg.Run(args...)
}

func buildArgs(preset QualitySettings, encoder EncoderPresets, config config.Config) []string {
	args := []string{
		"-i", config.VideoPath,
		"-i", config.AudioPath,
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

	outputPath := config.Output + "-" + config.Quality + encoder.Filetype
	args = append(args, outputPath)

	return args
}
