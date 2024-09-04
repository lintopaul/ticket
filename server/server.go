package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"sync"

	"github.com/lintopaul/ticket/proto"
	"google.golang.org/grpc"
)

const (
    TICKET_PRICE = 20
    SECTION_CAPACITY = 10
)

// server struct to implement the gRPC service.
type server struct {
    proto.UnimplementedTicketServiceServer
    mu        sync.Mutex
    users     map[string]*proto.PurchaseResponse
    sectionA  int
    sectionB  int
}

// newServer initializes a new Server with default values.
func newServer() *server {
    return &server{
        users:    make(map[string]*proto.PurchaseResponse),
        sectionA: 1,
        sectionB: 1,
    }
}

// PurchaseTicket handles the purchase of a train ticket, assigns a seat to the user,
// and returns a receipt with details of the purchase.
func (s *server) PurchaseTicket(ctx context.Context, req *proto.PurchaseRequest) (*proto.PurchaseResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Allocate a seat to the user
    section := "A"
    seatNumber := s.sectionA
    if s.sectionA > SECTION_CAPACITY {
        section = "B"
        seatNumber = s.sectionB
        s.sectionB++
    } else {
        s.sectionA++
    }

    // Generate a receipt with the ticket details
    receipt := &proto.PurchaseResponse{
        ReceiptId:  strconv.Itoa(rand.Intn(100000)),
        From:       "London",
        To:         "France",
        User:       req.User,
        PricePaid:  TICKET_PRICE,
        SeatNumber: fmt.Sprintf("%d", seatNumber),
        Section:    section,
    }

    // Store the user's seat information
    s.users[req.User.Email] = receipt

    log.Printf("User %s purchased a ticket: %+v\n", req.User.Email, receipt)
    return receipt, nil
}

// GetReceipt returns the receipt with ticket details
func (s *server) GetReceipt(ctx context.Context, req *proto.ReceiptRequest) (*proto.ReceiptResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    if receipt, ok := s.users[req.Email]; ok {
        return &proto.ReceiptResponse{
            ReceiptId:  receipt.ReceiptId,
            From:       receipt.From,
            To:         receipt.To,
            User:       receipt.User,
            PricePaid:  receipt.PricePaid,
            SeatNumber: receipt.SeatNumber,
            Section:    receipt.Section,
        }, nil
    }

    return nil, fmt.Errorf("no receipt found for user with email %s", req.Email)
}

// GetUsersBySection returns the seat allocation details for a specific section.
func (s *server) GetUsersBySection(ctx context.Context, req *proto.SectionRequest) (*proto.SectionResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    var users []*proto.SeatAllocation
    for _, receipt := range s.users {
        if receipt.Section == req.Section {
            users = append(users, &proto.SeatAllocation{
                User:       receipt.User,
                SeatNumber: receipt.SeatNumber,
            })
        }
    }

    return &proto.SectionResponse{Users: users}, nil
}

// RemoveUser removes a user from the train by their email, freeing up their seat.
func (s *server) RemoveUser(ctx context.Context, req *proto.RemoveUserRequest) (*proto.RemoveUserResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, ok := s.users[req.Email]; ok {
        delete(s.users, req.Email)
        log.Printf("User %s removed from the train.\n", req.Email)
        return &proto.RemoveUserResponse{Success: true}, nil
    }

    return &proto.RemoveUserResponse{Success: false}, fmt.Errorf("user with email %s not found", req.Email)
}

// ModifySeat allows a user to change their seat to a different section or seat number.
func (s *server) ModifySeat(ctx context.Context, req *proto.ModifySeatRequest) (*proto.ModifySeatResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    if receipt, ok := s.users[req.Email]; ok {
        receipt.SeatNumber = req.NewSeatNumber
        log.Printf("Seat for user %s changed to %s.\n", req.Email, req.NewSeatNumber)
        return &proto.ModifySeatResponse{Success: true}, nil
    }

    return &proto.ModifySeatResponse{Success: false}, fmt.Errorf("user with email %s not found", req.Email)
}

func main() {
    // Listen on a specific port for incoming gRPC requests
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    // Create a new gRPC server
    s := grpc.NewServer()

    // Register the TicketService with the gRPC server
    proto.RegisterTicketServiceServer(s, newServer())

    log.Println("gRPC server is running on port 50051...")

    // Start the server and listen for incoming connections
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
