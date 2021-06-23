package action

import (
	"fmt"

	"github.com/JonLMyers/GoFxell/game"
)

const (
	NetworkMonitorFootprintValue = 5
)

func CheckMonitor(ipAddr string, team *game.Team, PID int) ([]game.MonitorLog, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, monitor := range team.DiscoveredNodes[index].Node.Monitors {
		if monitor.TeamName == team.Name {
			fmt.Println(monitor.CollectedLogs)
			//monitor.DumpMonitor()
			return monitor.CollectedLogs, nil
		}
	}
	return nil, nil
}

func DeployMonitor(monitorType string, ipAddr string, team *game.Team, gameMap game.Map) bool {
	// TODO: Use a switch statement (*forehead*)
	if monitorType == "Network" {
		monitor, _ := DeployNetworkMonitor(ipAddr, team, gameMap)
		if monitor.PID == 0 {
			return false
		}
	}
	if monitorType == "Process" {

	}
	if monitorType == "Filesystem" {

	}
	return false

}

func DeployNetworkMonitor(ipAddr string, team *game.Team, gameMap game.Map) (game.Monitor, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return game.Monitor{}, err
	}
	for _, proc := range team.DiscoveredNodes[index].Node.Processes {
		if proc.CMD == "netmon.exe" {
			if proc.TeamName == team.Name {
				fmt.Println("Network Monitor Already deployed")
				return game.Monitor{}, nil
			}
		}
	}
	PID, ok := CreatePid(ipAddr, team, NetworkMonitorFootprintValue)
	monitor := game.CreateMonitor("Network", PID, team)
	//monitor := game.Monitor{"Network", &team.DiscoveredNodes[index].Node, PID, []game.MonitorLog{}, team.Name}
	fmt.Printf("%p\n", &monitor)
	go monitor.MonitorNetwork(&team.DiscoveredNodes[index].Node, team)
	game.CreateProcess(ipAddr, team, PID, "netmon.exe")
	if ok != nil {
		return game.Monitor{}, ok
	}
	fmt.Println("Process: netmon.exe@", PID, "Deployed")
	fmt.Println("Process: netmon.Monitor() Initiated")
	team.DiscoveredNodes[index].Node.Monitors = append(team.DiscoveredNodes[index].Node.Monitors, monitor)
	team.Monitors = append(team.Monitors, monitor)
	return monitor, nil
}
