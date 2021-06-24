package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewProcess(t *testing.T) {
	gameMap := NewMap("testdata/devmap.json")
	testTeam := NewTeam("Red", "A", gameMap)
	startNode := testTeam.StartNode
	testTeam.Scan(startNode.IPAddr)
	testTeam.ServiceExploit(startNode.IPAddr, "http")
	proc, _ := testTeam.NewProcess(startNode.IPAddr, "GetFiles.exe")
	require.Equal(t, "GetFiles.exe", proc.CMD)
}

func ShowProcesses(t *testing.T) {
	gameMap := NewMap("testdata/devmap.json")
	testTeam := NewTeam("Red", "A", gameMap)
	startNode := testTeam.StartNode
	testTeam.Scan(startNode.IPAddr)
	testTeam.ServiceExploit(startNode.IPAddr, "http")
	testTeam.NewProcess(startNode.IPAddr, "GetFiles.exe")
	procs, _ := testTeam.ShowProcesses(startNode.IPAddr, *gameMap)
	require.Equal(t, "GetFiles.exe", procs[0].CMD)
}

func KillProcess(t *testing.T) {
	gameMap := NewMap("testdata/devmap.json")
	testTeam := NewTeam("Red", "A", gameMap)
	startNode := testTeam.StartNode
	testTeam.Scan(startNode.IPAddr)
	testTeam.ServiceExploit(startNode.IPAddr, "http")
	proc, _ := testTeam.NewProcess(startNode.IPAddr, "GetFiles.exe")
	procs, _ := testTeam.ShowProcesses(startNode.IPAddr, *gameMap)
	require.Equal(t, "GetFiles.exe", procs[0].CMD)
	testTeam.KillProcess(startNode.IPAddr, proc.PID)
	procs, _ = testTeam.ShowProcesses(startNode.IPAddr, *gameMap)
	require.Len(t, procs, 0)
}
