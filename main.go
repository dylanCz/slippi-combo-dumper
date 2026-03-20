package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dylanCz/slippi-combo-dumper/internal/config"
	"github.com/dylanCz/slippi-combo-dumper/internal/encoder"
)

func setConfig() config.FfmpegConfig {
	vPath := flag.String("videoPath", "resources/framedump0.avi", "The path to your frame dumped video")
	aPath := flag.String("audioPath", "resources/dspdump.wav", "The path to your frame dumped audio")
	quality := flag.String("quality", "Best", "The quality of the final clip, choices are Lossless, Best, Good, Average and Low")
	codec := flag.String("codec", "webm", "The codec to use, choices are webm or mp4")
	output := flag.String("output", "output", "The name prefix of the output clip")
	discord := flag.Bool("discord", false, "Whether or not you want the clip to be under 10mb for discord")

	flag.Parse()

	return config.FfmpegConfig{VideoPath: *vPath, AudioPath: *aPath, Quality: *quality, Codec: *codec, Output: *output, Discord: *discord}
}

func main() {
	if err := encoder.Encode(setConfig()); err != nil {
		log.Fatalf("ffmpeg failed: %v", err)
	}
	fmt.Println("Done!")
}
