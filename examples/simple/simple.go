package main

import (
	"os"

	"golang.org/x/net/context"

	"github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
)

func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))

	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention).Subrouter()
	toMe.Hear("(?i)(hi|hello).*").MessageHandler(HelloHandler)
	bot.Hear("(?i)how are you(.*)").MessageHandler(HowAreYouHandler)
	bot.Run()
}

func HelloHandler(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent) {
	bot.ReplyAndType(msg, "Oh hello!")
}

func HowAreYouHandler(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent) {
	bot.ReplyAndType(msg, "A bit tired. You get it? A bit?")
}
