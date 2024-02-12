package zibalgo

import (
	"context"
	"reflect"
	"testing"
	"time"
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

func TestNewPayment_WithTimeout(t *testing.T) {
	// Create a new ZibalClient instance
	client := NewClient("zibal")

	// Create a PaymentRequest instance
	paymentRequest := PaymentRequest{
		CallbackURL: "https://example.com/callback",
		Description: "Test payment",
		Amount:      10000,
	}

	// Call the NewPayment method with a context
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond)
	defer cancel()

	_, err := client.NewPayment(ctx, paymentRequest)
	if err == nil {
		t.Errorf("expected error but got nil")
	}

	expected := `HTTP request failed: Post "https://gateway.zibal.ir/v1/request": context deadline exceeded`
	if err.Error() != expected {
		t.Errorf("expected timeout error but got: %s", err.Error())
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
	verificationRes, err := client.VerifyPayment(ctx, verificationRequest)
	if err != nil && reflect.DeepEqual(verificationRes, VerificationRequest{}) {
		t.Errorf("VerifyPayment failed: %v", err)
	}
}

func TestVerifyPayment_WithTimeout(t *testing.T) {
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

	dctx, cancel := context.WithTimeout(ctx, time.Millisecond)
	defer cancel()
	// Call the VerifyPayment method with a context
	_, err = client.VerifyPayment(dctx, verificationRequest)
	if err == nil {
		t.Errorf("expected timeout error but got nil")
	}

	expected := `HTTP request failed: Post "https://gateway.zibal.ir/v1/verify": context deadline exceeded`
	if err.Error() != expected {
		t.Errorf("expected timeout error but got: %s", err.Error())
	}
}

func TestVerifyPayment_Result(t *testing.T) {
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
	verificationRes, err := client.VerifyPayment(ctx, verificationRequest)
	if err != nil {
		t.Errorf("VerifyPayment failed: %v", err)
	}

	if verificationRes.Result != NotPaid {
		t.Errorf("VerifyPayment Result failed, expected: %d, got: %d", NotPaid, verificationRes.Result)
	}
}

func TestNewPayment_Result_Success(t *testing.T) {
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

	if res.Result != SuccessCode {
		t.Errorf("expected result code: %d, got: %d", SuccessCode, res.Result)
	}
}

func TestNewPayment_Result_AmountTooSmall(t *testing.T) {
	// Create a new ZibalClient instance
	client := NewClient("zibal")

	// Create a PaymentRequest instance
	paymentRequest := PaymentRequest{
		CallbackURL: "https://example.com/callback",
		Description: "Test payment",
		Amount:      999,
	}

	// Call the NewPayment method with a context
	ctx := context.Background()
	res, err := client.NewPayment(ctx, paymentRequest)
	if err != nil {
		t.Errorf("NewPayment failed: %v", err)
	}

	if res.Result != AmountTooSmall {
		t.Errorf("expected result code: %d, got: %d", AmountTooSmall, res.Result)
	}
}

func TestNewPayment_Result_AmountExeeded(t *testing.T) {
	// Create a new ZibalClient instance
	client := NewClient("zibal")

	// Create a PaymentRequest instance
	paymentRequest := PaymentRequest{
		CallbackURL: "ton://blackestwhite.ton",
		Description: "Test payment",
		Amount:      1000000001,
	}

	// Call the NewPayment method with a context
	ctx := context.Background()
	res, err := client.NewPayment(ctx, paymentRequest)
	if err != nil {
		t.Errorf("NewPayment failed: %v", err)
	}

	if res.Result != AmountExeeded {
		t.Errorf("expected result code: %d, got: %d", AmountExeeded, res.Result)
	}
}
