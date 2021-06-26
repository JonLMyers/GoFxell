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
	mutex         *sync.RWMutex
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

func CheckMonitor(ipAddr string, team *Team) ([]Log, error) {
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

func (team *Team) NewMonitor(monitorType string, process *Process) *Monitor {
	monitor := Monitor{MonitorType: monitorType, StartTime: time.Now(), LastLog: time.Time{}, Process: process, CollectedLogs: nil, TeamName: team.Name, mutex: &sync.RWMutex{}}
	return &monitor
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
				if log.Type == "Network" && log.Value >= MonitorLogThreshold && !log.Deleted && log.Timestamp.After(monitor.LastLog) {
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

func (monitor *Monitor) MonitorProc(team *Team) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if !monitor.Process.Alive {
				return
			}
			// Get host nodes Network logs
			for _, proc := range monitor.Process.Node.Processes {
				// Check to see if a team tripped 100% footprint
				GetLogs(monitor.Process.Node.IPAddr, team)
				if proc.Viewable {
					monitor.mutex.Lock()
					defer monitor.mutex.Unlock()
					monitor.CollectedLogs = append(monitor.CollectedLogs, Log{Id: proc.PID, Type: "Process", LogString: proc.CMD, Timestamp: time.Now()})
					monitor.LastLog = time.Now()
					monitor.mutex.Unlock()
				}
			}
		}
	}
}
