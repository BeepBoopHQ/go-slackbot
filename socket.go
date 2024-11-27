package slackbot

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"golang.org/x/net/context"
)

func (b *Bot) RunSocketMode() {

	b.Socket = socketmode.New(b.Client,
		socketmode.OptionDebug(b.Debug),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)))

	go func() {
		for evt := range b.Socket.Events {
			ctx := context.Background()
			ctx = AddBotToContext(ctx, b)
			if b.Debug {
				ctx = SetDebug(ctx)
				log.Printf("Event Type: %s", evt.Type)
			}

			switch evt.Type {
			case socketmode.EventTypeConnecting:
				fmt.Println("Connecting to Slack with Socket Mode...")
				// conn := evt.Data.(socketmode.ConnectedEvent)\
			case socketmode.EventTypeConnectionError:
				fmt.Println("Connection failed. Retrying later...")
			case socketmode.EventTypeConnected:
				fmt.Println("Connected to Slack with Socket Mode.")
			case socketmode.EventTypeHello:
				fmt.Println("Slack says hello")
			case socketmode.EventTypeEventsAPI:
				eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
				if !ok {
					fmt.Printf("Ignored %+v\n", evt)
					continue
				}

				fmt.Printf("Event received: %+v\n", eventsAPIEvent)

				b.Socket.Ack(*evt.Request)

				switch eventsAPIEvent.Type {
				case slackevents.CallbackEvent:
					innerEvent := eventsAPIEvent.InnerEvent
					switch ev := innerEvent.Data.(type) {
					case *slackevents.MessageEvent:
						if b.botUserID == ev.User {
							continue
						}
						ctx = AddMessageToContext(ctx, b.ConvertMessageEvent(ev))
						var match RouteMatch
						if matched, ctx := b.Match(ctx, &match); matched {
							match.Handler(ctx)
						}
					case *slackevents.ReactionAddedEvent:
						// Handle reaction events
						if b.botUserID == ev.User {
							continue
						}

						ctx = AddReactionAddedToContext(ctx, ConvertReactionAdded(ev))
						var match RouteMatch
						if matched, ctx := b.Match(ctx, &match); matched {
							match.Handler(ctx)
						}

					case *slackevents.ReactionRemovedEvent:
						// Handle reaction events
						if b.botUserID == ev.User {
							continue
						}

						ctx = AddReactionRemovedToContext(ctx, ConvertReactionRemoved(ev))
						var match RouteMatch
						if matched, ctx := b.Match(ctx, &match); matched {
							match.Handler(ctx)
						}

					}
				default:
					b.Socket.Debugf("unsupported Events API event received")
				}
			case socketmode.EventTypeInteractive:
				// callback, ok := evt.Data.(slack.InteractionCallback)
				// if !ok {
				// 	fmt.Printf("Ignored %+v\n", evt)

				// 	continue
				// }

				// fmt.Printf("Interaction received: %+v\n", callback)

				// var payload interface{}

				// switch callback.Type {
				// case slack.InteractionTypeBlockActions:
				// 	// See https://api.slack.com/apis/connections/socket-implement#button

				// 	client.Debugf("button clicked!")
				// case slack.InteractionTypeShortcut:
				// case slack.InteractionTypeViewSubmission:
				// 	// See https://api.slack.com/apis/connections/socket-implement#modal
				// case slack.InteractionTypeDialogSubmission:
				// default:

				// }

				// client.Ack(*evt.Request, payload)
			case socketmode.EventTypeSlashCommand:
				// cmd, ok := evt.Data.(slack.SlashCommand)
				// if !ok {
				// 	fmt.Printf("Ignored %+v\n", evt)

				// 	continue
				// }

				// client.Debugf("Slash command received: %+v", cmd)

				// payload := map[string]interface{}{
				// 	"blocks": []slack.Block{
				// 		slack.NewSectionBlock(
				// 			&slack.TextBlockObject{
				// 				Type: slack.MarkdownType,
				// 				Text: "foo",
				// 			},
				// 			nil,
				// 			slack.NewAccessory(
				// 				slack.NewButtonBlockElement(
				// 					"",
				// 					"somevalue",
				// 					&slack.TextBlockObject{
				// 						Type: slack.PlainTextType,
				// 						Text: "bar",
				// 					},
				// 				),
				// 			),
				// 		),
				// 	}}

				// client.Ack(*evt.Request, payload)
			default:
				fmt.Fprintf(os.Stderr, "Unexpected event type received: %s\n", evt.Type)
			}
		}
	}()

	b.Socket.Run()
}

func (b *Bot) ConvertMessageEvent(evt *slackevents.MessageEvent) *slack.MessageEvent {

	msg := slack.MessageEvent{
		Msg: slack.Msg{
			ClientMsgID:     evt.ClientMsgID,
			Type:            evt.Type,
			User:            evt.User,
			Text:            evt.Text,
			ThreadTimestamp: evt.ThreadTimeStamp,
			Timestamp:       evt.TimeStamp,
			Channel:         evt.Channel,
			// ChannelType
			EventTimestamp: evt.EventTimeStamp,
			// UserTeam
			// SourceTeam
			SubType:  evt.SubType,
			BotID:    evt.BotID,
			Username: evt.Username,
			Icons:    (*slack.Icon)(evt.Icons),
			Upload:   evt.Upload,
			// Files:    evt.Files,

		},
	}

	return &msg
}

func ConvertReactionAdded(evt *slackevents.ReactionAddedEvent) *slack.ReactionAddedEvent {

	react := slack.ReactionAddedEvent{
		Type:           evt.Type,
		User:           evt.User,
		ItemUser:       evt.ItemUser,
		Reaction:       evt.Reaction,
		EventTimestamp: evt.EventTimestamp,
		Item: slack.ReactionItem{
			Type:      evt.Item.Type,
			Channel:   evt.Item.Channel,
			Timestamp: evt.Item.Timestamp,
		},
	}

	return &react
}

func ConvertReactionRemoved(evt *slackevents.ReactionRemovedEvent) *slack.ReactionRemovedEvent {

	react := slack.ReactionRemovedEvent{
		Type:           evt.Type,
		User:           evt.User,
		ItemUser:       evt.ItemUser,
		Reaction:       evt.Reaction,
		EventTimestamp: evt.EventTimestamp,
		Item: slack.ReactionItem{
			Type:      evt.Item.Type,
			Channel:   evt.Item.Channel,
			Timestamp: evt.Item.Timestamp,
		},
	}

	return &react
}
