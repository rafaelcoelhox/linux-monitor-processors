package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	MonitoringInterval int      `json:"monitoringInterval"` // em segundos
	CPUThreshold       float64  `json:"cpuThreshold"`
	MemoryThreshold    float64  `json:"memoryThreshold"`
	ProcessesToWatch   []string `json:"processesToWatch"`
	WhatsAppNumber     string   `json:"whatsappNumber"`
	OllamaEndpoint     string   `json:"ollamaEndpoint"`
}

func Load() (*Config, error) {
	file, err := os.Open("configs/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
