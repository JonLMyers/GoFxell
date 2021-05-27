package game

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
)

type Map struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	IPAddr    string   `json:"IP_addr"`
	Platform  string   `json:"platform"`
	Services  []string `json:"services"`
	Routes    []string `json:"routes"`
	StartNode bool     `json:"start_node"`
	LootRate  int      `json:"loot_rate"`
}

type StartNodes []Node

func (nodes StartNodes) Select() (Node, StartNodes) {
	if len(nodes) < 1 {
		panic(fmt.Errorf("cannot select team start node: no starting nodes remaining"))
	}
	max := big.NewInt(int64(len(nodes)))
	i, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}

	node := nodes[i.Int64()]

	nodes[i.Int64()] = nodes[len(nodes)-1] // Copy last node to index i
	nodes = nodes[:len(nodes)-1]           // Remove last node

	return node, nodes
}

func FilterStartNodes(nodes []Node) (startNodes StartNodes) {
	for _, node := range nodes {
		if node.StartNode {
			startNodes = append(startNodes, node)
		}
	}
	return
}

func NewMap(json_path string) Map {
	byteValue, err := ioutil.ReadFile(json_path)
	if err != nil {
		panic(err)
	}
	var gameMap Map
	if err := json.Unmarshal(byteValue, &gameMap); err != nil {
		panic(err)
	}
	return gameMap
}
