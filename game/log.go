package game

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	id        string
	timestamp time.Time
	Type      string
	LogString string
	Value     int
	Deleted   bool
	DeletedAt time.Time
}

func NewLog(ipAddr string, logType string, details string, value int, team *Team) (Log, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return Log{}, err
	}

	log := Log{uuid.New().String(), time.Now(), logType, details, value, false, time.Time{}}
	team.DiscoveredNodes[index].Node.Logs = append(team.DiscoveredNodes[index].Node.Logs, log)
	team.UpdateNodeFootprint(ipAddr, value)
	return log, nil
}

func GetLogs(ipAddr string, team *Team) ([]Log, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return nil, err
	}
	return team.DiscoveredNodes[index].Node.Logs, nil
}

func DeleteLog(logId string, ipAddr string, team *Team) (bool, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return false, err
	}
	logs := team.DiscoveredNodes[index].Node.Logs
	for i, log := range logs {
		if log.id == logId {
			if team.Name == "Red" {
				team.DiscoveredNodes[index].RedFootprint = team.DiscoveredNodes[index].RedFootprint - log.Value
			} else {
				team.DiscoveredNodes[index].BlueFootprint = team.DiscoveredNodes[index].BlueFootprint - log.Value
			}
			logs[i].Deleted = true
			logs[i].DeletedAt = time.Now()
			return true, nil
		}
	}
	return false, nil
}
