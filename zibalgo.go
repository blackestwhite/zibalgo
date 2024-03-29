package zibalgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// NewClient creates a new ZibalClient with the provided configuration.
func NewClient(merchant string) *ZibalClient {
	// Initialize the HTTP client
	httpClient := &http.Client{}

	return &ZibalClient{
		httpClient: httpClient,
		merchant:   merchant,
	}
}

func (c *ZibalClient) NewPayment(ctx context.Context, paymentRequest PaymentRequest) (paymentResponse PaymentResponse, err error) {
	paymentRequest.Merchant = c.merchant
	requestBody, err := json.Marshal(paymentRequest)
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("failed to marshal request data: %w", err)
	}

	url := fmt.Sprint(BaseURL, "v1/request")
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&paymentResponse)
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return paymentResponse, nil
}

func (c *ZibalClient) VerifyPayment(ctx context.Context, vericationRequest VerificationRequest) (verificationResponse VerificationResponse, err error) {
	vericationRequest.Merchant = c.merchant
	requestBody, err := json.Marshal(vericationRequest)
	if err != nil {
		return VerificationResponse{}, fmt.Errorf("failed to marshal request data: %w", err)
	}

	url := fmt.Sprint(BaseURL, "v1/verify")
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return VerificationResponse{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return VerificationResponse{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&verificationResponse)
	if err != nil {
		return VerificationResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return verificationResponse, nil
}
