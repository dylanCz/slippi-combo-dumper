package encoder

type QualitySettings struct {
	CRFvp9       string
	CRFx264      string
	Resolution   string
	AudioBitrate string
	CPUUsed      string
}

type EncoderPresets struct {
	Filetype string
	Encoder  string
}

func (q QualitySettings) CRF(codec string) string {
	if codec == "libx264" {
		return q.CRFx264
	}
	return q.CRFvp9
}
