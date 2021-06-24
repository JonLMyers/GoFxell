package game

import (
	"fmt"
	"sync"
	"time"
)

const (
	MonitorLogThreshold = 10
)

type Monitor struct {
	mutex         sync.RWMutex
	StartTime     time.Time
	LastLog       time.Time
	MonitorType   string
	Process       *Process
	CollectedLogs []Log
	TeamName      string
}

const (
	NetworkMonitorFootprintValue = 5
)

func CheckMonitor(ipAddr string, team *Team, PID int) ([]Log, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, monitor := range team.DiscoveredNodes[index].Node.Monitors {
		if monitor.TeamName == team.Name {
			fmt.Println(monitor.CollectedLogs)
			//monitor.DumpMonitor()
			return monitor.CollectedLogs, nil
		}
	}
	return nil, nil
}

func (team *Team) NewMonitor(monitorType string, process *Process) Monitor {
	monitor := Monitor{MonitorType: monitorType, StartTime: time.Now(), LastLog: time.Time{}, Process: process, CollectedLogs: nil, TeamName: team.Name}
	return monitor
}

func (monitor *Monitor) MonitorNetwork(team *Team) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if !monitor.Process.Alive {
				return
			}
			// Get host nodes Network logs
			for _, log := range monitor.Process.Node.Logs {
				if log.Type == "Network" && log.Value >= MonitorLogThreshold && !log.Deleted && log.timestamp.After(monitor.LastLog) {
					monitor.mutex.Lock()
					defer monitor.mutex.Unlock()
					monitor.CollectedLogs = append(monitor.CollectedLogs, log)
					monitor.LastLog = time.Now()
					monitor.mutex.Unlock()
				}
			}
		}
	}
}
