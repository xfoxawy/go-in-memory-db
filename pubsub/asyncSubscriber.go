package pubsub

// Import Go and NATS packages
import (
	"log"
	"runtime"

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
	log.Println("Connected to " + as.Url)

	log.Printf("Subscribing to subject " + as.Subject + "\n")
	natsConnection.Subscribe(as.Subject, func(msg *nats.Msg) {
		log.Printf("Received message '%s\n", string(msg.Data)+"'")
	})

	// Keep the connection alive
	runtime.Goexit()
}
