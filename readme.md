<div align="center">
    <img src="https://heuris.io/assets/img/logo.png" alt="Heuris" width="200">
    <h1>Heuris</h1>
</div>

[![Build Status](https://travis-ci.org/nullseed/heuris.svg?branch=master)](https://travis-ci.org/nullseed/heuris)
[![Go Report Card](https://goreportcard.com/badge/github.com/nullseed/heuris)](https://goreportcard.com/report/github.com/nullseed/heuris)

## Build

From the root directory, run:

```
go build
```

## Run

From the root directory, run:

```
./heuris
```

## Use

### Subscribing

```
cd tools
npm install -g ws
node client.js
```

### Publishing

```
curl --data "{}" http://localhost:8080/foo
```

## Docker

Run the container:

```
docker run -p "8080:8080" nullseed/heuris
```

Build the container:

```
docker build -t nullseed/heuris:latest .
```
