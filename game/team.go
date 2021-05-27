package game

const (
	defaultBandwidth = 100
	defaultIO        = 100
	defaultCPU       = 100
	defaultEntropy   = 100
)

type Team struct {
	Name      string
	Bandwidth int
	Io        int
	Cpu       int
	Entropy   int
	StartNode Node
}

func NewTeam(name string, options ...func(*Team)) Team {
	team := Team{name, defaultBandwidth, defaultIO, defaultCPU, defaultEntropy, Node{}}
	for _, opt := range options {
		opt(&team)
	}
	return team
}
