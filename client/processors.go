package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/JonLMyers/GoFxell/game"
)

func ScanProcessor(cmdParts []string) []byte {
	playerTeam.Scan(cmdParts[1])
	return ShowTargetsProcessor()
}

func ExploitProcessor(cmdParts []string) []byte {
	ok, message := RequiredArguments(cmdParts, 2)
	if !ok {
		return message
	}
	ok, message = DiscoveredNode(cmdParts[1], playerTeam)
	if !ok {
		return message
	}

	if len(cmdParts) == 2 {
		ok, _ := playerTeam.PlatformExploit(cmdParts[1])
		if !ok {
			return Message("Platform Exploit Failed")
		}
		return ShowTargetsProcessor()
	}

	if cmdParts[2] == "Web" || cmdParts[2] == "SSH" || cmdParts[2] == "SMTP" || cmdParts[2] == "FTP" || cmdParts[2] == "Mail" || cmdParts[2] == "SMB" {
		ok, _ := playerTeam.ServiceExploit(cmdParts[1], cmdParts[2])
		if !ok {
			return Message("Service Exploit Failed")
		}
		return ShowTargetsProcessor()
	}
	return Message("Exploit Function Failure")
}

func DosProcessor(cmdParts []string) []byte {
	ok, message := RequiredArguments(cmdParts, 2)
	if !ok {
		return message
	}
	ok, message = DiscoveredNode(cmdParts[1], playerTeam)
	if !ok {
		return message
	}
	ok, err := playerTeam.DenialOfService(cmdParts[1])
	if !ok {
		return Message(err.Error())
	}
	return Message("Denail of Service Initiated")
}

func ShowTargetsProcessor() []byte {
	nodes := playerTeam.ShowTargets()
	var targets []TargetResponse
	for _, node := range nodes {
		j := TargetResponse{IPAddr: node.IPAddr, Platform: node.Platform, Services: node.Services, Routes: node.Routes, NodeOwned: node.NodeOwned}
		targets = append(targets, j)
	}
	bytes, err := json.Marshal(targets)
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}

func ShowResourcesProcessor() []byte {
	resources := playerTeam.GetResources()
	var response []string
	for _, resource := range resources {
		response = append(response, resource)
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}

func ConnectProcessor(cmdParts []string) []byte {
	ok, message := RequiredArguments(cmdParts, 2)
	if !ok {
		return message
	}
	ok, message = DiscoveredNode(cmdParts[1], playerTeam)
	if !ok {
		return message
	}

	ok, err := Connect(cmdParts[1], playerTeam, *gameMap)
	if !ok || err != nil {
		return Message("Connection Refused")
	}
	return Message("Connection Closed")
}

func ShowRoutesProcessor() []byte {
	playerTeam.ShowRoutes(n.IPAddr, *gameMap)
	return ShowTargetsProcessor()
}

func DeployMinerProcessor(cmdParts []string) []byte {
	ok, message := RequiredArguments(cmdParts, 3)
	if !ok {
		return message
	}
	playerTeam.DeployMiner(n.IPAddr, cmdParts[2])
	return Messages([]string{"miner Deployed on Target", "miner.Mine() Initiated"})
}

func DeployFirewallProcessor(cmdParts []string) []byte {
	ok, _ := game.DeployFirewall(n.IPAddr, playerTeam)
	if !ok {

		return Message("Firewall Deployment Failed")
	}
	return Message("Firewall Deployment Successful")
}

func DeployNetmonProcessor(cmdParts []string) []byte {
	proc, _ := playerTeam.NewProcess(n.IPAddr, "netmon.exe")
	monitor := playerTeam.NewMonitor("Network", &proc)
	if monitor.Process.CMD != "netmon.exe" {
		return Message("Network Monitor Deployment Failed")
	}
	return Messages([]string{"netmon Deployed on Target", "netmon.Monitor() Initiated"})
}

func CheckMonitorProcessor(cmdParts []string) []byte {
	ok, message := RequiredArguments(cmdParts, 2)
	if !ok {
		return message
	}

	monLogs, _ := game.CheckMonitor(n.IPAddr, playerTeam)
	if len(monLogs) == 0 {
		return Message("No Logs Available")
	}
	var response []LogResponse
	for _, log := range monLogs {
		response = append(response, LogResponse{Id: log.Id, Timestamp: log.Timestamp.Format("Mon Jan _2 15:04:05"), Type: log.Type, LogMessage: log.LogString})
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}

func ShowProcessesProcessor(cmdParts []string) []byte {
	processes, _ := playerTeam.ShowProcesses(n.IPAddr, *gameMap)
	var response []ProcessResponse
	for _, proc := range processes {
		response = append(response, ProcessResponse{PID: proc.PID, CMD: proc.CMD, Team: proc.TeamName})
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}

func ShowFootprintProcessor(cmdParts []string) []byte {
	footprint, _ := playerTeam.ViewNodeFootprint(n.IPAddr)
	message := "Node Footprint: " + strconv.Itoa(footprint)
	return Message(message)
}

func ShowLogsProcessor(cmdParts []string) []byte {
	logs, _ := game.GetLogs(n.IPAddr, playerTeam)
	var response []LogResponse
	for _, log := range logs {
		response = append(response, LogResponse{Id: log.Id, Timestamp: log.Timestamp.Format("Mon Jan _2 15:04:05"), Type: log.Type, LogMessage: log.LogString})
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}

func CleanLogProcessor(cmdParts []string) []byte {
	ok, message := RequiredArguments(cmdParts, 3)
	if !ok {
		return message
	}

	ok, _ = game.DeleteLog(cmdParts[2], n.IPAddr, playerTeam)
	if !ok {
		return Message("Failed to Delete Logs")
	}
	return Message("Log Cleaned Successfully")
}

func ExfiltrateProcessor(cmdParts []string) []byte {
	proc, _ := playerTeam.NewProcess(n.IPAddr, "GetFiles")
	//ticker := time.NewTicker(time.Second)
	//defer ticker.Stop()

	//for {
	//select {
	//case <-ticker.C:
	timeleft, _ := playerTeam.ExfiltrateObjective(n.IPAddr, &proc)
	return Message("Seconds Remaining: " + strconv.Itoa(timeleft))
	//}
	//}
}

func KillProcessor(cmdParts []string) []byte {
	ok, message := RequiredArguments(cmdParts, 2)
	if !ok {
		return message
	}
	ok, err := playerTeam.KillProcess(n.IPAddr, cmdParts[1])
	if !ok {
		fmt.Println(err)
		return Message("Failed to Kill Process")
	}
	return Message("Process Destroyed")
}

// Helper Function for Processors. Sends a single json message
func Message(message string) []byte {
	bytes, err := json.Marshal(SingleResponse{Message: message})
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}

// Helper Function for Processors Sends multiple (single) json messages
func Messages(messages []string) []byte {
	var responses []SingleResponse
	for _, message := range messages {
		j := SingleResponse{Message: message}
		responses = append(responses, j)
	}
	bytes, err := json.Marshal(responses)
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}
