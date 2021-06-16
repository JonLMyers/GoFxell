package game

import "github.com/google/uuid"

type Log struct {
	id        string
	Type      string
	LogString string
	TeamName  string
	Value     int
}

func CreateLog(ipAddr string, logType string, details string, value int, team *Team) error {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return err
	}

	log := Log{uuid.New().String(), logType, details, team.Name, value}
	team.DiscoveredNodes[index].Node.Logs = append(team.DiscoveredNodes[index].Node.Logs, log)
	team.UpdateNodeFootprint(ipAddr, value)
	return nil
}

func ViewLogs(ipAddr string, team *Team) ([]Log, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return nil, err
	}
	return team.DiscoveredNodes[index].Node.Logs, nil
}

func CleanLog(logId string, ipAddr string, team *Team) (bool, error) {
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
			//Removes the log from the slice
			logs[i] = logs[len(logs)-1]
			team.DiscoveredNodes[index].Node.Logs = logs[:len(logs)-1]
			return true, nil
		}
	}
	return false, nil
}

func (team Team) UpdateNodeFootprint(ipAddr string, value int) (int, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return 0, err
	}
	if team.Name == "Red" {
		team.DiscoveredNodes[index].RedFootprint = team.DiscoveredNodes[index].RedFootprint + value
		return team.DiscoveredNodes[index].RedFootprint, nil
	} else {
		team.DiscoveredNodes[index].BlueFootprint = team.DiscoveredNodes[index].BlueFootprint + value
		return team.DiscoveredNodes[index].BlueFootprint, nil
	}
}

func (team Team) UpdateOpSec() {
	//I am not sure what this will do so no implementation as of now
}
