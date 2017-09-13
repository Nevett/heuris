// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"time"
)

const (
	MessageTypeText   = iota
	MessageTypeBinary = iota
)

type Message struct {
	channel     string
	messageType int
	data        []byte
}

type Channel struct {
	name                 string
	clients              []*Client
	numMessagesPublished uint
	lastPublish          time.Time
}

type Instance struct {
	channels   map[string]*Channel
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

func newInstance() *Instance {
	return &Instance{
		channels:   make(map[string]*Channel),
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (inst *Instance) run() {
	for {
		select {
		case client := <-inst.register:
			inst.registerClient(client)
		case client := <-inst.unregister:
			inst.unregisterClient(client)
		case message := <-inst.broadcast:
			inst.broadcastMessage(message)
		}
	}
}

func (inst *Instance) registerClient(client *Client) {
	if _, ok := inst.channels[client.channel]; ok {
		inst.channels[client.channel].clients = append(inst.channels[client.channel].clients, client)
	} else {
		inst.channels[client.channel] = &Channel{client.channel, []*Client{client}, 0, time.Time{}}
	}

	go inst.raiseEvent(client.channel)
}

func (inst *Instance) unregisterClient(client *Client) {
	channel := inst.channels[client.channel]
	for i := range channel.clients {
		if channel.clients[i] == client {
			close(client.send)
			inst.channels[client.channel].clients = append(channel.clients[:i], channel.clients[i+1:]...)
			break
		}
	}

	go inst.raiseEvent(client.channel)
}

func (inst *Instance) broadcastMessage(message Message) {
	channel, ok := inst.channels[message.channel]
	if !ok {
		return
	}
	inst.channels[message.channel].numMessagesPublished++
	inst.channels[message.channel].lastPublish = time.Now()

	for i, client := range channel.clients {
		select {
		case client.send <- message:
		default:
			close(client.send)
			inst.channels[message.channel].clients = append(channel.clients[:i], channel.clients[i+1:]...)
		}
	}

	go inst.raiseEvent(message.channel)
}

func (inst *Instance) Channels() map[string]*Channel {
	channels := make(map[string]*Channel)
	for key, value := range inst.channels {
		if key != "_" {
			channels[key] = value
		}
	}
	return channels
}

func (inst *Instance) raiseEvent(channel string) {
	c := inst.channels[channel]
	if c.name == "_" {
		return
	}

	e := ChannelModel{c.name, uint(len(c.clients)), c.numMessagesPublished, c.lastPublish}
	b, _ := json.Marshal(e)
	inst.broadcast <- Message{"_", MessageTypeText, b}
}
