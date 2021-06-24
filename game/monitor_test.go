package game

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNetworkMonitor(t *testing.T) {
	gameMap := NewMap("testdata/devmap.json")

	testTeam := NewTeam("Red", "A", gameMap)
	startNode := testTeam.StartNode
	testTeam.Scan(startNode.IPAddr)
	testTeam.PlatformExploit(startNode.IPAddr)
	process, _ := testTeam.NewProcess(startNode.IPAddr, "netmon.exe")
	monitor := testTeam.NewMonitor("Network", &process)
	require.Equal(t, "Red", monitor.TeamName)

	NewLog(startNode.IPAddr, "Network", "Testdata", 50, testTeam)
	require.Equal(t, "Testdata", testTeam.DiscoveredNodes[0].Logs[2].LogString)
	go monitor.MonitorNetwork(testTeam)
	time.Sleep(2 * time.Second)
	require.Equal(t, monitor.CollectedLogs[0].LogString, "Unknown Network Connection")
}
