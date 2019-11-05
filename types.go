package slackbot

import (
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

// Matcher type for matching message routes
type Matcher interface {
	Match(context.Context) (bool, context.Context)
	SetBotID(botID string)
}
