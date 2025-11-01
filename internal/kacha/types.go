package kacha

// PaymentRequest represents the request body for initiating a payment request
type PaymentRequest struct {
	Phone      string `json:"phone" validate:"required"`
	Amount     int    `json:"amount" validate:"required,min=1"`
	TraceNumber string `json:"trace_number" validate:"required"`
	Reason     string `json:"reason" validate:"required"`
}

// PaymentAuthorizeRequest represents the request body for authorizing a payment
type PaymentAuthorizeRequest struct {
	Reference string `json:"reference" validate:"required"`
	OTP       int    `json:"otp" validate:"required"`
}

// PushUSSDRequest represents the request body for Push USSD payment
type PushUSSDRequest struct {
	Phone       string `json:"phone" validate:"required"`
	Amount      int    `json:"amount" validate:"required,min=1"`
	TraceNumber string `json:"trace_number" validate:"required"`
	CallbackURL string `json:"callback_url" validate:"required,url"`
}

// PaymentRequestResponse represents the response from payment request endpoint
type PaymentRequestResponse struct {
	Success   bool   `json:"success,omitempty"`
	Reference string `json:"reference,omitempty"`
	Message   string `json:"message,omitempty"`
	Status    string `json:"status,omitempty"`
	TraceNumber string `json:"trace_number,omitempty"`
}

// PaymentAuthorizeResponse represents the response from payment authorize endpoint
type PaymentAuthorizeResponse struct {
	Success   bool   `json:"success,omitempty"`
	Message   string `json:"message,omitempty"`
	Status    string `json:"status,omitempty"`
	Reference string `json:"reference,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
}

// PushUSSDResponse represents the response from Push USSD endpoint
type PushUSSDResponse struct {
	Success   bool   `json:"success,omitempty"`
	Message   string `json:"message,omitempty"`
	Status    string `json:"status,omitempty"`
	TraceNumber string `json:"trace_number,omitempty"`
}

// CallbackNotification represents the callback notification from Kacha
type CallbackNotification struct {
	Success      bool   `json:"success,omitempty"`
	Status       string `json:"status,omitempty"`
	TraceNumber  string `json:"trace_number,omitempty"`
	Reference    string `json:"reference,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	Amount       int    `json:"amount,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Message      string `json:"message,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
}

// TransferRequest represents the request body for transfer operations (B2C)
type TransferRequest struct {
	To        string `json:"to" validate:"required"`
	Amount    int    `json:"amount" validate:"required,min=1"`
	Reason    string `json:"reason" validate:"required"`
	ShortCode string `json:"short_code" validate:"required"`
}

// TransferValidateResponse represents the response from transfer validate endpoint
type TransferValidateResponse struct {
	Success   bool   `json:"success,omitempty"`
	Status    string `json:"status,omitempty"`
	Message   string `json:"message,omitempty"`
	To        string `json:"to,omitempty"`
	Amount    int    `json:"amount,omitempty"`
	Reason    string `json:"reason,omitempty"`
	ShortCode string `json:"short_code,omitempty"`
	CustomerInfo *CustomerInfo `json:"customer_info,omitempty"`
}

// CustomerInfo represents customer information returned from validation
type CustomerInfo struct {
	Phone     string `json:"phone,omitempty"`
	Name      string `json:"name,omitempty"`
	AccountID string `json:"account_id,omitempty"`
}

// TransferResponse represents the response from transfer endpoint
type TransferResponse struct {
	Success       bool   `json:"success,omitempty"`
	Status        string `json:"status,omitempty"`
	Message       string `json:"message,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	To            string `json:"to,omitempty"`
	Amount        int    `json:"amount,omitempty"`
	Reference     string `json:"reference,omitempty"`
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Code    string `json:"code,omitempty"`
}

