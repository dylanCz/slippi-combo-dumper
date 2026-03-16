package main

import (
	"fmt"
	"log"

	"github.com/dylanCz/slippi-combo-dumper/internal/ffmpeg"
)

const (
	dolphinPath = "C:\\Users\\Maskl\\Appdata\\Roaming\\Slippi Launcher\\playback\\dolphin.exe"
	isoPath     = "C:\\Roms\\gc\\Super Smash Bros. Melee (USA) (En,Ja) (v1.02).iso"
	// videoDump   = "C:\\Users\\Maskl\\Appdata\\Roaming\\Slippi Launcher\\playback\\User\\Dump\\Frames\\framedump0.avi"
	videoDump = "resources/framedump0.avi"
	// audioDump   = "C:\\Users\\Maskl\\Appdata\\Roaming\\Slippi Launcher\\playback\\User\\Dump\\Audio\\dspdump.wav"
	audioDump = "resources/dspdump.wav"
	output    = "output50.webm"

	chosenQuality = "Low"
)

func main() {
	// dolphinConf := dolphin.InitDolphin(dolphinPath, isoPath)
	// err := dolphin.Run(*dolphinConf)
	// fmt.Println(err)
	qualityMap := map[string]string{
		"Best":    "18",
		"Good":    "25",
		"Average": "30",
		"Low":     "50",
	}
	quality, ok := qualityMap[chosenQuality]
	if !ok {
		// handle invalid input
		log.Fatalf("invalid quality: %s", chosenQuality)
	}
	err := ffmpeg.Run("-i", videoDump, "-i", audioDump, "-vf", "scale=-2:720:flags=lanczos", "-c:v", "libvpx-vp9", "-crf", quality, "-b:v", "0", "-c:a", "libopus", "-b:a", "32k", output)
	fmt.Println(err)
	fmt.Println("Done!")
}
