package game

type Log struct {
	id        int
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
	log := Log{len(team.DiscoveredNodes[index].Node.Logs), logType, details, team.Name, value}
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
