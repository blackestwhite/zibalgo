package zibalgo

import (
	"context"
	"testing"
)

func TestNewPayment(t *testing.T) {
	// Create a new ZibalClient instance
	client := NewClient("zibal")

	// Create a PaymentRequest instance
	paymentRequest := PaymentRequest{
		CallbackURL: "https://example.com/callback",
		Description: "Test payment",
		Amount:      10000,
	}

	// Call the NewPayment method with a context
	ctx := context.Background()
	_, err := client.NewPayment(ctx, paymentRequest)
	if err != nil {
		t.Errorf("NewPayment failed: %v", err)
	}
}

func TestVerifyPayment(t *testing.T) {
	// Create a new ZibalClient instance
	client := NewClient("zibal")

	// Create a PaymentRequest instance
	paymentRequest := PaymentRequest{
		CallbackURL: "https://example.com/callback",
		Description: "Test payment",
		Amount:      10000,
	}

	// Call the NewPayment method with a context
	ctx := context.Background()
	res, err := client.NewPayment(ctx, paymentRequest)
	if err != nil {
		t.Errorf("NewPayment failed: %v", err)
	}

	verificationRequest := VerificationRequest{
		TrackID: res.TrackID,
	}
	// Call the VerifyPayment method with a context
	_, err = client.VerifyPayment(ctx, verificationRequest)
	if err != nil {
		t.Errorf("VerifyPayment failed: %v", err)
	}
}
