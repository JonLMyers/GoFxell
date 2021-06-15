package client

import "github.com/c-bata/go-prompt"

func teamCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "Red", Description: "Select Red Team"},
		{Text: "Blue", Description: "Select Blue Team"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func ActionCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "scan", Description: "scan {ip address}"},
		{Text: "exploit", Description: "exploit {Platform/Service} {ip address}"},
		{Text: "connect", Description: "connect {ip address"},
		{Text: "show targets", Description: "show targets"},
		{Text: "show routes", Description: "show routes {ip address}"},
		{Text: "show resources", Description: "show resources"},
		{Text: "exit", Description: "Exit CLI"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

/*-----------------------------------------------------------------------------*/

func cmdCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "show routes", Description: "show routes"},
		{Text: "show proc", Description: "show processes"},
		{Text: "show logs", Description: "show logs"},
		{Text: "deploy miner", Description: "deploy miner {Entropy/CPU/Io/Bandwidth}"},
		{Text: "kill", Description: "kill {Process ID (PID)"},
		{Text: "exit", Description: "Exit CLI"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
