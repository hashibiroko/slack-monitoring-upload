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
	token := flag.String("token", os.Getenv("SLACK_BOT_TOKEN"), "Set Slack Bot Token")
	channelName := flag.String("channel", "general", "Set the channel name on which the message is posted")
	flag.Parse()

	api := slack.New(*token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	channelId := ""
	fileId := ""
	channels, _ := rtm.GetChannels(false)

	for _, channel := range channels {
		if channel.Name == *channelName {
			channelId = channel.ID
			break
		}
	}

	if channelId == "" {
		log.Fatal(fmt.Sprintf("Not found channel <%s>\n", *channelName))
		os.Exit(1)
	}

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
					file, _, _, _ := rtm.GetFileInfo(fileId, 1, 1)
					text := fmt.Sprintf("File uploaded <%s|%s> (%dKB) by <@%s>",
						file.URLPrivate, file.Name, file.Size/1000, file.User)
					params := slack.PostMessageParameters{
						AsUser:      true,
						UnfurlLinks: false,
						UnfurlMedia: false,
					}
					rtm.PostMessage(channelId, text, params)
				}()
			default:
			}
		}
	}
}
