## go-slackbot - Build Slackbots in Go

The go-slackbot project hopes to ease development of Slack bots by adding helpful
methods and a mux-router style interface to the github.com/nlopes/slack package.

Incoming Slack RTM events are mapped to a handler in the following form:

	bot.Hear("(?i)how are you(.*)").MessageHandler(HowAreYouHandler)

In addition to several useful functions in the utils.go file, the slackbot.Bot struct provides handy Reply and ReplyWithAttachments methods:

	func HowAreYouHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
		bot.Reply(evt, "A bit tired. You get it? A bit?", slackbot.WithTyping)
	}
&nbsp;

	func HowAreYouAttachmentsHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
		txt := "Beep Beep Boop is a ridiculously simple hosting platform for your Slackbots."
		attachment := slack.Attachment{
			Pretext:   "We bring bots to life. :sunglasses: :thumbsup:",
			Title:     "Host, deploy and share your bot in seconds.",
			TitleLink: "https:beepboophq.com/",
			Text:      txt,
			Fallback:  txt,
			ImageURL:  "https:storage.googleapis.com/beepboophq/_assets/bot-1.22f6fb.png",
			Color:     "#7CD197",
		}

		attachments := []slack.Attachment{attachment}
		bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
	}
  
But wait, there's more! Well, until there's more, the slackbot package exposes github.com/nlopes/slack RTM and Client objects enabling a consumer to interact with the lower level package directly:

    func HowAreYouHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
      bot.RTM.NewOutgoingMessage("Hello", "#random")
    }


If you want to kick the tires, we would love feedback. Check out these two examples:

- [simple.go](https://github.com/BeepBoopHQ/go-slackbot/blob/master/examples/simple/simple.go)
- [wit.go](https://github.com/BeepBoopHQ/go-slackbot/blob/master/examples/wit/wit.go).
