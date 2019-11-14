package main

import (
	"./input"
	"errors"
	"fmt"
	"strconv"
)

// Insert your own Discord client ID here, if you want to use your own application with own custom game presets.
const DiscordClientId = "644313712567648287"

func main() {
	// Set console title.
	err := SetConsoleTitle("Switchcord v1.0")
	if err != nil {
		fmt.Println("failed to change console title: ", err)
	}
	// Read game from user input.
	client := selectGame()
	if client == nil {
		return
	}
	// Wait for exit.
	fmt.Println("Press enter to exit...")
	input.DefaultInput.WaitForInput()
	// Close Discord connection.
	err = client.Close()
	if err != nil {
		fmt.Println("close error: ", err)
	}
}

func selectGame() *DiscordClient {
	fmt.Println("Search Nintendo Switch game by name: ")
	name, err := input.DefaultInput.ReadString()
	if err != nil {
		fmt.Println("input error: ", err)
		return nil
	}

	// Search for game at IGDB (only Nintendo Switch games)
	gameList, err := SearchGame(name)
	if err != nil {
		fmt.Println("parent error: ", err)
		fmt.Println("child error: ", errors.Unwrap(err))
		return nil
	}

	// Let user select game if more than one game has been found.
	count := len(gameList)
	var index int
	if count == 1 {
		index = 0
	} else if count < 1 {
		fmt.Println("No games with the specified name found. Check spelling and go to https://igdb.com to check if it exists.")
		return nil
	} else {
		// List games returned by the API.
		fmt.Println("Select a game: ")
		fmt.Println("Try to state the game name more precisely if it is missing from the list!")
		for i := 0; i < len(gameList); i++ {
			fmt.Println(strconv.Itoa(i + 1) + ") " + gameList[i].Name)
		}

		index, err = input.DefaultInput.ReadInteger()
		if err != nil {
			fmt.Println("input error: ", err)
			return nil
		}

		index = index - 1
		if index < 0 || index > len(gameList) {
			fmt.Println("Selection is too small or to big. Must be between 1 and " + strconv.Itoa(len(gameList)))
			return nil
		}
	}

	// Print selected game.
	game := gameList[index]
	fmt.Println("Selected game: " + game.Name)

	// Set activity.
	client, err := NewClient(DiscordClientId)
	if err != nil {
		fmt.Println("failed to construct discord client: ", err)
		return nil
	}

	err = client.SetActivity(game)
	if err != nil {
		_ = client.Close()
		fmt.Println("failed to update profile status: ", err)
		return nil
	}

	return client
}