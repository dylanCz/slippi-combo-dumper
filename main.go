package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dylanCz/slippi-combo-dumper/internal/encoder"
)

const (
	dolphinPath = "C:\\Users\\Maskl\\Appdata\\Roaming\\Slippi Launcher\\playback\\dolphin.exe"
	isoPath     = "C:\\Roms\\gc\\Super Smash Bros. Melee (USA) (En,Ja) (v1.02).iso"
	// videoDump   = "C:\\Users\\Maskl\\Appdata\\Roaming\\Slippi Launcher\\playback\\User\\Dump\\Frames\\framedump0.avi"
	// videoDump = "resources/framedump0.avi"
	// audioDump   = "C:\\Users\\Maskl\\Appdata\\Roaming\\Slippi Launcher\\playback\\User\\Dump\\Audio\\dspdump.wav"
	// audioDump = "resources/dspdump.wav"
	// output = "videos/output"
)

type Config struct {
	VideoPath string `json:"videoPath"`
	AudioPath string `json:"audioPath"`
	Quality   string `json:"quality"`
	Codec     string `json:"codec"`
	Output    string `json:"output"`
	Discord   bool   `json:"discord"`
}

func setConfig() Config {
	vPath := flag.String("videoPath", "resources/framedump0.avi", "The path to your frame dumped video")
	aPath := flag.String("audioPath", "resources/dspdump.wav", "The path to your frame dumped audio")
	quality := flag.String("quality", "Best", "The quality of the final clip, choices are Lossless, Best, Good, Average and Low")
	codec := flag.String("codec", "webm", "The codec to use, choices are webm or mp4")
	output := flag.String("output", "output", "The name prefix of the output clip")
	discord := flag.Bool("discord", false, "Whether or not you want the clip to be under 10mb for discord")

	flag.Parse()

	return Config{VideoPath: *vPath, AudioPath: *aPath, Quality: *quality, Codec: *codec, Output: *output, Discord: *discord}
}

func main() {
	// dolphinConf := dolphin.InitDolphin(dolphinPath, isoPath)
	// err := dolphin.Run(*dolphinConf)
	// fmt.Println(err)
	config := setConfig()

	if err := encoder.Encode(config.VideoPath, config.AudioPath, config.Output, config.Quality, config.Codec, config.Discord); err != nil {
		log.Fatalf("ffmpeg failed: %v", err)
	}
	fmt.Println("Done!")
}
