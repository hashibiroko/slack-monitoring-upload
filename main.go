package main

import (
	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
	"flag"
	"os"
	"fmt"
	"log"
)

func main() {
	godotenv.Load()
	token := flag.String("token", os.Getenv("SLACK_USER_TOKEN"),
		"Set Slack User Token")
	channelName := flag.String("channel", os.Getenv("SLACK_CHANNEL_NAME"),
		"Set the channel name on which the message is posted")
	flag.Parse()

	if *channelName == "" {
		*channelName = "general"
	}

	api := slack.New(*token)
	if _, err := api.AuthTest(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fileId := ""
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.FilePublicEvent:
				if fileId == ev.FileID {
					break
				}
				fileId = ev.FileID
				go func() {
					channel, err := api.JoinChannel(*channelName)
					if err != nil {
						log.Println(err)
						return
					}
					file, _, _, _ := rtm.GetFileInfo(fileId, 1, 1)
					text := fmt.Sprintf("File uploaded <%s|%s> (%dKB) by <@%s>",
						file.URLPrivate, file.Name, file.Size/1000, file.User)
					params := slack.PostMessageParameters{
						AsUser:      true,
						UnfurlLinks: false,
						UnfurlMedia: false,
					}
					rtm.PostMessage(channel.ID, text, params)
				}()
			default:
			}
		}
	}
}
