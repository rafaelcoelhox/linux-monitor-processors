package monitoring

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/yourusername/linux-process-monitor/internal/config"
	"github.com/yourusername/linux-process-monitor/internal/whatsapp"
)

type ProcessMonitor struct {
	config   *config.Config
	whatsapp *whatsapp.Client
}

func NewProcessMonitor(cfg *config.Config, whatsClient *whatsapp.Client) *ProcessMonitor {
	return &ProcessMonitor{
		config:   cfg,
		whatsapp: whatsClient,
	}
}

func (pm *ProcessMonitor) Start() {
	ticker := time.NewTicker(time.Duration(pm.config.MonitoringInterval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		pm.checkProcesses()
	}
}

func (pm *ProcessMonitor) checkProcesses() {
	processes, err := process.Processes()
	if err != nil {
		return
	}

	for _, proc := range processes {
		name, err := proc.Name()
		if err != nil {
			continue
		}

		if pm.shouldMonitorProcess(name) {
			pm.analyzeProcess(proc)
		}
	}
}

func (pm *ProcessMonitor) shouldMonitorProcess(name string) bool {
	for _, p := range pm.config.ProcessesToWatch {
		if p == name {
			return true
		}
	}
	return false
}

func (pm *ProcessMonitor) analyzeProcess(p *process.Process) {
	cpu, _ := p.CPUPercent()
	mem, _ := p.MemoryPercent()

	if cpu > pm.config.CPUThreshold || float64(mem) > pm.config.MemoryThreshold {
		name, _ := p.Name()
		message := fmt.Sprintf("Alerta: Processo %s está utilizando CPU: %.2f%% e Memória: %.2f%%", name, cpu, mem)
		pm.whatsapp.SendMessage(message)
	}
}
