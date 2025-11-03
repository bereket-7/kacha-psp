package kacha

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	// DefaultBaseURL is the default base URL for Kacha API
	DefaultBaseURL = "https://docs.kacha.net/v1"
	// PaymentRequestEndpoint is the endpoint for payment request
	PaymentRequestEndpoint = "/api/v1/orgs/payment/request"
	// PaymentAuthorizeEndpoint is the endpoint for payment authorization
	PaymentAuthorizeEndpoint = "/api/v1/orgs/payment/authorize"
	// PushUSSDEndpoint is the endpoint for Push USSD payment
	PushUSSDEndpoint = "/api/v1/orgs/payment/request/push_ussd"
	// TransferValidateEndpoint is the endpoint for validating transfers (B2C)
	TransferValidateEndpoint = "/api/v1/orgs/transfer/validate"
	// TransferEndpoint is the endpoint for executing transfers (B2C)
	TransferEndpoint = "/api/v1/orgs/transfer"
)

// Client represents a Kacha API client
type Client struct {
    username   string
    password   string
	baseURL    string
	httpClient *resty.Client
}

// NewClient creates a new Kacha API client
func NewClient(username, password string) *Client {
    return NewClientWithBaseURL(username, password, DefaultBaseURL)
}

// NewClientWithBaseURL creates a new Kacha API client with a custom base URL
func NewClientWithBaseURL(username, password, baseURL string) *Client {
	client := resty.New()
	
	// Set base URL
	client.SetBaseURL(baseURL)
	
	// Set Basic Auth header
    auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	client.SetHeader("Authorization", fmt.Sprintf("Basic %s", auth))
	
	// Set default headers
	client.SetHeader("Content-Type", "application/json")
	client.SetHeader("Accept", "application/json")
	
	// Enable debug mode for development (can be disabled in production)
	// client.SetDebug(true)
	
	return &Client{
        username:   username,
        password:   password,
		baseURL:    baseURL,
		httpClient: client,
	}
}

// SetDebug enables or disables debug mode
func (c *Client) SetDebug(debug bool) {
	c.httpClient.SetDebug(debug)
}

// SetTimeout sets the HTTP client timeout in seconds
func (c *Client) SetTimeout(timeoutSeconds int) {
	c.httpClient.SetTimeout(time.Duration(timeoutSeconds) * time.Second)
}

