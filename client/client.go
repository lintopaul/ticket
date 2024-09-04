package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/lintopaul/ticket/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Setup the connection using grpc.DialContext, but with proper arguments
    conn, err := grpc.DialContext(
        ctx,
        "localhost:50051", // The target address.
        grpc.WithTransportCredentials(insecure.NewCredentials()), // Use insecure transport credentials.
        grpc.WithBlock(), // Block until the connection is up.
    )
    if err != nil {
        log.Fatalf("failed to dial: %v", err)
    }
    defer conn.Close()

    client := proto.NewTicketServiceClient(conn)

    // Purchase a ticket
    user := &proto.User{
        FirstName: "Linto",
        LastName:  "Paul",
        Email:     "lintopaul@trainticketbooking.com",
    }
    purchaseRes, err := client.PurchaseTicket(context.Background(), &proto.PurchaseRequest{User: user})
    if err != nil {
        log.Fatalf("Could not purchase ticket: %v", err)
    }
    fmt.Printf("Purchase Response: %+v\n", purchaseRes)

    // Get the receipt for the user
    receiptRes, err := client.GetReceipt(context.Background(), &proto.ReceiptRequest{Email: user.Email})
    if err != nil {
        log.Fatalf("Could not get receipt: %v", err)
    }
    fmt.Printf("Receipt Response: %+v\n", receiptRes)

    // Get users by section
    sectionRes, err := client.GetUsersBySection(context.Background(), &proto.SectionRequest{Section: "A"})
    if err != nil {
        log.Fatalf("Could not get users by section: %v", err)
    }
    fmt.Printf("Section A Users: %+v\n", sectionRes.Users)

    // Modify seat
    _, err = client.ModifySeat(context.Background(), &proto.ModifySeatRequest{Email: user.Email, NewSeatNumber: "5"})
    if err != nil {
        log.Fatalf("Could not modify seat: %v", err)
    }
    fmt.Println("Seat modified successfully")

    // Remove user
    _, err = client.RemoveUser(context.Background(), &proto.RemoveUserRequest{Email: user.Email})
    if err != nil {
        log.Fatalf("Could not remove user: %v", err)
    }
    fmt.Println("User removed successfully")
}
