package game

import (
	"fmt"
)

const (
	ExfiltrateObjectiveTimeReduction = 10
)

func (team *Team) Scan(ipAddr string) (Node, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		//fmt.Println("IP Address not in teams discovered nodes")
		fmt.Println(err)
		return Node{}, err
	}

	team.DiscoverNodePlatform(*team.DiscoveredNodes[index].Node)
	team.DiscoverNodeServices(*team.DiscoveredNodes[index].Node)
	//fmt.Println("Successful Scan")
	return *team.DiscoveredNodes[index].Node, nil

}

// Make code more testable by making functions return data and not print text (no JSON)
func (team *Team) ShowRoutes(ipAddr string, gameMap Map) ([]string, error) {
	discoveredNodes := make([]DiscoveredNode, len(team.DiscoveredNodes))
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

func (team *Team) ShowTargets() []Node {
	targets := []Node{}
	for _, n := range team.DiscoveredNodes {
		targets = append(targets, team.View(n.Node))
	}
	return targets
}
