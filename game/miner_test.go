package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeployMiner(t *testing.T) {
	gameMap := NewMap("testdata/devmap.json")
	testTeam := NewTeam("Red", "A", gameMap)
	startNode := testTeam.StartNode
	testTeam.Scan(startNode.IPAddr)
	testTeam.ServiceExploit(startNode.IPAddr, "http")
	miner, _ := testTeam.DeployMiner(startNode.IPAddr, "CPU")
	require.Equal(t, "CPU", miner.minerType)
}
