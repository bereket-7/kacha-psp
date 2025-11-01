package kacha

import (
	"fmt"
	"net/http"
)

// ValidateTransfer validates a B2C transfer before execution
// This endpoint checks:
// - If the account number (phone number) is valid
// - If there is sufficient funds in the Business account
// Returns customer information and transfer details with status 'PREPARED' if successful
func (c *Client) ValidateTransfer(req TransferRequest) (*TransferValidateResponse, error) {
	var response TransferValidateResponse
	var errorResp ErrorResponse

	resp, err := c.httpClient.R().
		SetBody(req).
		SetResult(&response).
		SetError(&errorResp).
		Post(TransferValidateEndpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to validate transfer: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		if errorResp.Message != "" {
			return nil, fmt.Errorf("transfer validation failed: %s (code: %s)", errorResp.Message, errorResp.Code)
		}
		return nil, fmt.Errorf("transfer validation failed with status code: %d", resp.StatusCode())
	}

	return &response, nil
}

// Transfer executes a B2C transfer to a customer account
// This endpoint initiates the actual transfer after validation
func (c *Client) Transfer(req TransferRequest) (*TransferResponse, error) {
	var response TransferResponse
	var errorResp ErrorResponse

	resp, err := c.httpClient.R().
		SetBody(req).
		SetResult(&response).
		SetError(&errorResp).
		Post(TransferEndpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to execute transfer: %w", err)
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		if errorResp.Message != "" {
			return nil, fmt.Errorf("transfer failed: %s (code: %s)", errorResp.Message, errorResp.Code)
		}
		return nil, fmt.Errorf("transfer failed with status code: %d", resp.StatusCode())
	}

	return &response, nil
}

