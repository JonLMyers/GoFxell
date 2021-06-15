package client

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/JonLMyers/GoFxell/action"
	"github.com/JonLMyers/GoFxell/game"
)

func Executor(cmd string) {
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		return
	}
	cmdParts := strings.Split(cmd, " ")

	if strings.HasPrefix(cmd, "scan") {
		action.Scan(cmdParts[1], playerTeam)
		action.ShowTargets(*playerTeam)
		return
	}
	if strings.HasPrefix(cmd, "exploit") {
		if len(cmdParts) < 3 {
			fmt.Println("Invalid Exploit Syntax")
		}
		//Create a prompt here and for connect... oh wait doesn't work.
		if cmdParts[1] == "Windows" || cmdParts[1] == "Linux" {
			ok, _ := action.PlatformExploit(cmdParts[2], playerTeam)
			if !ok {
				fmt.Println("Platform Exploit Failed")
				return
			}
			fmt.Println("Platform Exploit Successful")
			action.ShowTargets(*playerTeam)
			return
		}
		if cmdParts[1] == "Web" || cmdParts[1] == "SSH" || cmdParts[1] == "SMTP" || cmdParts[1] == "FTP" || cmdParts[1] == "Mail" || cmdParts[1] == "SMB" {
			ok, _ := action.ServiceExploit(cmdParts[2], playerTeam, cmdParts[1])
			if !ok {
				fmt.Println("Service Exploit Failed")
				return
			}
			fmt.Println("Servicem Exploit Successful")
			action.ShowTargets(*playerTeam)
			return
		}
		return
	}

	if strings.HasPrefix(cmd, "show targets") {
		action.ShowTargets(*playerTeam)
		return
	}

	if strings.HasPrefix(cmd, "show resources") {
		resources := playerTeam.GetResources()
		fmt.Println(resources)
		return
	}

	if strings.HasPrefix(cmd, "connect") {
		if len(cmdParts) < 2 {
			fmt.Println("Invalid Connect Syntax")
			return
		}
		ok, err := Connect(cmdParts[1], playerTeam, *gameMap)
		if !ok || err != nil {
			fmt.Println("Connection Refused")
			return
		}
		fmt.Println("Connection Closed")
		return
	}

	if strings.HasPrefix(cmd, "exit") {
		os.Exit(0)
		return
	}

	fmt.Println("Invalid Command")
}

func CmdExecutor(cmd string) {
	cmd = strings.TrimSpace(cmd)
	cmdParts := strings.Split(cmd, " ")
	if cmd == "show routes" {
		action.ShowRoutes(n.IPAddr, playerTeam, *gameMap)
		return
	}
	if strings.HasPrefix(cmd, "deploy miner") {
		if len(cmdParts) < 3 {
			fmt.Println("Invalid Syntax. Ensure you are providing a miner type {Bandwidth, IO, Entropy, or CPU")
		}
		action.DeployMiner(n.IPAddr, cmdParts[2], playerTeam)
		return
	}
	if strings.HasPrefix(cmd, "show proc") {
		processes, _ := action.ShowProcesses(n.IPAddr, playerTeam, *gameMap)
		fmt.Println(processes)
		return
	}

	if strings.HasPrefix(cmd, "show footprint") {
		resources, _ := action.ViewNodeFootprint(n.IPAddr, playerTeam)
		fmt.Println(resources)
		return
	}

	if strings.HasPrefix(cmd, "show logs") {
		logs, _ := game.ViewLogs(n.IPAddr, playerTeam)
		fmt.Println(logs)
		return
	}

	if strings.HasPrefix(cmd, "kill") {
		i, err := strconv.Atoi(cmdParts[1])
		if err != nil {
			fmt.Println("Invalid PID")
			return
		}
		ok, err := action.KillProcess(n.IPAddr, i, playerTeam)
		if !ok {
			fmt.Println(err)
			fmt.Println("Failed to kill process")
			return
		}
		fmt.Println("Process Destroyed")
		return
	}
	/*if cmd == "exit" {
		return
	}*/
	fmt.Println("Invalid Command")
}
