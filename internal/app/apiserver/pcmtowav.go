package apiserver

import (
	"bytes"
	"encoding/binary"
)

func MonoPCMtoWav(pcm bytes.Buffer, wavConfig *WavFormat) (bytes.Buffer, error) {
	wav := bytes.NewBuffer([]byte{})
	binary.Write(wav, binary.BigEndian, wavConfig.riffTitle)
	binary.Write(wav, binary.LittleEndian, uint32(36+wavConfig.fileOneSize))
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
	binary.Write(wav, binary.LittleEndian, wavConfig.fileOneSize)

	err := binary.Write(wav, binary.LittleEndian, pcm.Bytes())
	return *wav, err
}
