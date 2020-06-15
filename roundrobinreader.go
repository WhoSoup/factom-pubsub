package pubsub

import "sync"

// order not guaranteed
type RoundRobinReader struct {
	channel IChannel

	listeners []chan interface{}
	pos       int

	close chan interface{}
	once  sync.Once
}

func NewRoundRobinReader(channel IChannel, count int) *RoundRobinReader {
	rr := new(RoundRobinReader)
	rr.channel = channel
	rr.listeners = make([]chan interface{}, count)
	for i := range rr.listeners {
		rr.listeners[i] = make(chan interface{})
	}
	rr.close = make(chan interface{})

	go rr.listen()
	return rr
}

func (rr *RoundRobinReader) GetListener(i int) <-chan interface{} {
	return rr.listeners[i]
}

func (rr *RoundRobinReader) Close() {
	rr.once.Do(func() {
		close(rr.close)
		for i := range rr.listeners {
			close(rr.listeners[i])
		}
	})
}

func (rr *RoundRobinReader) listen() {
	defer func() {
		// can happen when rr is closed and it writes
		// to a closed channel
		recover()
	}()

	reader := rr.channel.GetReader().Channel()
	for {
		select {
		case <-rr.close:
			return
		case v, ok := <-reader:
			if ok {
				rr.listeners[rr.pos] <- v
				rr.pos = (rr.pos + 1) % len(rr.listeners)
			} else {
				rr.Close()
			}
		}
	}
}
