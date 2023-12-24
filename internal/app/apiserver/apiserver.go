package apiserver

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type APIserver struct {
	Seaweedfs *Seaweedfs
	config    *Config
	router    *mux.Router
}

// Constructor
func New(config *Config, swfs *Seaweedfs) *APIserver {
	return &APIserver{
		config:    config,
		router:    mux.NewRouter(),
		Seaweedfs: swfs,
	}
}

// Start server
func (s *APIserver) Start() error {
	s.configRouter()
	log.Printf("Server started with addr %s", string(s.config.BindAddr))
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configRouter() {
	s.router.HandleFunc("/api/v1/pcmtowav", s.handleHello())
}

func (s *APIserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var wav bytes.Buffer
		audioConfig := NewAudioFile()
		err := json.Unmarshal([]byte(r.Header.Get("Data")), audioConfig)
		if err != nil {
			log.Printf("error unmarshal input json: %s", err.Error())
		}
		wavConfig := NewWavFormat(audioConfig)
		firstPCM := s.Seaweedfs.Download(audioConfig, 1)
		if audioConfig.Chanels == 1 {
			wav, err = MonoPCMtoWav(*firstPCM, wavConfig)
			if err != nil {
				log.Printf("error convert Monopcm2wav: %s", err.Error())
			}
		}
		log.Printf("THIS: %v", audioConfig.Chanels)
		if audioConfig.Chanels == 2 {
			secondPCM := s.Seaweedfs.Download(audioConfig, 2)
			wav, err = StereoPCMtoWav(*firstPCM, *secondPCM, wavConfig)
			if err != nil {
				log.Printf("error convert Stereopcm2wav: %s", err.Error())
			}
		}
		//wav, err := pcm2wav.ConvertBytes(firstPCM.Bytes(), int(audioFileConfig.Chanels), int(audioFileConfig.SampleRate), int(audioFileConfig.BitRate))
		//if err != nil {
		//	log.Printf("error convert pcm2wav: %s", err.Error())
		//}

		rez, err := s.Seaweedfs.Uploadtest(audioConfig, &wav)
		if err != nil {
			log.Printf("error upload to seaweedfs: %s", err.Error())
		}
		log.Printf("upload to seaweedfs: <%s> <%s> <%s>", rez.Name, rez.Size, rez.FileURL)

		file, err := os.Create("hello.wav") // test file
		defer file.Close()
		file.Write(wav.Bytes())

	}
}
