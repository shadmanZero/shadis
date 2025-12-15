# Shadis 🔴

> Building Redis from scratch to learn how it works under the hood.

## What is this?

Just me having fun building a Redis clone in Go. It's a side project — **not meant for production use**. Just vibes and learning by doing.

## What works so far

- [x] TCP server on port 6379
- [x] RESP protocol parser
- [x] Command registry
- [x] Basic commands: `PING`, `ECHO`, `GET`, `SET`
- [x] In-memory key-value store
- [x] Structured logging

## What's next

- [ ] More commands: `DEL`, `EXISTS`, `INCR`, `DECR`
- [ ] Key expiration (TTL)
- [ ] Persistence (RDB/AOF)
- [ ] Pub/Sub
- [ ] Transactions

## Running it

```bash
# Clone and enter
cd shadis

# Copy env file
cp .env.example .env

# Run in dev mode
make dev

# Or build and run
make run
```

## Testing

Connect with telnet:

```bash
telnet localhost 6379
```

Then type RESP commands. Example for PING:

```
*1
$4
PING
```

Response:

```
+PONG
```

Example for SET and GET:

```
*3
$3
SET
$4
name
$7
shadman
```

```
+OK
```

```
*2
$3
GET
$4
name
```

```
$7
shadman
```

**Tips:**

- Each line must end with Enter
- See [RESP protocol docs](https://redis.io/docs/reference/protocol-spec/) for more info

## Project Structure

```
.
├── cmd/shadis/          # Entry point
├── internal/
│   ├── command/         # Command implementations
│   ├── config/          # Configuration
│   ├── logger/          # Structured logging
│   ├── resp/            # RESP protocol parser/writer
│   └── store/           # Key-value store
├── .env.example         # Environment template
├── Makefile             # Build commands
└── README.md
```

---

*This is a learning project. If you need a real Redis, use [Redis](https://redis.io).*
