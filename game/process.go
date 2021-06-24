package game

import (
	"fmt"

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
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return Process{}, err
	}
	//Change this to false by default
	proc := Process{team.Name, team.DiscoveredNodes[index].Node, uuid.New().String(), cmd, true, true}
	team.DiscoveredNodes[index].Processes = append(team.DiscoveredNodes[index].Processes, proc)
	return proc, nil
}

func (team *Team) ShowProcesses(ipAddr string, gameMap Map) ([]Process, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return nil, err
	}
	var procs []Process

	for _, proc := range team.DiscoveredNodes[index].Node.Processes {
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
