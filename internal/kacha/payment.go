package kacha

import (
	"fmt"
	"net/http"
)

// RequestPayment initiates a payment request using OTP-based authentication
// This sends an OTP to the customer via SMS and returns payment request details
func (c *Client) RequestPayment(req PaymentRequest) (*PaymentRequestResponse, error) {
	var response PaymentRequestResponse
	var errorResp ErrorResponse

	resp, err := c.httpClient.R().
		SetBody(req).
		SetResult(&response).
		SetError(&errorResp).
		Post(PaymentRequestEndpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to make payment request: %w", err)
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		if errorResp.Message != "" {
			return nil, fmt.Errorf("payment request failed: %s (code: %s)", errorResp.Message, errorResp.Code)
		}
		return nil, fmt.Errorf("payment request failed with status code: %d", resp.StatusCode())
	}

	return &response, nil
}

// AuthorizePayment authorizes a payment request using reference and OTP
// This endpoint approves a payment request that was initiated via RequestPayment
func (c *Client) AuthorizePayment(req PaymentAuthorizeRequest) (*PaymentAuthorizeResponse, error) {
	var response PaymentAuthorizeResponse
	var errorResp ErrorResponse

	resp, err := c.httpClient.R().
		SetBody(req).
		SetResult(&response).
		SetError(&errorResp).
		Post(PaymentAuthorizeEndpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to authorize payment: %w", err)
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		if errorResp.Message != "" {
			return nil, fmt.Errorf("payment authorization failed: %s (code: %s)", errorResp.Message, errorResp.Code)
		}
		return nil, fmt.Errorf("payment authorization failed with status code: %d", resp.StatusCode())
	}

	return &response, nil
}

// RequestPushUSSD initiates a Push USSD payment request
// This pushes a USSD prompt to the customer's mobile phone
// The transaction result will be sent asynchronously to the callback_url
func (c *Client) RequestPushUSSD(req PushUSSDRequest) (*PushUSSDResponse, error) {
	var response PushUSSDResponse
	var errorResp ErrorResponse

	resp, err := c.httpClient.R().
		SetBody(req).
		SetResult(&response).
		SetError(&errorResp).
		Post(PushUSSDEndpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to initiate push USSD payment: %w", err)
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		if errorResp.Message != "" {
			return nil, fmt.Errorf("push USSD payment failed: %s (code: %s)", errorResp.Message, errorResp.Code)
		}
		return nil, fmt.Errorf("push USSD payment failed with status code: %d", resp.StatusCode())
	}

	return &response, nil
}

