package localchannel

import (
	"sync"

	pubsub "github.com/WhoSoup/factom-pubsub"
)

type channel struct {
	channel chan interface{}

	reader *reader
	writer *writer

	close  sync.Once
	closed bool
	mtx    sync.RWMutex
}

var _ pubsub.IChannel = (*channel)(nil)

func New(size int) *channel {
	lc := new(channel)
	lc.channel = make(chan interface{}, size)

	lc.reader = new(reader)
	lc.reader.c = lc.channel

	lc.writer = new(writer)
	lc.writer.c = lc.channel
	return lc
}

func (c *channel) GetWriter() pubsub.IChannelWriter {
	return c.writer
}

func (c *channel) GetReader() pubsub.IChannelReader {
	return c.reader
}

func (c *channel) Close() {
	c.close.Do(func() {
		c.mtx.Lock()
		c.closed = true
		close(c.channel)
		c.mtx.Unlock()
	})
}

func (c *channel) IsClosed() bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.closed
}

func (c *channel) WriteCount() int64 {
	return c.writer.count
}
