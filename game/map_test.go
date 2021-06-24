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
	gameMap := &Map{
		Nodes: []Node{
			{
				IPAddr:    "127.0.0.1",
				StartNode: true,
			},
			{
				IPAddr:    "10.0.0.1",
				StartNode: true,
			},
		},
	}
	gameMap.SelectStartNode()
	require.Len(t, gameMap.startNodes, 1)
}
