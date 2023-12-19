package apiserver

import (
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
		audioFileConfig := NewAudioFile()
		err := json.Unmarshal([]byte(r.Header.Get("Data")), audioFileConfig)
		if err != nil {
			log.Printf("error unmarshal input json: %s", err.Error())
		}
		firstPCM := s.Seaweedfs.Download(audioFileConfig, 1)
		//wav, err := pcm2wav.ConvertBytes(firstPCM.Bytes(), int(audioFileConfig.Chanels), int(audioFileConfig.SampleRate), int(audioFileConfig.BitRate))
		//if err != nil {
		//	log.Printf("error convert pcm2wav: %s", err.Error())
		//}
		file, err := os.Create("hello.wav")
		defer file.Close()
		file.Write(firstPCM.Bytes())
	}
}
