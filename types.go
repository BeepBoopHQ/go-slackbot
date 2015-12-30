package slackbot

import (
	"regexp"

	"golang.org/x/net/context"

	"github.com/nlopes/slack"
)

type MessageType string

const (
	DirectMessage MessageType = "direct_message"
	DirectMention MessageType = "direct_mention"
	Mention       MessageType = "mention"
	Ambient       MessageType = "ambient"
)

type Handler func(context.Context)
type MessageHandler func(ctx context.Context, bot *Bot, msg *slack.MessageEvent)
type Preprocessor func(context.Context) context.Context

// IsDM returns true is message is a DM
func IsDirectMessage(evt *slack.MessageEvent) bool {
	return regexp.MustCompile("^D.*").MatchString(evt.Channel)
}

// IsDM returns true is message is a DM
func IsDirectMention(evt *slack.MessageEvent, userID string) bool {
	return regexp.MustCompile("^<@" + userID + ">.*").MatchString(evt.Text)
}
