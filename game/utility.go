package game

import "fmt"

func FilterStartNodes(nodes []Node) (startNodes []Node) {
	for _, node := range nodes {
		if node.StartNode {
			startNodes = append(startNodes, node)
		}
	}
	return
}

func containsPid(proclist []Process, PID string) bool {
	for _, process := range proclist {
		if process.PID == PID {
			return true
		}
	}
	return false
}

func (team *Team) ProcessExists(ipAddr string, cmd string) bool {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return false
	}
	for _, proc := range team.DiscoveredNodes[index].Node.Processes {
		if proc.CMD == "netmon.exe" {
			if proc.TeamName == team.Name {
				//fmt.Println("Network Monitor Already deployed")
				return false
			}
		}
	}
	return true
}
