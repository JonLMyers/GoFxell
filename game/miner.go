package game

import (
	"time"

	"github.com/google/uuid"
)

const (
	defaultMineAmount = 10
	defaultInterval   = 10
)

type Miner struct {
	minerId   uuid.UUID
	PID       int
	CMD       string
	TeamName  string
	minerType string
	interval  int
	amount    int
}

func NewMiner(minerType string, team Team, PID int, cmd string) Miner {
	miner := Miner{uuid.New(), PID, cmd, team.Name, minerType, defaultInterval, defaultMineAmount}
	return miner
}

func (miner *Miner) Mine(team *Team, index int) {
	cont := true
	for cont {
		for _, m := range team.DiscoveredNodes[index].Node.Miners {
			cont = false
			time.Sleep(time.Duration(miner.interval) * time.Second)
			if m.minerId == miner.minerId {
				if miner.minerType == "Bandwidth" {
					team.Bandwidth = team.Bandwidth + miner.amount
				}
				if miner.minerType == "IO" {
					team.Io = team.Io + miner.amount
				}
				if miner.minerType == "Entropy" {
					team.Entropy = team.Entropy + miner.amount
				}
				if miner.minerType == "CPU" {
					team.Cpu = team.Cpu + miner.amount
				}
				cont = true
				break
			}
		}
	}
}
