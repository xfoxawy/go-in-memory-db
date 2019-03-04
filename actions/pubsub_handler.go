package actions

import (
	"strings"

	"github.com/go-in-memory-db/pubsub"
)

func (a *Actions) publishHandler() string {
	subject := a.StringArray[1]

	msg := strings.Join(a.StringArray[2:], "")

	pub := &pubsub.Publisher{
		Url:     "",
		Subject: subject,
		Msg:     msg,
	}
	return pub.Publish()
}

func (a *Actions) subscribeHandler() {
	subject := a.StringArray[1]

	sub := &pubsub.AsyncSubscriber{
		Url:     "",
		Subject: subject,
	}
	sub.ASub()
}
