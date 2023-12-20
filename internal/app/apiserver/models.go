package apiserver

import "log"

type AudioFile struct {
	Bucket      string `json:"bucket"`
	FileOne     string `json:"fileOne"`
	FileOneSize int32  `json:"fileOneSize"`
	FileTwo     string `json:"fileTwo,omitempty"`
	FileTwoSize int32  `json:"fileTwoSize,omitempty"`
	BitRate     int8   `json:"bitRate"`
	SampleRate  int32  `json:"sampleRate"`
	Chanels     int8   `json:"chanels"`
}

func NewAudioFile() *AudioFile {
	return &AudioFile{}
}

type WavFormat struct {
	riffTitle     [4]byte
	wavTitle      [4]byte
	chunkID       [4]byte
	chunkSize     uint32
	audioFormat   uint16
	numChannels   uint16
	sampleRate    uint32
	byteRate      uint32
	blockAlign    uint16
	bitsPerSample uint16
	fileOneSize   uint32
	fileTwoSize   uint32
	data          [4]byte
}

func NewWavFormat(a *AudioFile) *WavFormat {
	log.Println(a.FileOneSize)
	log.Println(uint32(a.FileOneSize))
	return &WavFormat{
		riffTitle:     [4]byte{'R', 'I', 'F', 'F'},
		wavTitle:      [4]byte{'W', 'A', 'V', 'E'},
		chunkID:       [4]byte{'f', 'm', 't', ' '},
		chunkSize:     16,
		audioFormat:   1,
		numChannels:   1,
		sampleRate:    uint32(a.SampleRate),
		byteRate:      uint32(a.SampleRate * int32(a.Chanels) * int32(int8(a.BitRate)>>4)),
		blockAlign:    uint16(a.Chanels * (int8(a.BitRate) >> 4)),
		bitsPerSample: uint16(a.BitRate),
		fileOneSize:   uint32(a.FileOneSize),
		fileTwoSize:   uint32(a.FileTwoSize),
		data:          [4]byte{'d', 'a', 't', 'a'},
	}
}
