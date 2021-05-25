package main

import (
	"fmt"

	slack "github.com/pendo-io/nlopes-slack-fork"
)

func main() {
	api := slack.New("YOUR_TOKEN_HERE")
	channels, err := api.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, channel := range channels {
		fmt.Println(channel.Name)
		// channel is of type conversation & groupConversation
		// see all available methods in `conversation.go`
	}
}
