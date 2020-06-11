package localchannel

import (
	pubsub "github.com/WhoSoup/factom-pubsub"
)

type writer struct {
	c chan interface{}
}

var _ pubsub.IChannelWriter = (*writer)(nil)

func (w *writer) Write(v interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = pubsub.NewChannelIsClosedError()
		}
	}()

	w.c <- v
	return nil
}
