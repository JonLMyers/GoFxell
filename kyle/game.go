package kyle

import (
	"sync"
	"time"
)

type Log struct {
	LoggedAt  time.Time
	DeletedAt time.Time
}

type Map struct {
	mu   sync.RWMutex
	logs []Log
}

func (m *Map) Log(log Log) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.logs = append(m.logs, log)
}

func (m *Map) CollectLogs(monitor Monitor) (logs []Log) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, log := range m.logs {
		if !log.DeletedAt.IsZero() && monitor.StartTime.After(log.DeletedAt) {
			continue
		}

		logs = append(logs, log)

	}
	return
}

type Monitor struct {
	StartTime time.Time
}

func (m *Monitor) CollectLogs() {
	m.StartTime = time.Now()
}
