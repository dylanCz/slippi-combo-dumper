package main

import (
	"fmt"
	"log"

	"github.com/dylanCz/slippi-combo-dumper/internal/encoder"
)

const (
	dolphinPath = "C:\\Users\\Maskl\\Appdata\\Roaming\\Slippi Launcher\\playback\\dolphin.exe"
	isoPath     = "C:\\Roms\\gc\\Super Smash Bros. Melee (USA) (En,Ja) (v1.02).iso"
	// videoDump   = "C:\\Users\\Maskl\\Appdata\\Roaming\\Slippi Launcher\\playback\\User\\Dump\\Frames\\framedump0.avi"
	videoDump = "resources/framedump0.avi"
	// audioDump   = "C:\\Users\\Maskl\\Appdata\\Roaming\\Slippi Launcher\\playback\\User\\Dump\\Audio\\dspdump.wav"
	audioDump = "resources/dspdump.wav"
	output    = "videos/output"

	chosenQuality = "Lossless"
	chosenCodec   = "Slower"
)

func main() {
	// dolphinConf := dolphin.InitDolphin(dolphinPath, isoPath)
	// err := dolphin.Run(*dolphinConf)
	// fmt.Println(err)

	if err := encoder.Encode(videoDump, audioDump, output, chosenQuality, chosenCodec); err != nil {
		log.Fatalf("ffmpeg failed: %v", err)
	}
	fmt.Println("Done!")
}
