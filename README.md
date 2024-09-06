# Example API in Go

This repo contains a simple Go app that creates an API that reponds with facts about desert tortoises.

There is no database, so changes will not persist between executions.

To start the API:

```bash
git clone git@github.com:ngrok-samples/example-api-go.git
cd example-api-go
go run main.go
```

## Usage

Get a single random fact:

```bash
curl \
  -X GET \
  http://localhost:5000/random
```

Get all facts:

```bash
curl \
  -X GET \
  http://localhost:5000/facts
```

Get a specific fact:

```bash
curl \
  -X GET \
  https://localhost:5000/fact?id=DT001
```

Add a new fact:

```bash
curl \
  -X POST \ 
  -H "Content-Type: application/json" \
  -d '{"fact": "This is a fact."}' \    
  http://localhost:5000/add
```
