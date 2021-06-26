package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLog(t *testing.T) {
	gameMap := NewMap("testdata/devmap.json")
	testTeam := NewTeam("Red", "A", gameMap)
	startNode := testTeam.StartNode
	testTeam.Scan(startNode.IPAddr)
	log, _ := NewLog(startNode.IPAddr, "Test", "TestTest", 10, testTeam)
	require.Equal(t, "Test", log.Type)
}
func TestGetLogs(t *testing.T) {
	gameMap := NewMap("testdata/devmap.json")
	testTeam := NewTeam("Red", "A", gameMap)
	startNode := testTeam.StartNode
	testTeam.Scan(startNode.IPAddr)
	NewLog(startNode.IPAddr, "Test", "TestTest", 10, testTeam)
	logs, _ := GetLogs(startNode.IPAddr, testTeam)
	require.Equal(t, "Test", logs[0].Type)
}

func TestDeleteLog(t *testing.T) {
	gameMap := NewMap("testdata/devmap.json")
	testTeam := NewTeam("Red", "A", gameMap)
	startNode := testTeam.StartNode
	testTeam.Scan(startNode.IPAddr)
	testTeam.ServiceExploit(startNode.IPAddr, "http")
	log, _ := NewLog(startNode.IPAddr, "Test", "TestTest", 10, testTeam)
	success, _ := DeleteLog(log.Id, startNode.IPAddr, testTeam)
	require.Equal(t, true, success)
	logs, _ := GetLogs(startNode.IPAddr, testTeam)
	require.Equal(t, true, logs[1].Deleted)
}
