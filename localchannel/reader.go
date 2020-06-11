package localchannel

import pubsub "github.com/WhoSoup/factom-pubsub"

type reader struct {
	c chan interface{}
}

var _ pubsub.IChannelReader = (*reader)(nil)

func (r *reader) Channel() <-chan interface{} {
	return r.c
}
func (r *reader) Read() (interface{}, bool) {
	v, ok := <-r.c
	return v, ok
}
