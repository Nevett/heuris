package main

import "time"

type HomeModel struct {
	Channels []ChannelModel
}

type ChannelModel struct {
	Name                 string    `json:"name"`
	NumSubscribers       uint      `json:"numSubscribers"`
	NumMessagesPublished uint      `json:"numMessagesPublished"`
	LastPublished        time.Time `json:"lastPublished"`
}

type ByLastPublished []ChannelModel

func (a ByLastPublished) Len() int {
	return len(a)
}

func (a ByLastPublished) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByLastPublished) Less(i, j int) bool {
	return a[i].LastPublished.Before(a[j].LastPublished)
}
