package action

import (
	"fmt"
	"testing"

	"github.com/JonLMyers/GoFxell/game"
	"github.com/stretchr/testify/require"
)

func TestShowRoutes(t *testing.T) {
	gameMap := game.NewMap("testdata/devmap.json")

	testTeam := game.NewTeam("Red", "A", gameMap)
	startNode := testTeam.StartNode
	//fmt.Println(startNode.IPAddr)
	//fmt.Println(testTeam.DiscoveredNodes)
	ok, _ := Scan(startNode.IPAddr, testTeam)
	if !ok {
		fmt.Println("Scan: ", ok)
		return
	}
	ok, _ = PlatformExploit(startNode.IPAddr, testTeam)
	if !ok {
		fmt.Println("Exploit: ", ok)
		return
	}

	routes, error := ShowRoutes(startNode.IPAddr, testTeam, *gameMap)

	if error != nil {
		fmt.Println(error)
	}
	require.Len(t, routes, 1)
	require.Equal(t, "1.1.1.2", routes[0])
}
