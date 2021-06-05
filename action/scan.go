package action

import (
	"fmt"

	"github.com/JonLMyers/GoFxell/game"
)

// There is also a version that likes on team but I like this one more :)
func Scan(ipAddr string, team *game.Team) (bool, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println("IP Address not in teams discovered nodes")
		return false, err
	}

	team.DiscoverNodePlatform(team.DiscoveredNodes[index].Node)
	team.DiscoverNodeServices(team.DiscoveredNodes[index].Node)
	fmt.Println("Successful Scan")
	return true, nil

}
