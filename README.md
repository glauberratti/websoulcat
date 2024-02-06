# WebSoulCat

A simple project to practice `goroutines`, `channels`, `websocket` and more as the project progresses.

The project has a `server` application and a `client` application. The server starts and manages chat rooms, obviously it must be run first. The client connects to a chat room, and multiple clients can be run.

## Server
To run the `server` you can use the following command:

```bash
go run cmd/server/main.go
```

## Client
To run the `client` you can use the following command:

```bash
go run cmd/client/main.go
```

Enter a name and then a code for the room.

Run more clients to chat in the same room (same code room).
