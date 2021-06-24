package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeployExfiltrateObjective(t *testing.T) {
	gameMap := NewMap("testdata/devmap.json")
	testTeam := NewTeam("Red", "A", gameMap)
	startNode := testTeam.StartNode
	testTeam.Scan(startNode.IPAddr)
	testTeam.ServiceExploit(startNode.IPAddr, "http")
	proc, _ := testTeam.NewProcess(startNode.IPAddr, "GetFiles.exe")
	timeleft, _ := testTeam.ExfiltrateObjective(startNode.IPAddr, &proc)
	require.Equal(t, 110, timeleft)
}
