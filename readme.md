<div align="center">
    <img src="https://heuris.io/assets/img/logo.png" alt="Heuris" width="200">
    <h1>Heuris</h1>
</div>

[![Build Status](https://travis-ci.org/nullseed/heuris.svg?branch=master)](https://travis-ci.org/nullseed/heuris)
[![Go Report Card](https://goreportcard.com/badge/github.com/nullseed/heuris)](https://goreportcard.com/report/github.com/nullseed/heuris)

Heuris is a HTTP Pub/Sub service that uses POST requests to publish messages and WebSockets to subscribe to them.

The simplest way to run Heuris is using docker:

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


## Development

To build heuris, from the root directory, run:

```
./build.sh
```
