package main

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/lintopaul/ticket/mocks"
	"github.com/lintopaul/ticket/proto"
)

func TestTicketPurchase(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockClient := mocks.NewMockTicketServiceClient(ctrl)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Set up the mock expectation
    mockClient.EXPECT().PurchaseTicket(gomock.Any(), gomock.Any()).Return(&proto.PurchaseResponse{
        ReceiptId:  "12345",
        From:       "London",
        To:         "France",
        User:       &proto.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
        PricePaid:  20,
        SeatNumber: "1A",
        Section:    "A",
    }, nil)

    // Test PurchaseTicket
    resp, err := mockClient.PurchaseTicket(ctx, &proto.PurchaseRequest{
        User: &proto.User{
            FirstName: "John",
            LastName:  "Doe",
            Email:     "john.doe@example.com",
        },
    })

    if err != nil {
        t.Fatalf("PurchaseTicket failed: %v", err)
    }

    if resp.PricePaid != 20 {
        t.Errorf("Expected price 20, got %v", resp.PricePaid)
    }

    if resp.Section != "A" {
        t.Errorf("Expected section A, got %v", resp.Section)
    }

    if resp.SeatNumber != "1A" {
        t.Errorf("Expected seat number 1A, got %v", resp.SeatNumber)
    }
}
