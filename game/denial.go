package game

import (
	"fmt"
	"time"
)

const (
	StandardDosTime      = 30
	StandardDosFootprint = 50
	StandardDosPower     = 5
)

//Check this with kyle: I think target is dumb.
func (team *Team) DenialOfService(ipAddr string) (bool, error) {
	index, _ := team.DiscoveredNodes.IndexOf(ipAddr)
	if team.DiscoveredNodes[index].DiscoveredPlatform {
		target := "node"
		if FirewallExists(index, team) {
			target = "firewall"
		}
		NewLog(ipAddr, "Network", "Heavy Incoming Network Traffic", StandardDosFootprint, team)
		if target == "node" {
			//fmt.Println("Target Node Disabeld for 30 Seconds. All network processes are now broken.")
			go Deny(index, team, StandardDosTime)
		} else {
			team.BreakFirewall(index, StandardDosPower)
		}
		return true, nil
	}
	return false, fmt.Errorf("DenialOfService Function Failed")
}

func Deny(index int, team *Team, locktime int) {
	team.DiscoveredNodes[index].Node.Active = false
	time.Sleep(time.Duration(locktime) * time.Second)
	team.DiscoveredNodes[index].Node.Active = true
}

func (team *Team) BreakFirewall(index int, power int) {
	for {
		if team.DiscoveredNodes[index].FirewallStrength > 0 {
			team.DiscoveredNodes[index].FirewallStrength = team.DiscoveredNodes[index].FirewallStrength - StandardDosPower
			fmt.Println("Packets delivered, Firewall Strength: ", team.DiscoveredNodes[index].FirewallStrength)
			time.Sleep(1 * time.Second)
		} else {
			for _, proc := range team.DiscoveredNodes[index].Processes {
				if proc.CMD == "firewall.exe" {
					team.KillProcess(team.DiscoveredNodes[index].IPAddr, proc.PID)
					//fmt.Println("Firewall Successfully Disabled")
					return
				}
			}
		}
	}
}
