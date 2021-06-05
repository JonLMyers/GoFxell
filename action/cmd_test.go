package action

import (
	"testing"

	"github.com/JonLMyers/GoFxell/game"
	"github.com/stretchr/testify/require"
)

func TestShowRoutes(t *testing.T) {
	gameMap := game.NewMap("testdata/testMaps.json")

	testTeam := game.NewTeam("test", gameMap)
	startNode := testTeam.StartNode
	routes, _ := ShowRoutes(startNode.IPAddr, testTeam, *gameMap)
	require.Len(t, routes, 3)
	require.Equal(t, "10.0.0.5", routes[1])
}
