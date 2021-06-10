package action

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/JonLMyers/GoFxell/game"
	"github.com/c-bata/go-prompt"
)

const (
	MaxPID = 10000
	MinPID = 1000
)

type Processes struct {
	IpAddr string
	Procs  []Process
}

type Process struct {
	TeamName string
	PID      int
	CMD      string
}

func cmdCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "show routes", Description: "show routes"},
		{Text: "show proc", Description: "show processes"},
		{Text: "deploy miner", Description: "deploy miner {Entropy/CPU/Io/Bandwidth}"},
		{Text: "kill", Description: "kill {Process ID (PID)"},
		{Text: "exit", Description: "Exit CLI"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
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

func ShowProcesses(ipAddr string, team *game.Team, gameMap game.Map) ([]Process, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return nil, err
	}
	var procs []Process
	//I think we use a process interface here
	for _, miner := range team.DiscoveredNodes[index].Node.Miners {
		if miner.TeamName != team.Name {
			// 50/50 chance of finding not implemented
			procs = append(procs, Process{miner.TeamName, miner.PID, miner.CMD})
		}
		procs = append(procs, Process{miner.TeamName, miner.PID, miner.CMD})
	}

	return procs, nil
}

func KillProcess(ipAddr string, PID int, team *game.Team) (bool, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return false, err
	}
	if contains(team.DiscoveredNodes[index].Node.PIDS, PID) {
		for i, miner := range team.DiscoveredNodes[index].Node.Miners {
			if miner.PID == PID {
				// Removes the miner
				team.DiscoveredNodes[index].Node.Miners = append(team.DiscoveredNodes[index].Node.Miners[:i], team.DiscoveredNodes[index].Node.Miners[i+1:]...)
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
	return false, nil
}

func DeployMiner(ipAddr string, minerType string, team *game.Team) (bool, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if minerType == "Bandwidth" || minerType == "IO" || minerType == "CPU" || minerType == "Entropy" {
		if team.DiscoveredNodes[index].MinersDeployed <= team.DiscoveredNodes[index].MaxMiners {
			//Generate a new PID and ensure it is not in the PID list
			for {
				PID := rand.Intn(MaxPID-MinPID+1) + MinPID
				if !contains(team.DiscoveredNodes[index].PIDS, PID) {
					miner := game.NewMiner(minerType, *team, PID, fmt.Sprintf("/%s/miner.exe", team.Name))
					team.DiscoveredNodes[index].Node.Miners = append(team.DiscoveredNodes[index].Node.Miners, miner)
					fmt.Println("Process: miner.exe@", PID, " Deployed")

					go miner.Mine(team, index)
					team.DiscoveredNodes[index].MinersDeployed = team.DiscoveredNodes[index].MinersDeployed + 1
					team.DiscoveredNodes[index].PIDS = append(team.DiscoveredNodes[index].PIDS, PID)
					fmt.Println("Process: miner.Mine() Initiated")
					return true, nil
				}
			}
		}
		fmt.Println("Not enough space on Target")
	}
	return false, nil

}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// This will turn into middleware, I think.
func Connect(ipAddr string, team *game.Team, gameMap game.Map) (bool, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	n := team.DiscoveredNodes[index].Node
	if team.DiscoveredNodes[index].NodeOwned {
		playerName := "player"
		currentSystem := n.IPAddr
		gamePrompt := fmt.Sprintf("%s@%s:~$ ", playerName, currentSystem)
		for {
			cmd := prompt.Input(gamePrompt, cmdCompleter)
			cmd = strings.TrimSpace(cmd)
			cmdParts := strings.Split(cmd, " ")
			if cmd == "show routes" {
				ShowRoutes(n.IPAddr, team, gameMap)
				continue
			}
			if strings.HasPrefix(cmd, "deploy miner") {
				DeployMiner(n.IPAddr, cmdParts[2], team)
				continue
			}
			if strings.HasPrefix(cmd, "show proc") {
				processes, _ := ShowProcesses(n.IPAddr, team, gameMap)
				fmt.Println(processes)
				continue
			}
			if strings.HasPrefix(cmd, "kill") {
				i, err := strconv.Atoi(cmdParts[1])
				if err != nil {
					fmt.Println("Invalid PID")
					continue
				}
				ok, err := KillProcess(n.IPAddr, i, team)
				if !ok {
					fmt.Println(err)
					fmt.Println("Failed to kill process")
					continue
				}
				fmt.Println("Process Destroyed")
				continue
			}
			if cmd == "exit" {
				return true, nil
			}
			fmt.Println("Invalid Command")

		}
	} else {
		return false, nil
	}

}
