package pubsub

import "sync"

type SubscriptionManager struct {
	mtx     sync.Mutex
	subs    []ISubscriber
	channel IChannel
	stop    chan interface{}
}

func NewSubscriptionManager(channel IChannel) *SubscriptionManager {
	sm := new(SubscriptionManager)
	sm.channel = channel
	sm.stop = make(chan interface{}, 1)
	return sm
}

func (sm *SubscriptionManager) Register() {

}
