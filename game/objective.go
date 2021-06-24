package game

type Objective struct {
	ObjectiveType string `json:"objective_type"`
	CompleteTime  int    `json:"complete_time"`
}

/*
//Validation must be included to ensure that players do not spawn multiple exfil procs
func (team *Team) NewExfiltrateObjective(ipAddr string, Process) (Process, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return Process{}, err
	}
	if team.DiscoveredNodes[index].Node.ObjectiveNode {

		process, _ := team.NewProcess(ipAddr, "GetFiles.exe")
		//fmt.Println("Process: GetFiles.exe@", PID, " Deployed")
		//fmt.Println("Process: GetFiles.Exfiltrate() Initiated")
		team.UpdateNodeFootprint(team.DiscoveredNodes[index].Node.IPAddr, 20)
		return process, nil
	}
	return Process{}, nil
}
*/
func (team *Team) ExfiltrateObjective(ipAddr string, process *Process) (int, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		return -1, err
	}
	if team.DiscoveredNodes[index].Node.ObjectiveNode && containsPid(team.DiscoveredNodes[index].Node.Processes, process.PID) {
		if team.DiscoveredNodes[index].Node.Objective.CompleteTime > 0 {
			team.DiscoveredNodes[index].Node.Objective.CompleteTime = team.DiscoveredNodes[index].Node.Objective.CompleteTime - ExfiltrateObjectiveTimeReduction
			//fmt.Printf("%d Seconds Remaining\n", team.DiscoveredNodes[index].Node.Objective.CompleteTime)
			return team.DiscoveredNodes[index].Node.Objective.CompleteTime, nil
		} else {
			team.ObjectiveComplete = true
			return team.DiscoveredNodes[index].Node.Objective.CompleteTime, nil
		}
	}
	return -1, nil
}
