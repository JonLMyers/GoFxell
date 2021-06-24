package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScan(t *testing.T) {
	gameMap := NewMap("testdata/devmap.json")
	testTeam := NewTeam("Red", "A", gameMap)
	startNode := testTeam.StartNode
	node, _ := testTeam.Scan(startNode.IPAddr)
	require.Equal(t, "Windows", node.Platform)
}

func TestShowRoutes(t *testing.T) {
	gameMap := NewMap("testdata/devmap.json")
	testTeam := NewTeam("Red", "A", gameMap)
	startNode := testTeam.StartNode
	testTeam.Scan(startNode.IPAddr)
	testTeam.PlatformExploit(startNode.IPAddr)
	routes, error := testTeam.ShowRoutes(startNode.IPAddr, *gameMap)

	if error != nil {
		fmt.Println(error)
	}
	require.Len(t, routes, 1)
	require.Equal(t, "1.1.1.2", routes[0])
}

func TestShowTargets(t *testing.T) {
	gameMap := NewMap("testdata/devmap.json")
	testTeam := NewTeam("Red", "A", gameMap)
	startNode := testTeam.StartNode
	testTeam.Scan(startNode.IPAddr)
	require.Len(t, testTeam.DiscoveredNodes, 1)

	testTeam.ShowRoutes(startNode.IPAddr, *gameMap)
	testTeam.Scan(testTeam.DiscoveredNodes[1].IPAddr)
	require.Len(t, testTeam.DiscoveredNodes, 2)

}
