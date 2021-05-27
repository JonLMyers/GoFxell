package main

import (
	"fmt"

	"github.com/JonLMyers/GoFxell/game"
	"github.com/c-bata/go-prompt"
)

func teamCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "Red", Description: "Select Red Team"},
		{Text: "Blue", Description: "Select Blue Team"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	gameMap := game.NewMap("maps.json")
	teamName := prompt.Input("Select Team Name (Red/Blue)> ", teamCompleter)

	/* Filters the StartNodes to be assigned to different teams*/
	startNodes := game.FilterStartNodes(gameMap.Nodes)

	/* Functional Parameters with optional params */
	/* https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis */
	playerTeam := game.NewTeam(teamName, func(t *game.Team) {
		t.StartNode, startNodes = startNodes.Select()
	})

	/* Main gameplay loop */

	fmt.Printf("%v\n", startNodes)
	fmt.Printf("%+v\n", playerTeam.StartNode)

}
