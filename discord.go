package main

import (
	"fmt"
	"github.com/ananagame/rich-go/client"
	"time"
)

// DiscordClient represents a client able to display a rich-presence text on the users Discord profile.
type DiscordClient struct {
	clientId string
}

// The default discord client identifier is bundled with a set of supported games, which means it is possible to display the game's cover image.
var gameImages = []string{"pokemon-sword", "pokemon-shield", "super-smash-bros-ultimate", "tetris-99", "fire-emblem-three-houses", "pokemon-lets-go-eevee", "pokemon-lets-go-pikachu", "the-legend-of-zelda-breath-of-the-wild"}
var gameImageMap = map[string]string {"the-legend-of-zelda-breath-of-the-wild": "breath-of-the-wild"}

// Creates a new instance of the DiscordClient with the specified client identifier referring to the application identifier.
// The default client identifier refers to an application with a default preset of available game images.
func NewClient(clientId string) (*DiscordClient, error) {
	// Try to log in to Discord with the client identifier.
	err := client.Login(clientId)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}

	// Create instance of DiscordClient and pass client identifier.
	return &DiscordClient{clientId}, nil
}

// SetActivity sets the activity displayed on the users Discord profile.
// If an image of the current game is available it is displayed in the rich-presence presentation.
// Else a generic image is used.
func (dc *DiscordClient) SetActivity(game Game) error {
	// Check if game image is available, else use generic image.
	imageTag := "nintendo-switch"
	if imageAvailable(&game) {
		imageTag = game.Slug
		// Try to map large game names to smaller image tags.
		if val, ok := gameImageMap[imageTag]; ok {
			imageTag = val
		}
	}

	err := client.SetActivity(client.Activity{
		Details: game.Name,
		State: "Spielt ein Nintendo Switch-Spiel",
		LargeImage: imageTag,
		LargeText: game.Name,
		Timestamps: &client.Timestamps{
			Start: time.Now(), // Game started since now.
			End: time.Unix(0, 0), // If not specified, time display will be buggy.
		},
	})
	return err
}

// Close the connection to Discord.
func (dc *DiscordClient) Close() error {
	client.Logout()
	return nil
}

// Checks whether there is an uploaded game cover available to display.
func imageAvailable(game *Game) bool {
	for _, n := range gameImages {
		if game.Slug == n {
			return true
		}
	}
	return false
}