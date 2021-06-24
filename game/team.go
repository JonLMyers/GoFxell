package game

import (
	"fmt"
	"sync"
)

const (
	defaultBandwidth  = 100
	defaultIO         = 100
	defaultCPU        = 100
	defaultEntropy    = 100
	defaultFootprint  = 0
	defaultOpSecMeter = 0
)

type DiscoveredNode struct {
	*Node
	DiscoveredIPAddr   bool
	DiscoveredPlatform bool
	DiscoveredServices bool
	DiscoveredRoutes   bool
	NodeOwned          bool
}

// DiscoveredNodes is a slice of DiscoveredNodes with some helper functions.
type DiscoveredNodes []DiscoveredNode

// IndexOf returns the index of the discovered node with the provided IP address.
func (nodes DiscoveredNodes) IndexOf(ipAddr string) (int, error) {
	for i, node := range nodes {
		if ipAddr == node.IPAddr && node.Active {
			return i, nil
		}
	}

	return 0, fmt.Errorf("Node (%q) is not currently accessible", ipAddr)
}

type Team struct {
	Name              string
	Bandwidth         int
	Io                int
	Cpu               int
	Entropy           int
	OpSecMeter        int
	ObjectiveType     string
	ObjectiveComplete bool
	StartNode         Node
	DiscoveredNodes   DiscoveredNodes
	Monitors          []Monitor
	mutex             sync.RWMutex
}

func NewTeam(name string, objType string, gameMap *Map) *Team {
	startNode := gameMap.SelectStartNode()
	var mutex = sync.RWMutex{}
	team := &Team{name, defaultBandwidth, defaultIO, defaultCPU, defaultEntropy, defaultOpSecMeter, objType, false, startNode, []DiscoveredNode{}, []Monitor{}, mutex}

	if team.StartNode.IPAddr != "" {
		team.DiscoverNodeIP(team.StartNode)
	}
	return team
}

func (team Team) GetResources() [4]string {
	var resources [4]string
	resources[0] = fmt.Sprintf("Bandwidth: %d", team.Bandwidth)
	resources[1] = fmt.Sprintf("Disc I/O: %d", team.Io)
	resources[2] = fmt.Sprintf("CPU Cycles: %d", team.Cpu)
	resources[3] = fmt.Sprintf("Entropy: %d", team.Entropy)
	return resources

}

// View restricts what properties of a node are visible.
func (team Team) View(node *Node) (visibleNode Node) {
	// Determine if this node has been discovered.
	i, err := team.DiscoveredNodes.IndexOf(node.IPAddr)
	if err != nil {
		return
	}
	discoveredNode := &team.DiscoveredNodes[i]

	// We haven't discovered this node, so return the zero value
	if discoveredNode == nil { // len(team.DiscoveredNodes) < 1 || index < 0 || index >= len(team.DiscoveredNodes) {
		return
	}

	if discoveredNode.DiscoveredIPAddr {
		visibleNode.IPAddr = discoveredNode.IPAddr
	}
	if discoveredNode.DiscoveredPlatform {
		visibleNode.Platform = discoveredNode.Platform
	}
	if discoveredNode.DiscoveredServices {
		visibleNode.Services = discoveredNode.Services
	}
	if discoveredNode.DiscoveredRoutes {
		visibleNode.Routes = discoveredNode.Routes
	}
	if discoveredNode.NodeOwned {
		visibleNode.NodeOwned = discoveredNode.NodeOwned
	}
	if team.Name == "Red" {
		visibleNode.RedFootprint = discoveredNode.RedFootprint
	} else {
		visibleNode.BlueFootprint = discoveredNode.BlueFootprint
	}

	return
}

// DiscoverNodeIP enables a team to view the ip address for a node.
func (team *Team) DiscoverNodeIP(node Node) {
	// Find an existing node with the provided IP
	for _, n := range team.DiscoveredNodes {
		if node.IPAddr == n.IPAddr {
			println("Node already discovered: " + node.IPAddr)
			return
		}
	}
	team.DiscoveredNodes = append(team.DiscoveredNodes, DiscoveredNode{Node: &node, DiscoveredIPAddr: true})
}

// DiscoverNodeRoutes enables a team to view the routes for a node.
func (team *Team) DiscoverNodeRoutes(node Node) error {
	// Find an existing node with the provided IP
	index, err := team.DiscoveredNodes.IndexOf(node.IPAddr)
	if err != nil {
		fmt.Println("IP Address not in teams discovered nodes")
		return err
	}
	team.DiscoveredNodes[index].DiscoveredRoutes = true
	return nil

	//team.DiscoveredNodes = append(team.DiscoveredNodes, DiscoveredNode{Node: node, DiscoveredIPAddr: true, DiscoveredRoutes: true})
}

// DiscoverNodeRoutes enables a team to view the services for a node.
func (team *Team) DiscoverNodeServices(node Node) error {
	// Find an existing node with the provided IP
	index, err := team.DiscoveredNodes.IndexOf(node.IPAddr)
	if err != nil {
		fmt.Println("IP Address not in teams discovered nodes")
		return err
	}
	team.DiscoveredNodes[index].DiscoveredServices = true
	return nil
	//team.DiscoveredNodes = append(team.DiscoveredNodes, DiscoveredNode{Node: node, DiscoveredIPAddr: true, DiscoveredServices: true})
}

// DiscoverNodeRoutes enables a team to view the services for a node.
func (team *Team) DiscoverNodePlatform(node Node) error {
	// Find an existing node with the provided IP
	index, err := team.DiscoveredNodes.IndexOf(node.IPAddr)
	if err != nil {
		fmt.Println("IP Address not in teams discovered nodes")
		return err
	}
	team.DiscoveredNodes[index].DiscoveredPlatform = true
	return nil
	//team.DiscoveredNodes = append(team.DiscoveredNodes, DiscoveredNode{Node: node, DiscoveredIPAddr: true, DiscoveredPlatform: true})
}
