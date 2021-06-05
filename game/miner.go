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
	teamName  string
	minerType string
	interval  int
	amount    int
}

func NewMiner(minerType string, team Team) Miner {
	miner := Miner{uuid.New(), team.Name, minerType, defaultInterval, defaultMineAmount}
	return miner
}

func (miner *Miner) Mine(team *Team, node Node) {
	cont := true
	for cont {
		for _, m := range node.Miners {
			cont = false
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
		time.Sleep(time.Duration(miner.interval) * time.Second)
	}
}
