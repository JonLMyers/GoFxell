package game

type Firewall struct {
	teamName string
}

func NewFirewall(team Team) Firewall {
	firewall := Firewall{team.Name}
	return firewall
}
