package slackbot

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	"github.com/nlopes/slack"
)

func New(slackToken string) *Bot {
	b := &Bot{Client: slack.New(slackToken)}
	return b
}

type Bot struct {
	SimpleRouter
	// Routes to be matched, in order.
	routes []*Route
	// Slack UserID of the bot UserID
	botUserID string
	// Slack API
	Client *slack.Client
	RTM    *slack.RTM
}

// Run start the bot
func (b *Bot) Run() {
	b.RTM = b.Client.NewRTM()
	go b.RTM.ManageConnection()
	for {
		select {
		case msg := <-b.RTM.IncomingEvents:
			ctx := context.Background()
			ctx = AddBotToContext(ctx, b)
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Printf("Connected: %#v\n", ev.Info.User)
				b.SetBotID(ev.Info.User.ID)
			case *slack.MessageEvent:
				// ignore messages from the current user, the bot user
				if b.botUserID == ev.User {
					continue
				}

				fmt.Printf("Message: %#v\n", ev.Text)
				ctx = AddMessageToContext(ctx, ev)
				var match RouteMatch
				if matched, ctx := b.Match(ctx, &match); matched {
					match.Handler(ctx)
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break

			default:

				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}

}

// Reply to a message and simulate typing through the realtime messaging API
func (b *Bot) ReplyAndType(evt *slack.MessageEvent, text string) {
	b.Type(evt.Channel, text)
	b.Reply(evt, text)
}

// Reply to a message through the realtime messaging API
func (b *Bot) Reply(evt *slack.MessageEvent, text string) {
	b.RTM.SendMessage(b.RTM.NewOutgoingMessage(text, evt.Channel))
}

// Type sends a typing message and calls time.Sleep to simulate a delay
func (b *Bot) Type(channel, text string) {
	sleepDuration := time.Minute * time.Duration(len(text)) / 3000
	b.RTM.SendMessage(b.RTM.NewTypingMessage(channel))
	time.Sleep(sleepDuration)
}
