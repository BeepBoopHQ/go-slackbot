package slackbot

import (
	"testing"

	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

func TestStripDirectMention(t *testing.T) {
	assert := assert.New(t)

	pairs := []string{
		"a message", "a message",
		"    space", "    space",
		"<@USOMEUSER> I 😍 u", "I 😍 u",
		"<@USOMEUSER>       space", "space",
		"<@USOMEUSER>:  abc", "abc",
		"<@USOMEUSER>: you <@USOMEUSER>", "you <@USOMEUSER>",
		"space    ", "space    ",
	}

	for i := 0; i < len(pairs); {
		assert.Equal(pairs[i+1], StripDirectMention(pairs[i]))
		i += 2
	}
}

func TestIsDirectMessage(t *testing.T) {
	assert := assert.New(t)
	msg := &slack.MessageEvent{}

	msg.Channel = "DABCDEF"
	assert.True(true, IsDirectMessage(msg))

	msg.Channel = " DABCDEF"
	assert.False(IsDirectMessage(msg))

	msg.Channel = "UABCDEF"
	assert.False(IsDirectMessage(msg))

	msg.Channel = ""
	assert.False(IsDirectMessage(msg))
}

func TestIsDirectMention(t *testing.T) {
	assert := assert.New(t)
	msg := &slack.MessageEvent{}

	userID := "U12345"
	msg.Text = "<@" + userID + "> some message"
	assert.True(true, IsDirectMention(msg, userID))

	msg.Text = " <@" + userID + "> some othermessage"
	assert.False(IsDirectMention(msg, userID))

	msg.Text = "<@" + userID + "> some longer message <@MENTION> message   "
	assert.True(IsDirectMention(msg, userID))

	msg.Text = ""
	assert.False(IsDirectMention(msg, userID))

	msg.Text = " "
	assert.False(IsDirectMention(msg, userID))
}

func TestWhoMentioned(t *testing.T) {
	assert := assert.New(t)
	msg := &slack.MessageEvent{}

	msg.Text = "<@U123456> hi there <@UABCDEF>"
	assert.Equal([]string{"U123456", "UABCDEF"}, WhoMentioned(msg))

	msg.Text = "<@> <@DD <@U123456> hi there <@UABCDEF> LL"
	assert.Equal([]string{"U123456", "UABCDEF"}, WhoMentioned(msg))
}

func TestIsMention(t *testing.T) {
	assert := assert.New(t)
	msg := &slack.MessageEvent{}

	msg.Text = "<@U123456> hi there <@UABCDEF>"
	assert.True(IsMention(msg))

	msg.Text = "this is something"
	assert.False(IsMention(msg))
}

func TestIsMentioned(t *testing.T) {
	assert := assert.New(t)
	msg := &slack.MessageEvent{}

	msg.Text = "<@U123456> hi there <@UABCDEF>"
	assert.True(IsMentioned(msg, "UABCDEF"))

	msg.Text = "<@U123456> hi there <@UABCDEF>"
	assert.False(IsMentioned(msg, "UABCDE"))

	msg.Text = "<@U123456> hi there <@UABCDEF>"
	assert.False(IsMentioned(msg, "UXXXXXX"))

	msg.Text = "this is something"
	assert.False(IsMentioned(msg, "UAAAAAA"))
}
