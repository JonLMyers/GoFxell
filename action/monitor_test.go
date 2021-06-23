package action

import (
	"fmt"
	"testing"

	"github.com/JonLMyers/GoFxell/game"
	"github.com/stretchr/testify/require"
)

func TestNetworkMonitor(t *testing.T) {
	gameMap := game.NewMap("testdata/devmap.json")

	testTeam := game.NewTeam("Red", "A", gameMap)
	startNode := testTeam.StartNode
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
	monitor, _ := DeployNetworkMonitor(startNode.IPAddr, testTeam, *gameMap)
	require.Equal(t, "Red", monitor.TeamName)

	game.CreateLog(startNode.IPAddr, "Network", "Testdata", 50, testTeam)
	require.Equal(t, "Testdata", testTeam.DiscoveredNodes[0].Logs[1].LogString)
	fmt.Println(testTeam.DiscoveredNodes[0].Logs)

	monLogs, _ := CheckMonitor(startNode.IPAddr, testTeam, monitor.PID)
	require.Equal(t, monLogs[0].Message, "Testdata")
}
