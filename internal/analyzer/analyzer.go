package analyzer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yourusername/linux-process-monitor/internal/config"
)

type Analyzer struct {
	config *config.Config
}

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func NewAnalyzer(cfg *config.Config) *Analyzer {
	return &Analyzer{
		config: cfg,
	}
}

func (a *Analyzer) AnalyzeProcessBehavior(processInfo string) (string, error) {
	request := OllamaRequest{
		Model:  "llama2",
		Prompt: fmt.Sprintf("Analise o comportamento deste processo: %s", processInfo),
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(a.config.OllamaEndpoint+"/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.Response, nil
}
