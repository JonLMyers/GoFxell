package game

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	defaultMineAmount = 10
	defaultInterval   = 10
)

type Miner struct {
	minerId   uuid.UUID
	Process   *Process
	TeamName  string
	minerType string
	interval  int
	amount    int
}

// Refactor this due to the new process system :)
func (team *Team) DeployMiner(ipAddr string, minerType string) (Miner, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return Miner{}, err
	}
	if minerType == "Bandwidth" || minerType == "IO" || minerType == "CPU" || minerType == "Entropy" {
		if team.DiscoveredNodes[index].MinersDeployed <= team.DiscoveredNodes[index].MaxMiners {
			process, _ := team.NewProcess(ipAddr, "miner.exe")
			miner := team.NewMiner(minerType, &process)
			team.DiscoveredNodes[index].Node.Miners = append(team.DiscoveredNodes[index].Node.Miners, miner)
			team.DiscoveredNodes[index].MinersDeployed = team.DiscoveredNodes[index].MinersDeployed + 1
			//fmt.Println("Process: miner.exe@", PID, " Deployed")
			go miner.Mine(team, index)

			//fmt.Println("Process: miner.Mine() Initiated")
			return miner, nil
		}
		fmt.Println("Not enough space on Target")
	}
	return Miner{}, nil

}

func (team *Team) NewMiner(minerType string, process *Process) Miner {
	miner := Miner{uuid.New(), process, team.Name, minerType, defaultInterval, defaultMineAmount}
	return miner
}

// This runs forever even if the miner dies. It just doesn't add value.
func (miner *Miner) Mine(team *Team, index int) {
	ticker := time.NewTicker(time.Second * time.Duration(miner.interval))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if !team.DiscoveredNodes[index].Active || !miner.Process.Alive {
				return
			}

			team.mutex.Lock()
			defer team.mutex.Unlock()
			switch miner.minerType {
			case "Bandwidth":
				team.Bandwidth = team.Bandwidth + miner.amount
				continue
			case "IO":
				team.Io = team.Io + miner.amount
				continue
			case "Entropy":
				team.Entropy = team.Entropy + miner.amount
				continue
			case "CPU":
				team.Cpu = team.Cpu + miner.amount
				continue
			}
			team.mutex.Unlock()
		}
	}
}
