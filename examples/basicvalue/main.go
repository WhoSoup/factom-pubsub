package main

import (
	"fmt"
	"log"

	pubsub "github.com/WhoSoup/factom-pubsub"
	"github.com/WhoSoup/factom-pubsub/localchannel"
)

func main() {
	local := localchannel.New(5)

	pubsub.GlobalRegistry().Register("/basic", local)

	go Write()

	Read()
}

func Read() {
	c, ok := pubsub.GlobalRegistry().Get("/basic")
	if !ok {
		log.Fatalln("channel /basic doesn't exist")
	}

	for v := range c.GetReader().Channel() {
		fmt.Printf("<- %+v\n", v)
	}

	fmt.Println("channel was closed")
}

func Write() {
	c, ok := pubsub.GlobalRegistry().Get("/basic")
	if !ok {
		log.Fatalln("channel /basic doesn't exist")
	}

	for i := 0; i < 16; i++ {
		c.GetWriter().Write(i)
	}

	c.Close()
}
