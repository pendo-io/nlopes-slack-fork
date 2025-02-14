package main

import (
	"fmt"

	slack "github.com/pendo-io/nlopes-slack-fork"
)

func main() {
	api := slack.New("YOUR_TOKEN_HERE")

	userID := "USER_ID"

	_, _, channelID, err := api.OpenIMChannel(userID)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	api.PostMessage(channelID, "Hello World!", slack.PostMessageParameters{})
}
