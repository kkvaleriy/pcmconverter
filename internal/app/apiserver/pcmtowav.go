package apiserver

import (
	"bytes"
	"encoding/binary"
)

func createWavTitle(wav *bytes.Buffer, wavConfig *WavFormat) *bytes.Buffer {
	binary.Write(wav, binary.BigEndian, wavConfig.riffTitle)
	if wavConfig.numChannels == 1 {
		binary.Write(wav, binary.LittleEndian, uint32(36+wavConfig.fileOneSize))
	} else {
		binary.Write(wav, binary.LittleEndian, uint32(36+wavConfig.fileTwoSize))
	}
	binary.Write(wav, binary.BigEndian, wavConfig.wavTitle)
	binary.Write(wav, binary.BigEndian, wavConfig.chunkID)
	binary.Write(wav, binary.LittleEndian, wavConfig.chunkSize)
	binary.Write(wav, binary.LittleEndian, wavConfig.audioFormat)
	binary.Write(wav, binary.LittleEndian, wavConfig.numChannels)
	binary.Write(wav, binary.LittleEndian, wavConfig.sampleRate)
	binary.Write(wav, binary.LittleEndian, wavConfig.byteRate)
	binary.Write(wav, binary.LittleEndian, wavConfig.blockAlign)
	binary.Write(wav, binary.LittleEndian, wavConfig.bitsPerSample)
	binary.Write(wav, binary.BigEndian, wavConfig.data)
	if wavConfig.numChannels == 1 {
		binary.Write(wav, binary.LittleEndian, wavConfig.fileOneSize)
	} else {
		binary.Write(wav, binary.LittleEndian, wavConfig.fileOneSize+wavConfig.fileTwoSize)
	}
	return wav
}

func MonoPCMtoWav(pcm bytes.Buffer, wavConfig *WavFormat) (bytes.Buffer, error) {
	wav := bytes.NewBuffer([]byte{})
	wav = createWavTitle(wav, wavConfig)
	err := binary.Write(wav, binary.LittleEndian, pcm.Bytes())
	return *wav, err
}

func StereoPCMtoWav(pcm1, pcm2 bytes.Buffer, wavConfig *WavFormat) (bytes.Buffer, error) {
	fileOneSize := int(wavConfig.fileOneSize)
	fileTwoSize := int(wavConfig.fileTwoSize)
	var (
		count_max int
		err       error
	)
	wav := bytes.NewBuffer([]byte{})
	wav = createWavTitle(wav, wavConfig)
	if fileOneSize > fileTwoSize {
		count_max = int(wavConfig.fileOneSize)
	} else {
		count_max = int(wavConfig.fileTwoSize)
	}
	//log.Println(fileOneSize)
	for i := 0; i < count_max; i++ {
		if i < fileOneSize {
			err = binary.Write(wav, binary.LittleEndian, pcm1.Bytes()[i])
		} else {
			err = binary.Write(wav, binary.LittleEndian, uint8(0))
		}
		if i < fileTwoSize {
			err = binary.Write(wav, binary.LittleEndian, pcm2.Bytes()[i])
		} else {
			err = binary.Write(wav, binary.LittleEndian, uint8(0))
		}
	}

	return *wav, err
}
