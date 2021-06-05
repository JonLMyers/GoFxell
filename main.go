package main

import (
	"fmt"
	"strings"

	"github.com/JonLMyers/GoFxell/action"
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
func actionCompleter(d prompt.Document) []prompt.Suggest {
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

//var playerTeam *game.Team
//var gameMap *game.Map

func main() {
	gameMap := game.NewMap("maps.json") //NewMapFromFile
	teamName := prompt.Input("Select Team Name (Red/Blue)> ", teamCompleter)

	/* Functional Parameters with optional params */
	/* https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis */
	playerTeam := game.NewTeam(teamName, gameMap)

	/* Main gameplay loop */
	// https://youtu.be/-GV814cWiAw
	fmt.Println(playerTeam.View(playerTeam.StartNode))
	for {
		playerName := "player"
		currentSystem := "fxell"
		gamePrompt := fmt.Sprintf("%s@%s:~$ ", playerName, currentSystem)

		cmd := prompt.Input(gamePrompt, actionCompleter)
		cmd = strings.TrimSpace(cmd)
		cmdParts := strings.Split(cmd, " ")

		if strings.HasPrefix(cmd, "scan") {
			action.Scan(cmdParts[1], playerTeam)
			continue
		}
		if strings.HasPrefix(cmd, "exploit") {
			if len(cmdParts) < 3 {
				fmt.Println("Invalid Exploit Syntax")
				continue
			}
			//Create a prompt here and for connect... oh wait doesn't work.
			if cmdParts[1] == "Windows" || cmdParts[1] == "Nix" {
				ok, _ := action.PlatformExploit(cmdParts[2], playerTeam)
				if !ok {
					fmt.Println("Platform Exploit Failed")
					continue
				}
				fmt.Println("Platform Exploit Successful")
				action.ShowTargets(*playerTeam)
				continue
			}
			fmt.Println("Service Exploit")
			continue
		}

		if strings.HasPrefix(cmd, "show targets") {
			action.ShowTargets(*playerTeam)
			continue
		}

		if strings.HasPrefix(cmd, "show resources") {
			resources := playerTeam.GetResources()
			fmt.Println(resources)
			continue
		}

		if strings.HasPrefix(cmd, "connect") {
			if len(cmdParts) < 2 {
				fmt.Println("Invalid Connect Syntax")
				continue
			}
			ok, err := action.Connect(cmdParts[1], playerTeam, *gameMap)
			if !ok || err != nil {
				fmt.Println("Connection Refused")
				continue
			}
			fmt.Println("Connection Closed")
			continue
		}

		if strings.HasPrefix(cmd, "exit") {
			break
		}

		fmt.Println("Invalid Command")
	}

	//fmt.Printf("%v\n", startNodes)
	//fmt.Printf("%+v\n", playerTeam.StartNode)
	// We will need to remove some of the extra info (lootrate) when we return the data back to the client.
	//playerTeamNodeView := playerTeam.View(playerTeam.StartNode)
	//fmt.Printf("%+v\n", playerTeamNodeView)

	//ctx := context.Background()
	//teamCTX := game.WithViewer(ctx, playerTeam)
	//game.DoThing(teamCTX, playerTeam.StartNode)

	//action.Scan("10.0.0.1", *playerTeam)
	//reqNode := gameMap.FindNode("10.0.0.1")
	//playerTeamNodeView = playerTeam.View(reqNode)
	//fmt.Printf("%+v\n", playerTeam.DiscoveredNodes)

}
