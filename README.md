# MarxquoteAPI

A tiny HTTP API that returns a random Karl Marx quote, rewritten in Go.

## Run

```sh
go run .
```

Then open http://localhost:3000. The port can be changed with the `PORT` environment variable.

## Endpoints

- `GET /` — plain text in the form `<quote text> - <author>`
- `GET /getquotejson` — the full quote object as JSON

## Build

With [just](https://github.com/casey/just):

```sh
just build v1
```

Or directly:

```sh
go build -o main .
```

## Docker

```sh
docker build -t marxquoteapi .
docker run -p 3000:3000 marxquoteapi
```
