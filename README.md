# PubSub

## The idea

Allowing independent sections of code to communicate with each other through channels without needing to import each other's code. 

## Requirements

* Code should only have to import the pubsub module





## Channels

* local channel: this is a wrapper for golang channels
* remote channel (tcp): this is a tcp/ip wrapper
* remote channel (udp): an udp wrapper


Registry:

A channel is registered to a specific path, ie "/FNode0/state/missing_messages", which is considered the base path. The factory can be retrieved via this path. Every new subscriber is assigned an internal id staring with 0 and has a dedicated path to use in logging, ie "/FNode0/state/missing_messages/0". 