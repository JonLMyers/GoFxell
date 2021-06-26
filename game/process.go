package game

import (
	"github.com/google/uuid"
)

type ProcessList struct {
	IpAddr string
	Procs  []Process
}

type Process struct {
	TeamName string
	Node     *Node
	PID      string
	CMD      string
	Viewable bool
	Alive    bool
}

func (team *Team) NewProcess(ipAddr string, cmd string) (Process, error) {
	index, _ := team.DiscoveredNodes.IndexOf(ipAddr)

	proc := Process{team.Name, team.DiscoveredNodes[index].Node, uuid.New().String(), cmd, false, true}
	team.DiscoveredNodes[index].Processes = append(team.DiscoveredNodes[index].Processes, proc)
	return proc, nil
}

func (team *Team) ShowProcesses(ipAddr string, gameMap Map) ([]Process, error) {
	index, _ := team.DiscoveredNodes.IndexOf(ipAddr)
	var procs []Process

	for _, proc := range team.DiscoveredNodes[index].Node.Processes {
		if proc.TeamName == "Red" && team.DiscoveredNodes[index].Node.BlueFootprint >= 100 {
			if proc.TeamName == "Blue" {
				proc.Viewable = true
			}
		}
		if proc.TeamName == "Blue" && team.DiscoveredNodes[index].Node.RedFootprint >= 100 {
			if proc.TeamName == "Red" {
				proc.Viewable = true
			}
		}
		if proc.Viewable || proc.TeamName == team.Name {
			procs = append(procs, proc)
		}
	}
	return procs, nil
}

func (team *Team) KillProcess(ipAddr string, PID string) (bool, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return false, err
	}

	for _, proc := range team.DiscoveredNodes[index].Node.Processes {
		if proc.PID == PID {
			proc.Alive = false
			//team.DiscoveredNodes[index].Node.Processes = append(team.DiscoveredNodes[index].Node.Processes[:i], team.DiscoveredNodes[index].Node.Processes[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}
