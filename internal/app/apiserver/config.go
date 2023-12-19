package apiserver

import (
	"io"
	"log"
	"os"
)

// Config Parse
type Config struct {
	BindAddr        string `json:"bind_addr"`
	MasterSeaweedFS string `json:"master_seaweedfs"`
	FilerSeaweedFS  string `json:"filer_seaweedfs"`
	PortSeaweedFS   string `json:"port_seaweedfs_filer"`
}

// Read Config file, return json string for unmarshal
func (c Config) ReadConfigFile(path string) string {
	log.Printf("read config file")
	fileConfig, err := os.Open(path)
	if err != nil {
		log.Printf("Error of reading config file, start with default config %s", err.Error())
		return ""
	}
	defer fileConfig.Close()

	data := make([]byte, 128)
	var jsonString string
	for {
		n, err := fileConfig.Read(data)
		if err == io.EOF {
			break
		}
		jsonString += string(data[:n])

	}
	return jsonString
}

func NewConfig() *Config {
	return &Config{
		BindAddr:        ":8085",
		MasterSeaweedFS: "seaweedfs",
		FilerSeaweedFS:  "seaweedfs",
		PortSeaweedFS:   ":8888",
	}
}
