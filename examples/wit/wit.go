package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/chris-skud/go-wit"
	"github.com/nlopes/slack"
)

func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))
	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention).Preprocess(WitPreprocess).Subrouter()
	toMe.AddMatcher(&IntentMatcher{intent: "hello"}).MessageHandler(HelloHandler)
	toMe.AddMatcher(&IntentMatcher{intent: "how_are_you"}).MessageHandler(HowAreYouHandler)
	toMe.MessageHandler(ConfusedHandler)
	bot.Run()
}

func HelloHandler(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent) {
	bot.Reply(msg, "Oh hello!", slackbot.WithTyping)
}

func HowAreYouHandler(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent) {
	bot.Reply(msg, "A bit tired. You get it? A bit?", slackbot.WithTyping)
}

func ConfusedHandler(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent) {
	bot.Reply(msg, "I don't understand 😰", slackbot.WithTyping)
}

func WitPreprocess(ctx context.Context) context.Context {
	msg := slackbot.MessageFromContext(ctx)
	text := slackbot.StripDirectMention(msg.Text)
	req := &wit.MessageRequest{Query: text}
	witMessage, err := wit.NewClient(os.Getenv("WIT_TOKEN")).Message(req)
	if err != nil {
		bot := slackbot.BotFromContext(ctx)
		bot.Reply(msg, "Uh oh, I seem to be out of sorts :dizzy_face", slackbot.WithTyping)
		return ctx
	}
	fmt.Printf("WIT: %#v\n", witMessage)
	ctx = AddWitToContext(ctx, witMessage)

	return ctx
}

type IntentMatcher struct {
	slackbot.RegexpMatcher
	intent     string
	confidence float32
}

func (it *IntentMatcher) Match(ctx context.Context) (bool, context.Context) {
	// default confidence to 50%
	if it.confidence == 0 {
		it.confidence = 0.5
	}
	witMessage := WitFromContext(ctx)
	if witMessage != nil && len(witMessage.Outcomes) > 0 {
		outcome := witMessage.Outcomes[0]
		if outcome.Intent == it.intent && outcome.Confidence >= it.confidence {
			return true, ctx
		}
	}
	return false, ctx
}

func WitFromContext(ctx context.Context) *wit.Message {
	if result, ok := ctx.Value("__WIT__").(*wit.Message); ok {
		return result
	}
	return nil
}

// AddLoggerToContext sets the logger and returns the newly derived context
func AddWitToContext(ctx context.Context, witMessage *wit.Message) context.Context {
	return context.WithValue(ctx, "__WIT__", witMessage)
}
