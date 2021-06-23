package action

import (
	"fmt"
	"math/rand"

	"github.com/JonLMyers/GoFxell/game"
)

const (
	MaxPID                           = 10000
	MinPID                           = 1000
	ExfiltrateObjectiveTimeReduction = 10
)

type ProcessList struct {
	IpAddr string
	Procs  []ViewableProcess
}

type ViewableProcess struct {
	TeamName string
	PID      int
	CMD      string
}

// Make code more testable by making functions return data and not print text (no JSON)
func ShowRoutes(ipAddr string, team *game.Team, gameMap game.Map) ([]string, error) {
	discoveredNodes := make([]game.DiscoveredNode, len(team.DiscoveredNodes))
	copy(discoveredNodes, team.DiscoveredNodes)

	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return nil, err
	}
	routes := team.DiscoveredNodes[index].Node.Routes
	team.DiscoveredNodes[index].DiscoveredRoutes = true
	for _, route := range routes {
		for i, node := range gameMap.Nodes {
			if route == node.IPAddr {
				team.DiscoverNodeIP(gameMap.Nodes[i])
			}
		}

	}
	return routes, nil
}

func ShowTargets(team game.Team) {
	for _, n := range team.DiscoveredNodes {
		fmt.Println(team.View(n.Node))
	}
}

func ShowProcesses(ipAddr string, team *game.Team, gameMap game.Map) ([]ViewableProcess, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return nil, err
	}
	var procs []ViewableProcess
	//Code for miners which can be lumped into the next section
	/*for _, miner := range team.DiscoveredNodes[index].Node.Miners {
		if miner.TeamName != team.Name {
			// 50/50 chance of finding not implemented
			// Increase the opposite teams Footprint?
			procs = append(procs, Process{miner.TeamName, miner.PID, miner.CMD})
		}
		procs = append(procs, Process{miner.TeamName, miner.PID, miner.CMD})
	}*/

	for _, proc := range team.DiscoveredNodes[index].Node.Processes {
		procs = append(procs, ViewableProcess{proc.TeamName, proc.PID, proc.CMD})
	}
	return procs, nil
}

func ViewNodeFootprint(ipAddr string, team *game.Team) (int, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return 0, err
	}
	if team.Name == "Red" {
		return team.DiscoveredNodes[index].Node.RedFootprint, nil
	}
	return team.DiscoveredNodes[index].Node.BlueFootprint, nil
}

func KillProcess(ipAddr string, PID int, team *game.Team) (bool, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return false, err
	}
	if contains(team.DiscoveredNodes[index].Node.PIDS, PID) {

		for i, proc := range team.DiscoveredNodes[index].Node.Processes {
			if proc.PID == PID {
				team.DiscoveredNodes[index].Node.Processes = append(team.DiscoveredNodes[index].Node.Processes[:i], team.DiscoveredNodes[index].Node.Processes[i+1:]...)
				break
			}
		}
		for i, p := range team.DiscoveredNodes[index].Node.PIDS {
			if p == PID {
				// Removes the PID
				team.DiscoveredNodes[index].Node.PIDS = append(team.DiscoveredNodes[index].Node.PIDS[:i], team.DiscoveredNodes[index].Node.PIDS[i+1:]...)
				return true, nil
			}
		}
	}
	// Miners are special and need to be explicitly removed due to go routine functionality
	/*for i, miner := range team.DiscoveredNodes[index].Node.Miners {
		if miner.PID == PID {
			// Removes the miner
			team.DiscoveredNodes[index].Node.Miners = append(team.DiscoveredNodes[index].Node.Miners[:i], team.DiscoveredNodes[index].Node.Miners[i+1:]...)
			break
		}
	}*/
	return false, nil
}

//Validation must be included to ensure that players do not spawn multiple exfil procs
func DeployExfiltrateObjective(ipAddr string, team *game.Team) (int, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	if team.DiscoveredNodes[index].Node.ObjectiveNode {
		for {
			PID := rand.Intn(MaxPID-MinPID+1) + MinPID
			if !contains(team.DiscoveredNodes[index].PIDS, PID) {
				game.CreateProcess(ipAddr, team, PID, "GetFiles.exe")
				fmt.Println("Process: GetFiles.exe@", PID, " Deployed")
				team.DiscoveredNodes[index].PIDS = append(team.DiscoveredNodes[index].PIDS, PID)
				fmt.Println("Process: GetFiles.Exfiltrate() Initiated")
				team.UpdateNodeFootprint(team.DiscoveredNodes[index].Node.IPAddr, 20)
				return PID, nil
			}
		}
	}
	return 0, nil
}

// Refactor this due to the new process system :)
func DeployMiner(ipAddr string, minerType string, team *game.Team) (bool, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if minerType == "Bandwidth" || minerType == "IO" || minerType == "CPU" || minerType == "Entropy" {
		// This is probably a bug as the miner limit applies to all teams. Maybe I want this. idk.
		if team.DiscoveredNodes[index].MinersDeployed <= team.DiscoveredNodes[index].MaxMiners {
			//Generate a new PID and ensure it is not in the PID list
			PID, _ := CreatePid(ipAddr, team, 5)
			miner := game.NewMiner(minerType, *team, PID, fmt.Sprintf("/%s/miner.exe", team.Name))
			team.DiscoveredNodes[index].Node.Miners = append(team.DiscoveredNodes[index].Node.Miners, miner)
			game.CreateProcess(ipAddr, team, PID, "miner.exe")
			fmt.Println("Process: miner.exe@", PID, " Deployed")
			go miner.Mine(team, index)
			team.DiscoveredNodes[index].MinersDeployed = team.DiscoveredNodes[index].MinersDeployed + 1
			fmt.Println("Process: miner.Mine() Initiated")
			return true, nil
		}
		fmt.Println("Not enough space on Target")
	}
	return false, nil

}
func CreatePid(ipAddr string, team *game.Team, footprintValue int) (int, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	for {
		PID := rand.Intn(MaxPID-MinPID+1) + MinPID
		if !contains(team.DiscoveredNodes[index].PIDS, PID) {
			team.DiscoveredNodes[index].PIDS = append(team.DiscoveredNodes[index].PIDS, PID)
			team.UpdateNodeFootprint(team.DiscoveredNodes[index].Node.IPAddr, footprintValue)
			fmt.Println(team.DiscoveredNodes[index].PIDS)
			return PID, nil
		}
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
