# Go UDS

Domain Socket example using Golang

Helpful topics glossed over in this page:

1. [have go installed](https://go.dev/doc/install)
2. Be familiar with git - I'm not going to walk through `git clone ..` here, but being on Github, you probably know how to do that. (There's probably a green `<> Code` button above).
3. Use [tmux](https://github.com/tmux/tmux/wiki) to manage terminal windows to run the server and the client at the same time.

# Run the server

```bash
cd domain_sockets/server
go run .
```

# Run the client

In another terminal window..

```bash
cd domain_sockets/client
go run .
```

Both applications point to a socket at `/tmp/my_uds_socket.sock` (with prefix `C:` in Windows) for communication between server and client.

The server application stays running and listens for new requests. The client makes a request, waits for a response to print to the screen, then quits.

