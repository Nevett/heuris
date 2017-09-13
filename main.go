// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func homeGetHandler(instance *Instance, w http.ResponseWriter, r *http.Request) {
	if websocket.IsWebSocketUpgrade(r) {
		subscribeHandler(instance, w, r, "_")
	} else {
		tmpl, err := template.ParseFiles("templates/home.html")
		if err != nil {
			panic(err)
		}

		var model HomeModel
		for _, channel := range instance.Channels() {
			model.Channels = append(model.Channels, ChannelModel{
				channel.name,
				uint(len(channel.clients)),
				channel.numMessagesPublished,
				channel.lastPublish})
		}

		sort.Sort(sort.Reverse(ByLastPublished(model.Channels)))

		err = tmpl.Execute(w, model)
		if err != nil {
			panic(err)
		}
	}
}

func channelGetHandler(instance *Instance, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	channel := params["channel"]

	if websocket.IsWebSocketUpgrade(r) {
		subscribeHandler(instance, w, r, channel)
	} else {
		w.WriteHeader(404)
	}
}

func channelPostHandler(instance *Instance, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	channel := params["channel"]

	publishHandler(instance, w, r, channel)
}

func publishHandler(instance *Instance, w http.ResponseWriter, r *http.Request, channel string) {
	messageType := getMessageType(r)

	body, _ := ioutil.ReadAll(r.Body)
	instance.broadcast <- Message{channel, messageType, body}

	w.WriteHeader(200)
}

func getMessageType(r *http.Request) int {
	contentType := r.Header.Get("Content-Type")

	if contentType == "application/octet-stream" {
		return MessageTypeBinary
	}

	return MessageTypeText
}

func subscribeHandler(instance *Instance, w http.ResponseWriter, r *http.Request, channel string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{instance, conn, make(chan Message, 256), channel}
	client.instance.register <- client

	go client.writeLoop()
	go client.readLoop()
}

func main() {
	addr := flag.String("addr", ":8080", "http service address")
	flag.Parse()

	instance := newInstance()
	go instance.run()

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homeGetHandler(instance, w, r)
	}).Methods("GET")

	router.HandleFunc("/{channel}", func(w http.ResponseWriter, r *http.Request) {
		channelGetHandler(instance, w, r)
	}).Methods("GET")

	router.HandleFunc("/{channel}", func(w http.ResponseWriter, r *http.Request) {
		channelPostHandler(instance, w, r)
	}).Methods("POST")

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",
		http.FileServer(http.Dir("public/assets"))))

	log.Fatal(http.ListenAndServe(*addr, router))
}
