package client

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Executor(cmd string) {
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		return
	}
	cmdParts := strings.Split(cmd, " ")

	switch {
	case strings.HasPrefix(cmd, "scan"):
		fmt.Println(string(ScanProcessor(cmdParts)))

	case strings.HasPrefix(cmd, "exploit"):
		fmt.Println(string(ExploitProcessor(cmdParts)))

	case strings.HasPrefix(cmd, "dos"):
		fmt.Println(string(DosProcessor(cmdParts)))

	case strings.HasPrefix(cmd, "show targets"):
		fmt.Println(string(ShowTargetsProcessor()))

	case strings.HasPrefix(cmd, "show resources"):
		fmt.Println(string(ShowResourcesProcessor()))

	case strings.HasPrefix(cmd, "connect"):
		fmt.Println(string(ConnectProcessor(cmdParts)))

	case strings.HasPrefix(cmd, "exit"):
		os.Exit(0)
		return

	default:
		fmt.Println("Invalid Command")
	}
}

func CmdExecutor(cmd string) {
	cmd = strings.TrimSpace(cmd)
	cmdParts := strings.Split(cmd, " ")
	switch {
	case strings.HasPrefix(cmd, "show routes"):
		fmt.Println(string(ShowRoutesProcessor()))

	case strings.HasPrefix(cmd, "deploy miner"):
		fmt.Println(string(DeployMinerProcessor(cmdParts)))

	case strings.HasPrefix(cmd, "deploy firewall"):
		fmt.Println(string(DeployFirewallProcessor(cmdParts)))

	case strings.HasPrefix(cmd, "deploy netmon"):
		fmt.Println(string(DeployNetmonProcessor(cmdParts)))

	case strings.HasPrefix(cmd, "check monitor"):
		fmt.Println(string(CheckMonitorProcessor(cmdParts)))

	case strings.HasPrefix(cmd, "show proc"):
		fmt.Println(string(ShowProcessesProcessor(cmdParts)))

	case strings.HasPrefix(cmd, "show footprint"):
		fmt.Println(string(ShowFootprintProcessor(cmdParts)))

	case strings.HasPrefix(cmd, "show logs"):
		fmt.Println(string(ShowLogsProcessor(cmdParts)))

	case strings.HasPrefix(cmd, "clean log"):
		fmt.Println(string(CleanLogProcessor(cmdParts)))

	case strings.HasPrefix(cmd, "exfiltrate"):
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println(string(ExfiltrateProcessor(cmdParts)))
			}
		}

	case strings.HasPrefix(cmd, "show resources"):
		fmt.Println(string(ShowResourcesProcessor()))

	case strings.HasPrefix(cmd, "kill"):
		fmt.Println(string(KillProcessor(cmdParts)))

	case strings.HasPrefix(cmd, "exit"):
		return

	case strings.HasPrefix(cmd, ""):
		//Do nothing :)
	default:
		fmt.Println("Invalid Command")
	}
}
