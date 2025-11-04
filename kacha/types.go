package kacha

type PaymentRequest struct {
	Phone       string `json:"phone" validate:"required"`
	Amount      int    `json:"amount" validate:"required,min=1"`
	TraceNumber string `json:"trace_number" validate:"required"`
	Reason      string `json:"reason" validate:"required"`
}

type PaymentAuthorizeRequest struct {
	Reference string `json:"reference" validate:"required"`
	OTP       int    `json:"otp" validate:"required"`
}

type PushUSSDRequest struct {
	Phone       string `json:"phone" validate:"required"`
	Amount      int    `json:"amount" validate:"required,min=1"`
	TraceNumber string `json:"trace_number" validate:"required"`
	CallbackURL string `json:"callback_url" validate:"required,url"`
}

type PaymentRequestResponse struct {
	Success     bool   `json:"success,omitempty"`
	Reference   string `json:"reference,omitempty"`
	Message     string `json:"message,omitempty"`
	Status      string `json:"status,omitempty"`
	TraceNumber string `json:"trace_number,omitempty"`
}

type PaymentAuthorizeResponse struct {
	Success       bool   `json:"success,omitempty"`
	Message       string `json:"message,omitempty"`
	Status        string `json:"status,omitempty"`
	Reference     string `json:"reference,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
}

type PushUSSDResponse struct {
	Success     bool   `json:"success,omitempty"`
	Message     string `json:"message,omitempty"`
	Status      string `json:"status,omitempty"`
	TraceNumber string `json:"trace_number,omitempty"`
}

type CallbackNotification struct {
	Success       bool   `json:"success,omitempty"`
	Status        string `json:"status,omitempty"`
	TraceNumber   string `json:"trace_number,omitempty"`
	Reference     string `json:"reference,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	Amount        int    `json:"amount,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Message       string `json:"message,omitempty"`
	Timestamp     string `json:"timestamp,omitempty"`
}

type TransferRequest struct {
	To        string `json:"to" validate:"required"`
	Amount    int    `json:"amount" validate:"required,min=1"`
	Reason    string `json:"reason" validate:"required"`
	ShortCode string `json:"short_code" validate:"required"`
}

type TransferValidateResponse struct {
	Success      bool          `json:"success,omitempty"`
	Status       string        `json:"status,omitempty"`
	Message      string        `json:"message,omitempty"`
	To           string        `json:"to,omitempty"`
	Amount       int           `json:"amount,omitempty"`
	Reason       string        `json:"reason,omitempty"`
	ShortCode    string        `json:"short_code,omitempty"`
	CustomerInfo *CustomerInfo `json:"customer_info,omitempty"`
}

type CustomerInfo struct {
	Phone     string `json:"phone,omitempty"`
	Name      string `json:"name,omitempty"`
	AccountID string `json:"account_id,omitempty"`
}

type TransferResponse struct {
	Success       bool   `json:"success,omitempty"`
	Status        string `json:"status,omitempty"`
	Message       string `json:"message,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	To            string `json:"to,omitempty"`
	Amount        int    `json:"amount,omitempty"`
	Reference     string `json:"reference,omitempty"`
}

type ErrorDetails struct {
	Status     string `json:"status,omitempty"`
	StatusCode string `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
	Detail     string `json:"detail,omitempty"`
}

type ErrorResponse struct {
	Success bool          `json:"success,omitempty"`
	Message string        `json:"message,omitempty"`
	Error   *ErrorDetails `json:"error,omitempty"`
}
