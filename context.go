package slackbot

import (
	"github.com/slack-go/slack"
	"golang.org/x/net/context"
)

const (
	BOT_CONTEXT      = "__BOT_CONTEXT__"
	MESSAGE_CONTEXT  = "__MESSAGE_CONTEXT__"
	REACTION_CONTEXT = "__REACTION_CONTEXT__"
	REACTION_EVENT   = "__REACTION_EVENT__"
	BOT_DEBUG        = "__BOT_DEBUG__"
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

// AddReactionAddedToContext
func AddReactionAddedToContext(ctx context.Context, react *slack.ReactionAddedEvent) context.Context {
	nctx := context.WithValue(ctx, REACTION_CONTEXT, "Added")
	return context.WithValue(nctx, REACTION_EVENT, react)
}

// AddReactionAddedToContext
func AddReactionRemovedToContext(ctx context.Context, react *slack.ReactionRemovedEvent) context.Context {
	nctx := context.WithValue(ctx, REACTION_CONTEXT, "Removed")
	return context.WithValue(nctx, REACTION_CONTEXT, react)
}

func ReactionTypeFromContext(ctx context.Context) string {
	if result, ok := ctx.Value(REACTION_CONTEXT).(string); ok {
		return result
	}
	return ""
}
func ReactionAddedFromContext(ctx context.Context) *slack.ReactionAddedEvent {
	if result, ok := ctx.Value(REACTION_CONTEXT).(*slack.ReactionAddedEvent); ok {
		return result
	}
	return nil
}

func ReactionRemovedFromContext(ctx context.Context) *slack.ReactionRemovedEvent {
	if result, ok := ctx.Value(REACTION_CONTEXT).(*slack.ReactionRemovedEvent); ok {
		return result
	}
	return nil
}

func SetDebug(ctx context.Context) context.Context {
	return context.WithValue(ctx, BOT_DEBUG, true)
}

func IsDebug(ctx context.Context) bool {
	if result, ok := ctx.Value(BOT_DEBUG).(bool); ok {
		return result
	}
	return false
}
