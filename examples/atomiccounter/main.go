package main

import (
	"fmt"
	"time"

	pubsub "github.com/WhoSoup/factom-pubsub"
	"github.com/WhoSoup/factom-pubsub/localchannel"
)

// Atomic counter keeps all subscribers on the same level

func main() {
	channel := localchannel.New(0)
	pubsub.GlobalRegistry().Register("/source", channel)

	reader, _ := pubsub.GlobalRegistry().Get("/source")
	multireader := pubsub.NewMultiReader(reader)

	for i := 0; i < 5; i++ {
		worker := i
		multireader.NewCallback(func(o interface{}) {
			if i, ok := o.(int64); ok {
				fmt.Printf("\t%d updated to %d\n", worker, i)
			}
		})
	}

	var i int64
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("Writing %d\n", i)
		channel.GetWriter().Write(i)
		i++
	}
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}
