package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilterStartNodes(t *testing.T) {
	gameMap := Map{Nodes: []Node{
		{
			IPAddr:    "127.0.0.1",
			StartNode: true,
		},
		{
			IPAddr: "10.0.0.1",
		},
	}}

	startNodes := FilterStartNodes(gameMap.Nodes)
	require.Len(t, startNodes, 1)
	require.Equal(t, "127.0.0.1", startNodes[0].IPAddr)
}

func TestStartNodeSelect(t *testing.T) {
	nodes := []Node{
		{
			IPAddr:    "127.0.0.1",
			StartNode: true,
		},
		{
			IPAddr:    "10.0.0.1",
			StartNode: true,
		},
	}
	startNodes := FilterStartNodes(nodes)
	require.Len(t, startNodes, 2)
	_, startNodes = startNodes.Select()
	require.Len(t, startNodes, 1)
}
