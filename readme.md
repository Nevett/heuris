<div align="center">
    <img src="https://heuris.io/assets/img/logo.png" alt="Heuris" width="200">
    <h1>Heuris</h1>
</div>

[![Build Status](https://travis-ci.org/nullseed/heuris.svg?branch=master)](https://travis-ci.org/nullseed/heuris)
[![Go Report Card](https://goreportcard.com/badge/github.com/nullseed/heuris)](https://goreportcard.com/report/github.com/nullseed/heuris)

Heuris is a HTTP Pub/Sub service that uses POST requests to publish messages and WebSockets to subscribe to them. Heuris doesn't require any configuration and so can be used easily as part of a docker service e.g. when using docker-compose.

The simplest way to run Heuris is by using docker:

```
docker run -p "8080:8080" nullseed/heuris
```

You can then subscribe to a channel using:

```
npm install -g ws
node client.js foo
```

And publish using:

```
curl --data "{}" http://localhost:8080/foo
```

## Monitoring

Heuris has built in monitoring that is available on `http://localhost:8080`. The
front end uses WebSockets to update the page in real-time. Event messages are sent
on channel `_` which means that this channel must not be used in your application. All other channels are free to use.

## Development

Heuris can be built and developed locally. [Go 1.9+](https://tip.golang.org/doc/go1.9) and [dep](https://github.com/golang/dep) are required. From the root directory run:

```
dep ensure
go build && ./heuris
```

To build the container using docker, from the root directory, run:

```
./build.sh
```
