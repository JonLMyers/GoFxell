import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/JonLMyers/GoFxell/game"
)

case strings.HasPrefix(cmd, "scan"):
case strings.HasPrefix(cmd, "exploit"):
case strings.HasPrefix(cmd, "dos"):
case strings.HasPrefix(cmd, "show targets"):
case strings.HasPrefix(cmd, "show resources"):
case strings.HasPrefix(cmd, "connect"):
func ProcessScan(){

}
func ExploitProcesor(cmdParts []string){
	
}
func DosProcesor(){
	
}
func ShowTargetsProcesor(){
	
}
func ShowResourcesProcesor(){
	
}
func ConnectProcesor(){
	
}
func Executor(cmd string) {
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		return
	}
	cmdParts := strings.Split(cmd, " ")

	if strings.HasPrefix(cmd, "scan") {
		playerTeam.Scan(cmdParts[1])
		fmt.Println(playerTeam.ShowTargets())
		return
	}

	if strings.HasPrefix(cmd, "exploit") {
		if len(cmdParts) < 3 {
			fmt.Println("Invalid Exploit Syntax")
		}
		//Create a prompt here and for connect... oh wait doesn't work.
		if cmdParts[1] == "Windows" || cmdParts[1] == "Linux" {
			ok, _ := playerTeam.PlatformExploit(cmdParts[2])
			if !ok {
				fmt.Println("Platform Exploit Failed")
				return
			}
			fmt.Println("Platform Exploit Successful")
			playerTeam.ShowTargets()
			return
		}
		if cmdParts[1] == "Web" || cmdParts[1] == "SSH" || cmdParts[1] == "SMTP" || cmdParts[1] == "FTP" || cmdParts[1] == "Mail" || cmdParts[1] == "SMB" {
			ok, _ := playerTeam.ServiceExploit(cmdParts[2], cmdParts[1])
			if !ok {
				fmt.Println("Service Exploit Failed")
				return
			}
			fmt.Println("Service Exploit Successful")
			playerTeam.ShowTargets()
			return
		}
		return
	}

	if strings.HasPrefix(cmd, "dos") {
		playerTeam.DenialOfService(cmdParts[1])
		return
	}

	if strings.HasPrefix(cmd, "show targets") {
		nodes := playerTeam.ShowTargets()
		fmt.Println(nodes)
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
		playerTeam.ShowRoutes(n.IPAddr, *gameMap)
		fmt.Println(playerTeam.ShowTargets())
		return
	}
	if strings.HasPrefix(cmd, "deploy miner") {
		if len(cmdParts) < 3 {
			fmt.Println("Invalid Syntax. Ensure you are providing a miner type {Bandwidth, IO, Entropy, or CPU")
		}
		playerTeam.DeployMiner(n.IPAddr, cmdParts[2])
		return
	}
	if strings.HasPrefix(cmd, "deploy firewall") {
		ok, error := game.DeployFirewall(n.IPAddr, playerTeam)
		if !ok {
			fmt.Println("Firewall Deployment Failed: ", error)
		}
		return
	}

	if strings.HasPrefix(cmd, "deploy netmon") {
		proc, _ := playerTeam.NewProcess(n.IPAddr, "netmon.exe")
		monitor := playerTeam.NewMonitor("Network", &proc)
		if monitor.Process.CMD != "netmon.exe" {
			fmt.Println("netmon.exe Deployment Failed.")
			return
		}
		return
	}

	if strings.HasPrefix(cmd, "check monitor") {
		if len(cmdParts) < 3 {
			fmt.Println("Invalid Syntax. Ensure you are providing the monitor's PID")
			return
		}
		i, err := strconv.Atoi(cmdParts[2])
		if err != nil {
			fmt.Println("Invalid PID")
			return
		}
		monLogs, error := game.CheckMonitor(n.IPAddr, playerTeam, i)
		if monLogs[0].Type == "" {
			fmt.Println("No Logs Available: ", error)
		}
		return
	}
	if strings.HasPrefix(cmd, "show proc") {
		processes, _ := playerTeam.ShowProcesses(n.IPAddr, *gameMap)
		fmt.Println(processes)
		return
	}

	if strings.HasPrefix(cmd, "show footprint") {
		resources, _ := playerTeam.ViewNodeFootprint(n.IPAddr)
		fmt.Println(resources)
		return
	}

	if strings.HasPrefix(cmd, "show logs") {
		logs, _ := game.GetLogs(n.IPAddr, playerTeam)
		fmt.Println(logs)
		return
	}

	if strings.HasPrefix(cmd, "clean log") {
		if len(cmdParts) < 3 {
			fmt.Println("Invalid Syntax. Ensure you are providing a LogId")
		}

		logs, _ := game.DeleteLog(cmdParts[2], n.IPAddr, playerTeam)
		fmt.Println(logs)
		return
	}

	if strings.HasPrefix(cmd, "exfiltrate") {
		if len(cmdParts) < 2 {
			fmt.Println("Invalid Syntax")
			return
		}
		proc, _ := playerTeam.NewProcess(n.IPAddr, "GetFiles.exe")
		ticker := time.NewTicker(time.Second * 10)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				timeleft, _ := playerTeam.ExfiltrateObjective(n.IPAddr, &proc)
				fmt.Println("Time Remaining:", timeleft)
			}
		}
	}

	if strings.HasPrefix(cmd, "show resources") {
		resources := playerTeam.GetResources()
		fmt.Println(resources)
		return
	}

	if strings.HasPrefix(cmd, "kill") {
		ok, err := playerTeam.KillProcess(n.IPAddr, cmdParts[1])
		if !ok {
			fmt.Println(err)
			fmt.Println("Failed to Kill Process")
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
