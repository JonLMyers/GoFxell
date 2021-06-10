package client

import (
	"fmt"

	"github.com/JonLMyers/GoFxell/game"
	"github.com/c-bata/go-prompt"
)

var playerTeam *game.Team
var gameMap *game.Map
var n game.Node

func SinglePlayer() {
	gameMap = game.NewMap("maps.json") //NewMapFromFile
	teamName := prompt.Input("Select Team Name (Red/Blue)> ", teamCompleter)

	/* Functional Parameters with optional params */
	/* https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis */
	playerTeam = game.NewTeam(teamName, gameMap)

	/* Main gameplay loop */
	// https://youtu.be/-GV814cWiAw
	fmt.Println(playerTeam.View(playerTeam.StartNode))
	playerName := "player"
	currentSystem := "fxell"
	promptPrefix := fmt.Sprintf("%s@%s:~$ ", playerName, currentSystem)

	//cmd := prompt.Input(gamePrompt, actionCompleter)
	//cmd = strings.TrimSpace(cmd)
	//cmdParts := strings.Split(cmd, " ")

	gamePrompt := prompt.New(
		Executor,
		ActionCompleter,
		prompt.OptionPrefix(promptPrefix),
		prompt.OptionPrefixTextColor(prompt.Green),
	)
	gamePrompt.Run()
}

func Connect(ipAddr string, team *game.Team, gameMap game.Map) (bool, error) {
	index, err := team.DiscoveredNodes.IndexOf(ipAddr)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	n = team.DiscoveredNodes[index].Node
	if team.DiscoveredNodes[index].NodeOwned {
		playerName := "player"
		currentSystem := n.IPAddr
		cmdPrefix := fmt.Sprintf("%s@%s:~$ ", playerName, currentSystem)
		cmdPrompt := prompt.New(
			CmdExecutor,
			cmdCompleter,
			prompt.OptionPrefix(cmdPrefix),
			prompt.OptionPrefixTextColor(prompt.Red),
			prompt.OptionSetExitCheckerOnInput(ExitChecker),
		)
		cmdPrompt.Run()
		return true, nil
	} else {
		return false, nil
	}
}

func ExitChecker(in string, breakline bool) bool {
	return in == "exit"
}
