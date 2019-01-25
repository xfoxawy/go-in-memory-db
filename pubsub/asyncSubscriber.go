package pubsub

// Import Go and NATS packages
import (
	"log"
	"runtime"

	"github.com/nats-io/go-nats"
)

// const (
// 	DEFAULTURL string = nats.DefaultURL
// )

type AsyncSubscriber struct {
	url     string
	subject string
}

func (as *AsyncSubscriber) ASub() {

	if as.url == "" {
		as.url = DEFAULTURL
	}

	natsConnection, _ := nats.Connect(as.url)
	log.Println("Connected to " + as.url)

	log.Printf("Subscribing to subject" + as.subject + "\n")
	natsConnection.Subscribe(as.subject, func(msg *nats.Msg) {
		log.Printf("Received message '%s\n", string(msg.Data)+"'")
	})

	// Keep the connection alive
	runtime.Goexit()

}
