
# TCP-Server (Go)

Overview
--------

This repository contains a minimal TCP client and server written in Go. It was created as a hands-on learning exercise to understand how raw TCP networking works not as production-ready software. The goal is to explore connection handling, byte-level I/O, partial reads/writes, EOF handling, and how Go's goroutines simplify concurrent connection handling.

Why I built this
----------------

- To learn the fundamentals of TCP (beyond HTTP and frameworks).
- To observe real network behaviour: partial reads, buffering, EOF, and connection lifecycle.
- To practice handling concurrent clients with goroutines and safe resource cleanup.
- To provide a small, extendable example others can read, run, and modify.

Repository layout
-----------------

- `server/` — TCP server program (server/main.go)
- `client/` — TCP client program (client/main.go)

Each folder is a standalone Go program. The server listens on a port you provide; the client connects to the configured host and port and exchanges a short message.

Prerequisites
-------------

- Go 1.20+ installed and on your PATH
- Basic familiarity with the command line

Verify Go is installed:

```bash
go version
```

How it works (detailed)
----------------------

- Server: listens on a TCP port, accepts incoming connections, and spins up a goroutine per connection to handle reading all bytes until EOF, then writes the same data back (echo).
- Client: connects to the server, writes a short message, calls `CloseWrite()` to signal end-of-stream, then reads until EOF and prints the echoed response.

Key learning points you will see while running the code
------------------------------------------------------

- TCP is a stream protocol — a single logical message may be received across multiple `Read` calls.
- You must implement framing or message boundaries in real applications (length-prefix, delimiters, etc.).
- `io.EOF` is the signal for clean remote close; handle it explicitly to avoid infinite loops.
- Goroutines make concurrency simple, but you still need to manage connection lifetimes and errors.

Running the example
-------------------

1. Start the server first (pick a port):

```bash
cd server
go run . 3000
```

The server prints a brief message when a client connects. It accepts connections on the specified port.

2. In another terminal, run the client:

```bash
cd client
go run .
```

The client in this repo connects to `localhost:3000` by default, sends a sample message, then reads and prints the echoed response.

Example client output:

```
Reading data...
Reading data...
Received message: Hello from Pratham
```

Notes on the code
-----------------

- Server implementation: `server/main.go`
	- Expects a port argument (e.g. `3000`). If missing, it prints a usage message.
	- Uses `net.Listen("tcp4", ":<port>")` and `Accept()` in a loop.
	- Spawns `hendleConnection(c)` (note: current filename uses that name) as a goroutine to read all bytes and echo them back.

- Client implementation: `client/main.go`
	- Connects to `localhost:3000` (constants at top of file).
	- Writes a test message, calls `CloseWrite()` and then reads until `io.EOF`.

Common gotchas and troubleshooting
---------------------------------

- If you run the client before the server you'll get `connection refused` — always start the server first.
- The server currently collects reads into a slice using `append(packet, tmp...)`. That works for demo purposes but may allocate extra memory; consider reading into a buffer and tracking length in production.
- The server reads until the client signals EOF (client calls `CloseWrite()`), so the client must close its write side after sending.
- If you see empty or padded data in the response, it's likely because the code appends full-sized temporary buffers; trimming by the number of bytes read fixes that.


Testing locally (quick)
-----------------------

Open two terminals. In one run the server with a chosen port. In the other run the client. Observe the logs and echoed message.

Contributing and licensing
-------------------------

This is a personal learning project. Feel free to fork, experiment, and open a PR or issue if you add useful improvements.
