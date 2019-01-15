package main

import (
	"fmt"
	nats "github.com/nats-io/go-nats-streaming"
	"time"
)

func main() {
	sc, connectError := nats.Connect("test-cluster", "testclient", nats.NatsURL("nats://localhost:4223"))
	defer sc.Close()
	if connectError != nil {
		panic(connectError)
	}

	sub, subScribeError := sc.Subscribe("foo",
		func(m *nats.Msg) {
			fmt.Println("received message ", string(m.Data))
		},
	)
	defer sub.Unsubscribe()

	if subScribeError != nil {
		panic(subScribeError)
	}

	if publishError := sc.Publish("foo", []byte("hello world")); publishError != nil {
		panic(publishError)
	} else {
		fmt.Println("sent a message")
	}

	// wait for the message to propagate
	time.Sleep(1 * time.Second)

}
