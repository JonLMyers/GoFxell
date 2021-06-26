package client

type TargetResponse struct {
	IPAddr    string
	Platform  string
	Services  []string
	Routes    []string
	NodeOwned bool
}

type LogResponse struct {
	Id         string
	Timestamp  string
	Type       string
	LogMessage string
}

type ProcessResponse struct {
	PID  string
	CMD  string
	Team string
}

type SingleResponse struct {
	Message string
}

type MultiResponse struct {
	Messages []string
}
