package action

import (
	"fmt"
	"strings"

	"github.com/JonLMyers/GoFxell/game"
	"github.com/c-bata/go-prompt"
)

func cmdCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "show routes", Description: "show routes"},
		{Text: "exit", Description: "Exit CLI"},
		{Text: "deploy miner", Description: "deploy miner {Entropy/CPU/Io/Bandwidth}"},
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

func DeployMiner(ipAddr string, minerType string, team *game.Team) (bool, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if minerType == "Bandwidth" || minerType == "IO" || minerType == "CPU" || minerType == "Entropy" {
		if len(team.DiscoveredNodes[index].Miners) <= team.DiscoveredNodes[index].MaxMiners {
			miner := game.NewMiner(minerType, *team)
			team.DiscoveredNodes[index].Node.Miners = append(team.DiscoveredNodes[index].Node.Miners, miner)
			go miner.Mine(team, team.DiscoveredNodes[index].Node)
			return true, nil
		}
		fmt.Println("Not enough space on Target")
	}
	return false, nil

}

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
				fmt.Println(team.DiscoveredNodes[index].Miners)
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
