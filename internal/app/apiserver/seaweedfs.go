package apiserver

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/linxGnu/goseaweedfs"
)

type Seaweedfs struct {
	Seaweedfs goseaweedfs.Seaweed
	Filer     goseaweedfs.Filer
	url       string
}

func NewSeaweedFS(config *Config) *Seaweedfs {
	_filer := make([]string, 1)
	_filer = append(_filer, config.FilerSeaweedFS)
	_sw, err := goseaweedfs.NewSeaweed(config.MasterSeaweedFS, _filer, 8096, &http.Client{Timeout: 5 * time.Minute})
	if err != nil {
		log.Printf("error create seaweedfs: %s", err.Error())
	}
	return &Seaweedfs{
		Seaweedfs: *_sw,
		Filer:     *_sw.Filers()[0],
		url:       config.FilerSeaweedFS + config.PortSeaweedFS + "/",
	}
}

func (swfs *Seaweedfs) Download(audioFile *AudioFile, chanel int8) *bytes.Buffer {
	var (
		buf      bytes.Buffer
		filename string
	)
	if chanel == 1 {
		filename = audioFile.FileOne
	} else {
		filename = audioFile.FileTwo
	}
	swfs.Filer.Download(swfs.url+audioFile.Bucket+filename, nil, func(r io.Reader) error {
		_, err := io.Copy(&buf, r)
		if err != nil {
			log.Printf("error download file %s", err.Error())
		}
		return nil
	})
	log.Println(swfs.url + audioFile.Bucket + filename)
	return &buf
}

func (swfs *Seaweedfs) Uploadtest(audioFile *AudioFile, wav *bytes.Buffer) (*goseaweedfs.FilerUploadResult, error) {
	rez, err := swfs.Filer.Upload(wav, 1048576, swfs.url+audioFile.Bucket+audioFile.FileOne+".wav", "col", "")
	log.Println(int64(wav.Available()))
	log.Println(audioFile.Bucket + audioFile.FileOne + ".wav")
	return rez, err
}
