# Ticket
Ticket is a simple in-memory train ticketing system using Golang and gRPC. Users can purchase tickets for a train journey from London to France, view receipts, view users and their seats, remove users, and modify user seats.

## Features

### Purchase Ticket
Allows a user to purchase a train ticket from London to France. Upon purchase, the user is allocated a seat in one of the two sections (A or B).

### View Receipt
Retrieve the details of a purchased ticket, including the user's details, seat number, and section.

### View Seat Allocation
View all users and their allocated seats within a specific section.

### Modify Seat
Change the seat allocation for a user.

### Remove User
Remove a user from the train, freeing up their seat.

## Project Structure

```
├── README.md
├── client
│   ├── client.go
│   └── client_test.go
├── go.mod
├── go.sum
├── mocks
│   └── ticket_service_mock.go
├── proto
│   ├── ticket.pb.go
│   ├── ticket.proto
│   └── ticket_grpc.pb.go
└── server
    └── server.go
```

## Prerequisites
- Golang: Version 1.21 or above
- gRPC: For server and client communication
- Protocol Buffers: For defining the gRPC service

## Setup
1. Clone the repository:
```
git clone https://github.com/lintopaul/ticket.git
cd ticket
```
2. Install dependencies:
Ensure you have Go modules enabled and run:
```
go mod tidy
```
3. Generate gRPC code from Protobuf:
Run the following command to generate the necessary Go code from the .proto file:
```
protoc --go_out=. --go_opt=paths=source_relative \ 
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
proto/ticket.proto
```
4. Run the server:
Start the server by running:
```
go run server/server.go
```
5. Run the client
From another terminal, run the client:
```
go run client/client.go
```

This will set up a complete gRPC service with the server and client implementing the train ticket booking system.

## Run Tests
A mock client for testing is used without requiring an actual gRPC server.

```
go test -v ./...
```

`Note`: Mocks were generated using:
```
mockgen -destination=mocks/ticket_service_mock.go -package=mocks github.com/lintopaul/ticket/proto TicketServiceClient  
```
where TicketServiceClient is the service interface

