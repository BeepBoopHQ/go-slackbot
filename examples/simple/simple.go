package main

import (
	"os"

	"golang.org/x/net/context"

	"github.com/BeepBoopHQ/go-slackbot"
)

func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))

	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention).Subrouter()
	toMe.Hear("(?i)(hi|hello).*").Handler(HelloHandler)
	bot.Hear("(?i)how are you(.*)").Handler(HowAreYouHandler)
	bot.Run()
}

func HelloHandler(ctx context.Context) {
	bot := slackbot.BotFromContext(ctx)
	msg := slackbot.MessageFromContext(ctx)
	bot.ReplyAndType(msg, "Oh hello!")
}

func HowAreYouHandler(ctx context.Context) {
	bot := slackbot.BotFromContext(ctx)
	msg := slackbot.MessageFromContext(ctx)
	bot.ReplyAndType(msg, "A bit tired. You get it? A bit?")
}
