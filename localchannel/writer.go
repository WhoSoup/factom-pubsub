package localchannel

import (
	"sync/atomic"

	pubsub "github.com/WhoSoup/factom-pubsub"
)

type writer struct {
	c     chan interface{}
	count int64
}

var _ pubsub.IChannelWriter = (*writer)(nil)

func (w *writer) Write(v interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = pubsub.NewChannelIsClosedError()
		}
	}()

	w.c <- v
	atomic.AddInt64(&w.count, 1)
	return nil
}

func (w *writer) Count() int64 {
	return w.count
}
