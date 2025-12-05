package bank

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/itua234/payment-bridge/internal/models"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (c *Client) Authorize(
	ctx context.Context,
	payment *models.Payment,
	idempotencyKey string,
) (*AuthorizeResponse, error) {
	reqBody := AuthorizeRequest{
		Amount:      payment.Amount,
		CardNumber:  payment.CardNumber,
		CVV:         payment.CVV,
		ExpiryMonth: payment.ExpiryMonth,
		ExpiryYear:  payment.ExpiryYear,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		c.baseURL+"/api/v1/authorizations",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Idempotency-Key", idempotencyKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("bank error %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var authResp AuthorizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &authResp, nil
}

func (c *Client) Capture(
	ctx context.Context,
	authRef string,
	amount int64,
	idempotencyKey string,
) (*CaptureResponse, error) {
	reqBody := CaptureRequest{
		Amount:          amount,
		AuthorizationID: authRef,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		c.baseURL+"/api/v1/captures",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Idempotency-Key", idempotencyKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("bank error %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var captureResp CaptureResponse
	if err := json.NewDecoder(resp.Body).Decode(&captureResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &captureResp, nil
}

func (c *Client) Void(
	ctx context.Context,
	authRef string,
	idempotencyKey string,
) (*VoidResponse, error) {
	reqBody := VoidRequest{
		AuthorizationID: authRef,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		c.baseURL+"/api/v1/voids",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Idempotency-Key", idempotencyKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("bank error %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var voidResp VoidResponse
	if err := json.NewDecoder(resp.Body).Decode(&voidResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &voidResp, nil
}

func (c *Client) Refund(
	ctx context.Context,
	captureRef string,
	amount int64,
	idempotencyKey string,
) (*RefundResponse, error) {
	reqBody := RefundRequest{
		Amount:    amount,
		CaptureID: captureRef,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		c.baseURL+"/api/v1/refunds",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Idempotency-Key", idempotencyKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("bank error %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var refundResp RefundResponse
	if err := json.NewDecoder(resp.Body).Decode(&refundResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &refundResp, nil
}
