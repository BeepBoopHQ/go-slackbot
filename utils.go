package slackbot

import "regexp"

func StripDirectMention(text string) string {
	return regexp.MustCompile(`(^<@[a-zA-Z0-9]+>[\:]*[\s]*)?(.*)`).FindStringSubmatch(text)[2]
}
