package actions

import (
	"strings"

	"github.com/go-in-memory-db/pubsub"
)

func (a *Actions) publishHandler() string {
	subject := a.StringArray[1]

	msg := strings.Join(a.StringArray[2:], "")

	publisher := &pubsub.Publisher{
		"",
		subject,
		msg,
	}
	return publisher.Publish()
}

// func (a *Actions) subscribeHandler() string {
//
// }
