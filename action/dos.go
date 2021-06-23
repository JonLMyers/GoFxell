package action

import (
	"fmt"
	"time"

	"github.com/JonLMyers/GoFxell/game"
)

const (
	StandardDosTime      = 30
	StandardDosFootprint = 50
	StandardDosPower     = 5
)

func DenialOfService(ipAddr string, team *game.Team) (bool, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println("IP Address not in teams discovered nodes")
		return false, err
	}
	if team.DiscoveredNodes[index].DiscoveredPlatform {
		target := "node"
		if FirewallExists(index, team) {
			fmt.Println("Firewall present on target. Preparing Firewall DoS.")
			target = "firewall"
		}
		game.CreateLog(ipAddr, "Network", "Heavy Incoming Network Traffic", StandardDosFootprint, team)
		if target == "node" {
			fmt.Println("Target Node Disabeld for 30 Seconds. All network processes are now broken.")
			go Deny(index, team, StandardDosTime)
		} else {
			BreakFirewall(index, team, StandardDosPower)
		}
		return true, nil
	}

	return false, nil

}

func Deny(index int, team *game.Team, locktime int) {
	team.DiscoveredNodes[index].Node.Active = false
	time.Sleep(time.Duration(locktime) * time.Second)
	team.DiscoveredNodes[index].Node.Active = true
}

func BreakFirewall(index int, team *game.Team, power int) {
	for {
		if team.DiscoveredNodes[index].FirewallStrength > 0 {
			team.DiscoveredNodes[index].FirewallStrength = team.DiscoveredNodes[index].FirewallStrength - StandardDosPower
			fmt.Println("Packets delivered, Firewall Strength: ", team.DiscoveredNodes[index].FirewallStrength)
			time.Sleep(1 * time.Second)
		} else {
			for _, proc := range team.DiscoveredNodes[index].Processes {
				if proc.CMD == "firewall.exe" {
					KillProcess(team.DiscoveredNodes[index].IPAddr, proc.PID, team)
					fmt.Println("Firewall Successfully Disabled")
					return
				}
			}
		}
	}
}
