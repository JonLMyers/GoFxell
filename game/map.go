package game

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
)

type Map struct {
	Nodes      []Node `json:"nodes"`
	startNodes []Node `json:"-"`
}

func (m *Map) SelectStartNode() Node {
	if m.startNodes == nil {
		m.startNodes = FilterStartNodes(m.Nodes)
	}

	if len(m.startNodes) < 1 {
		panic(fmt.Errorf("cannot select team start node: no starting nodes remaining"))
	}
	max := big.NewInt(int64(len(m.startNodes)))
	i, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}

	node := m.startNodes[i.Int64()]

	m.startNodes[i.Int64()] = m.startNodes[len(m.startNodes)-1] // Copy last node to index i
	m.startNodes = m.startNodes[:len(m.startNodes)-1]           // Remove last node

	return node
}

type Node struct {
	IPAddr    string   `json:"IP_addr"`
	Platform  string   `json:"platform"`
	Services  []string `json:"services"`
	Routes    []string `json:"routes"`
	StartNode bool     `json:"start_node"`
	LootRate  int      `json:"loot_rate"`
	MaxMiners int      `json:"max_miners"`
	NodeOwned bool
	Miners    []Miner
}

// func (n Node) String() string {
// 	return "You suck"
// }

func FilterStartNodes(nodes []Node) (startNodes []Node) {
	for _, node := range nodes {
		if node.StartNode {
			startNodes = append(startNodes, node)
		}
	}
	return
}

func (gameMap Map) FindNode(ipAddr string) (node Node) {
	for _, n := range gameMap.Nodes {
		if ipAddr == n.IPAddr {
			return n
		}
	}
	return
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
