package slackbot

import (
	"github.com/nlopes/slack"
	"golang.org/x/net/context"
)

const (
	BOT_CONTEXT     = "__BOT_CONTEXT__"
	MESSAGE_CONTEXT = "__MESSAGE_CONTEXT__"
)

func BotFromContext(ctx context.Context) *Bot {
	if result, ok := ctx.Value(BOT_CONTEXT).(*Bot); ok {
		return result
	}
	return nil
}

// AddBotToContext sets the bot reference in context and returns the newly derived context
func AddBotToContext(ctx context.Context, bot *Bot) context.Context {
	return context.WithValue(ctx, BOT_CONTEXT, bot)
}

func MessageFromContext(ctx context.Context) *slack.MessageEvent {
	if result, ok := ctx.Value(MESSAGE_CONTEXT).(*slack.MessageEvent); ok {
		return result
	}
	return nil
}

// AddMessageToContext sets the Slack message event reference in context and returns the newly derived context
func AddMessageToContext(ctx context.Context, msg *slack.MessageEvent) context.Context {
	return context.WithValue(ctx, MESSAGE_CONTEXT, msg)
}
