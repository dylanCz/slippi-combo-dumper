package config

type FfmpegConfig struct {
	VideoPath string
	AudioPath string
	Quality   string
	Codec     string
	Output    string
	Discord   bool
}
