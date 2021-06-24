package game

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
)

const (
	MaxPID = 10000
	MinPID = 1000
)

type Map struct {
	Nodes      []Node `json:"nodes"`
	startNodes []Node `json:"-"`
}

type Node struct {
	IPAddr           string    `json:"IP_addr"`
	Platform         string    `json:"platform"`
	Services         []string  `json:"services"`
	Routes           []string  `json:"routes"`
	StartNode        bool      `json:"start_node"`
	ObjectiveNode    bool      `json:"objective_node"`
	Objective        Objective `json:"objective"`
	LootRate         int       `json:"loot_rate"`
	Active           bool      `json:"active"`
	MaxMiners        int       `json:"max_miners"`
	NodeOwned        bool
	RedFootprint     int
	BlueFootprint    int
	Processes        []Process
	Miners           []Miner
	Monitors         []Monitor
	Logs             []Log
	FirewallStrength int
	MinersDeployed   int
}

func NewMap(json_path string) *Map {
	byteValue, err := ioutil.ReadFile(json_path)
	if err != nil {
		panic(err)
	}
	var gameMap Map
	if err := json.Unmarshal(byteValue, &gameMap); err != nil {
		panic(err)
	}
	return &gameMap
}

func (gameMap Map) FindNode(ipAddr string) (node Node) {
	for _, n := range gameMap.Nodes {
		if ipAddr == n.IPAddr {
			return n
		}
	}
	return
}

func (gameMap *Map) SelectStartNode() Node {
	if gameMap.startNodes == nil {
		gameMap.startNodes = FilterStartNodes(gameMap.Nodes)
	}

	if len(gameMap.startNodes) < 1 {
		panic(fmt.Errorf("cannot select team start node: no starting nodes remaining"))
	}
	max := big.NewInt(int64(len(gameMap.startNodes)))
	i, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}

	node := gameMap.startNodes[i.Int64()]

	gameMap.startNodes[i.Int64()] = gameMap.startNodes[len(gameMap.startNodes)-1] // Copy last node to index i
	gameMap.startNodes = gameMap.startNodes[:len(gameMap.startNodes)-1]           // Remove last node

	return node
}

func (team *Team) UpdateNodeFootprint(ipAddr string, value int) (int, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return 0, err
	}
	if team.Name == "Red" {
		team.DiscoveredNodes[index].RedFootprint = team.DiscoveredNodes[index].RedFootprint + value
		return team.DiscoveredNodes[index].RedFootprint, nil
	} else {
		team.DiscoveredNodes[index].BlueFootprint = team.DiscoveredNodes[index].BlueFootprint + value
		return team.DiscoveredNodes[index].BlueFootprint, nil
	}
}

func (team *Team) ViewNodeFootprint(ipAddr string) (int, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return 0, err
	}
	if team.Name == "Red" {
		return team.DiscoveredNodes[index].Node.RedFootprint, nil
	}
	return team.DiscoveredNodes[index].Node.BlueFootprint, nil
}

// func (n Node) String() string {
// 	return "You suck"
// }
