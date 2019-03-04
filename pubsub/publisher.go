package pubsub

// Import packages
import (
	"log"

	"github.com/nats-io/go-nats"
)

const (
	// DEFAULTURL const
	DEFAULTURL string = nats.DefaultURL
)

// Publisher struct
type Publisher struct {
	Url     string
	Subject string
	Msg     string
}

// Publish function
func (p *Publisher) Publish() string {
	if p.Url == "" {
		p.Url = DEFAULTURL
	}

	natsConnection, err := nats.Connect(p.Url)
	if err != nil {
		log.Println("not Connected to " + p.Url)
	}
	defer natsConnection.Close()

	// Publish message on subject
	natsConnection.Publish(p.Subject, []byte(p.Msg))

	return "Connected to " + p.Url + "\n" + "Published message on subject " + p.Subject
}
