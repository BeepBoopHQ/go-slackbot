package slackbot

import (
	"regexp"

	"github.com/nlopes/slack"
)

// StripDirectMention removes a leading mention (aka direct mention) from a message string
func StripDirectMention(text string) string {
	return regexp.MustCompile(`(^<@[a-zA-Z0-9]+>[\:]*[\s]*)?(.*)`).FindStringSubmatch(text)[2]
}

// IsDirectMessage returns true if this message is in a direct message conversation
func IsDirectMessage(evt *slack.MessageEvent) bool {
	return regexp.MustCompile("^D.*").MatchString(evt.Channel)
}

// IsDirectMention returns true is message is a Direct Mention that mentions a specific user. A
// direct mention is a mention at the very beginning of the message
func IsDirectMention(evt *slack.MessageEvent, userID string) bool {
	return regexp.MustCompile("^<@" + userID + ">.*").MatchString(evt.Text)
}

// IsMentioned returns true if this message contains a mention of a specific user
func IsMentioned(evt *slack.MessageEvent, userID string) bool {
	userIDs := WhoMentioned(evt)
	for _, u := range userIDs {
		if u == userID {
			return true
		}
	}
	return false
}

// IsMention returns true the message contains a mention
func IsMention(evt *slack.MessageEvent) bool {
	results := regexp.MustCompile(`<@([a-zA-z0-9]+)?>`).FindAllStringSubmatch(evt.Text, -1)
	return len(results) > 0
}

// WhoMentioned returns a list of userIDs mentioned in the message
func WhoMentioned(evt *slack.MessageEvent) []string {
	results := regexp.MustCompile(`<@([a-zA-z0-9]+)>`).FindAllStringSubmatch(evt.Text, -1)
	matches := make([]string, len(results))
	for i, r := range results {
		matches[i] = r[1]
	}
	return matches
}
