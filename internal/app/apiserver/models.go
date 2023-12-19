package apiserver

type AudioFile struct {
	Bucket     string `json:"bucket"`
	FileOne    string `json:"fileOne"`
	FileTwo    string `json:"fileTwo,omitempty"`
	BitRate    int8   `json:"bitRate"`
	SampleRate int32  `json:"sampleRate"`
	Chanels    int8   `json:"chanels"`
}

func NewAudioFile() *AudioFile {
	return &AudioFile{}
}
