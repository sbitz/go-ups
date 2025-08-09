# Domain Sockets

Run the server

```bash
cd domainSockets/server
go run .
```

Run the client

```bash
cd domainSockets/client
go run .
```

Both need to define the socket located at `C:/tmp/my_dms_socket.sock` which is used for communication.

Using protocol buffers, processes can communicate, eliminating the need for monolith applications in Go. 

# What now?

1. Define a client module to include as a dependency to talk with another service.
2. How do you define a protocol buffer in golang to send objects across a socket?
3. Run golang applications on K8s?

# What next?

1. Feature flags - manage active features
2. Service discovery - registry of active services/processes (self-registration?)
