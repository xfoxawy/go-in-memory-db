package pubsub

// Import Go and NATS packages
import (
	"github.com/go-in-memory-db/connection"
	"github.com/nats-io/go-nats"
)

// AsyncSubscriber struct
type AsyncSubscriber struct {
	Url     string
	Subject string
}

// ASub function
func (as *AsyncSubscriber) ASub() {

	if as.Url == "" {
		as.Url = DEFAULTURL
	}

	natsConnection, _ := nats.Connect(as.Url)
	c := connection.GetSharedConnection()
	c.WriteString("Connected to " + as.Url)
	c.WriteString("Subscribing to subject " + as.Subject)

	natsConnection.Subscribe(as.Subject, func(msg *nats.Msg) {
		connection.ShareConnection(c)
		c.WriteString("Received message \n" + string(msg.Data))
	})
}
