package action

import (
	"fmt"

	"github.com/JonLMyers/GoFxell/game"
)

func ExfiltrateObjective(ipAddr string, PID int, team *game.Team) (bool, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return false, err
	}
	if team.DiscoveredNodes[index].Node.ObjectiveNode && contains(team.DiscoveredNodes[index].Node.PIDS, PID) {
		if team.DiscoveredNodes[index].Node.Objective.CompleteTime > 0 {
			team.DiscoveredNodes[index].Node.Objective.CompleteTime = team.DiscoveredNodes[index].Node.Objective.CompleteTime - ExfiltrateObjectiveTimeReduction
			fmt.Printf("%d Seconds Remaining\n", team.DiscoveredNodes[index].Node.Objective.CompleteTime)
			return true, nil
		} else {
			team.ObjectiveComplete = true
			return true, nil
		}
	}
	return false, nil
}
