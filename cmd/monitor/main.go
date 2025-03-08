package main

import (
	"log"

	"github.com/yourusername/linux-process-monitor/internal/config"
	"github.com/yourusername/linux-process-monitor/internal/monitoring"
	"github.com/yourusername/linux-process-monitor/internal/whatsapp"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	whatsClient, err := whatsapp.NewClient(cfg)
	if err != nil {
		log.Fatalf("Erro ao inicializar cliente WhatsApp: %v", err)
	}

	monitor := monitoring.NewProcessMonitor(cfg, whatsClient)
	monitor.Start()
}
