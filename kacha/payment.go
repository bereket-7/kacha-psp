package kacha

import (
	"fmt"
	"log"
	"net/http"
)

// RequestPayment initiates a payment request using OTP-based authentication
func (c *Client) RequestPayment(req PaymentRequest) (*PaymentRequestResponse, error) {
	var response PaymentRequestResponse
	var errorResp ErrorResponse

	log.Printf("[Kacha] RequestPayment -> endpoint=%s payload=%+v", PaymentRequestEndpoint, req)

	resp, err := c.httpClient.R().
		SetBody(req).
		SetResult(&response).
		SetError(&errorResp).
		Post(PaymentRequestEndpoint)

	if err != nil {
		log.Printf("[Kacha] RequestPayment error: %v", err)
		return nil, fmt.Errorf("failed to make payment request: %w", err)
	}

	log.Printf("[Kacha] RequestPayment <- status=%d response=%+v errorResponse=%+v",
		resp.StatusCode(), response, errorResp)

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		if errorResp.Error != nil {
			return nil, fmt.Errorf("payment request failed: %s (status_code: %s, detail: %s)",
				errorResp.Error.Message,
				errorResp.Error.StatusCode,
				errorResp.Error.Detail,
			)
		}
		return nil, fmt.Errorf("payment request failed with status code: %d", resp.StatusCode())
	}

	return &response, nil
}

// AuthorizePayment authorizes a payment request using reference and OTP
func (c *Client) AuthorizePayment(req PaymentAuthorizeRequest) (*PaymentAuthorizeResponse, error) {
	var response PaymentAuthorizeResponse
	var errorResp ErrorResponse

	log.Printf("[Kacha] AuthorizePayment -> endpoint=%s payload=%+v", PaymentAuthorizeEndpoint, req)

	resp, err := c.httpClient.R().
		SetBody(req).
		SetResult(&response).
		SetError(&errorResp).
		Post(PaymentAuthorizeEndpoint)

	if err != nil {
		log.Printf("[Kacha] AuthorizePayment error: %v", err)
		return nil, fmt.Errorf("failed to authorize payment: %w", err)
	}

	log.Printf("[Kacha] AuthorizePayment <- status=%d response=%+v errorResponse=%+v",
		resp.StatusCode(), response, errorResp)

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		if errorResp.Error != nil {
			return nil, fmt.Errorf("payment authorization failed: %s (status_code: %s, detail: %s)",
				errorResp.Error.Message,
				errorResp.Error.StatusCode,
				errorResp.Error.Detail,
			)
		}
		return nil, fmt.Errorf("payment authorization failed with status code: %d", resp.StatusCode())
	}

	return &response, nil
}

// RequestPushUSSD initiates a Push USSD payment request
func (c *Client) RequestPushUSSD(req PushUSSDRequest) (*PushUSSDResponse, error) {
	var response PushUSSDResponse
	var errorResp ErrorResponse

	log.Printf("[Kacha] RequestPushUSSD -> endpoint=%s payload=%+v", PushUSSDEndpoint, req)

	resp, err := c.httpClient.R().
		SetBody(req).
		SetResult(&response).
		SetError(&errorResp).
		Post(PushUSSDEndpoint)

	if err != nil {
		log.Printf("[Kacha] RequestPushUSSD error: %v", err)
		return nil, fmt.Errorf("failed to initiate push USSD payment: %w", err)
	}

	log.Printf("[Kacha] RequestPushUSSD <- status=%d response=%+v errorResponse=%+v",
		resp.StatusCode(), response, errorResp)

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		if errorResp.Error != nil {
			return nil, fmt.Errorf("push USSD payment failed: %s (status_code: %s, detail: %s)",
				errorResp.Error.Message,
				errorResp.Error.StatusCode,
				errorResp.Error.Detail,
			)
		}
		return nil, fmt.Errorf("push USSD payment failed with status code: %d", resp.StatusCode())
	}

	return &response, nil
}
