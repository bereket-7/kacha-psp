package kacha

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	DefaultBaseURL = "https://docs.kacha.net/api/v1"
	PaymentRequestEndpoint     = "/orgs/payment/request"
	PaymentAuthorizeEndpoint   = "/orgs/payment/authorize"
	PushUSSDEndpoint           = "/orgs/payment/request/push_ussd"
	TransferValidateEndpoint   = "/orgs/transfer/validate"
	TransferEndpoint           = "/orgs/transfer"
)

type Client struct {
    username   string
    password   string
	baseURL    string
	httpClient *resty.Client
}

func NewClient(username, password string) *Client {
    return NewClientWithBaseURL(username, password, DefaultBaseURL)
}

func NewClientWithBaseURL(username, password, baseURL string) *Client {
	client := resty.New()
	
	client.SetBaseURL(baseURL)
	
    auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	client.SetHeader("Authorization", fmt.Sprintf("Basic %s", auth))
	
	client.SetHeader("Content-Type", "application/json")
	client.SetHeader("Accept", "application/json")
	
	return &Client{
        username:   username,
        password:   password,
		baseURL:    baseURL,
		httpClient: client,
	}
}

func (c *Client) SetDebug(debug bool) {
	c.httpClient.SetDebug(debug)
}

func (c *Client) SetTimeout(timeoutSeconds int) {
	c.httpClient.SetTimeout(time.Duration(timeoutSeconds) * time.Second)
}

