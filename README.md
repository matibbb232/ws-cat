# Project ws-cat

This is a simple utility for testing WebSockets connections.

## Example usage

```bash
$ ./wscat --url wss://websocket-echo.com
>> Hi!
<< Hi!
>> ^C2025/03/11 16:34:47 Received shutdown signal. Shutting down gracefully...
```

## Building from source
### Prerequisites
 - Golang >= 1.24

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Clean up binary from the last build:
```bash
make clean
```
