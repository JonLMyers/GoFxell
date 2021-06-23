package action

import (
	"fmt"

	"github.com/JonLMyers/GoFxell/game"
)

const (
	FirewallFootprintValue = 20
)

func DeployFirewall(ipAddr string, team *game.Team) (bool, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	for _, proc := range team.DiscoveredNodes[index].Node.Processes {
		if proc.CMD == "firewall.exe" {
			if proc.TeamName == team.Name {
				fmt.Println("Firewall Already deployed")
				return false, nil
			}
		}
	}
	PID, ok := CreatePid(ipAddr, team, FirewallFootprintValue)
	game.CreateProcess(ipAddr, team, PID, "firewall.exe")
	if ok != nil {
		return false, ok
	}
	team.DiscoveredNodes[index].Node.FirewallStrength = 100
	fmt.Println("Process: firewall.exe@", PID, " Deployed")
	fmt.Println("Process: firewall.Defend() Initiated")
	//team.DiscoveredNodes[index].Node.Firewalls = append(team.DiscoveredNodes[index].Node.Firewalls, team.Name)
	return true, nil
}

func FirewallExists(index int, team *game.Team) bool {
	for _, process := range team.DiscoveredNodes[index].Node.Processes {
		if process.CMD == "firewall.exe" {
			return true
		} else {
			return false
		}
	}
	return false
}
