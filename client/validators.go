package client

import (
	"github.com/JonLMyers/GoFxell/game"
)

func RequiredArguments(cmdParts []string, length int) (bool, []byte) {
	if len(cmdParts) < length {
		return false, Message("Invalid Number of Arguments. Check Command Syntax")
	}
	return true, nil
}

func DiscoveredNode(ipAddr string, team *game.Team) (bool, []byte) {
	_, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return false, Message("Invalid IP Address")
	}
	return true, nil
}

func OwnedNode(ipAddr string, team *game.Team) (bool, []byte) {
	i, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return false, Message("Invalid IP Address")
	}
	if !team.DiscoveredNodes[i].NodeOwned {
		return false, Message("No Remote Access Tool on Target. Run `exploit` before Connecting")
	}
	return true, nil
}
