# Kacha PSP Integration

This Go application provides integration with the Kacha Payment Service Provider (PSP) API for handling Consumer-to-Business (C2B) payments.

## Features

- **OTP-Based Payment Authentication**: Initiate payment requests that send OTPs to customers via SMS
- **Payment Authorization**: Authorize payments using reference numbers and OTPs
- **Push USSD Payment**: Initiate direct push payment transactions via USSD authentication
- **Callback Handling**: Receive and process asynchronous transaction notifications

## API Endpoints

### OTP-Based Payment Flow

1. **Initiate Payment Request** (`POST /api/payment/request`)
   - Sends an OTP to the customer via SMS
   - Returns payment request details including reference number

2. **Authorize Payment** (`POST /api/payment/authorize`)
   - Approves a payment request using reference and OTP

### Push USSD Payment Flow

1. **Push USSD Payment Request** (`POST /api/payment/push-ussd`)
   - Initiates a push USSD payment transaction
   - Requires a callback URL for asynchronous notifications

2. **Callback Handler** (`POST /api/payment/callback`)
   - Receives transaction results from Kacha
   - Process payment status updates

## Setup

### Prerequisites

- Go 1.25.3 or higher
- Kacha App ID and API Key from Merchant Portal

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd kacha-psp
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set environment variables:
```bash
export KACHA_APP_ID="your-app-id"
export KACHA_API_KEY="your-api-key"
export KACHA_BASE_URL="https://api.kacha.com"  # Optional
export PORT="8080"  # Optional, defaults to 8080
```

### Running the Server

```bash
go run cmd/server/main.go
```

The server will start on port 8080 (or the port specified in the PORT environment variable).

## Usage Examples

### OTP-Based Payment Flow

#### 1. Initiate Payment Request

```bash
curl -X POST http://localhost:8080/api/payment/request \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "251989840054",
    "amount": 100,
    "trace_number": "70RNVPO548",
    "reason": "payment"
  }'
```

Response:
```json
{
  "success": true,
  "reference": "2CU210EXT4",
  "message": "OTP sent successfully",
  "status": "pending",
  "trace_number": "70RNVPO548"
}
```

#### 2. Authorize Payment

```bash
curl -X POST http://localhost:8080/api/payment/authorize \
  -H "Content-Type: application/json" \
  -d '{
    "reference": "2CU210EXT4",
    "otp": 657894
  }'
```

### Push USSD Payment Flow

```bash
curl -X POST http://localhost:8080/api/payment/push-ussd \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "251989840054",
    "amount": 100,
    "trace_number": "70RNVPO549",
    "callback_url": "https://your-domain.com/api/payment/callback"
  }'
```

## Using the Kacha Client Directly

You can also use the Kacha client directly in your application:

```go
package main

import (
    "kacha-psp/internal/kacha"
)

func main() {
    // Initialize client
    client := kacha.NewClient("your-app-id", "your-api-key")
    
    // Make payment request
    req := kacha.PaymentRequest{
        Phone:       "251989840054",
        Amount:      100,
        TraceNumber: "70RNVPO548",
        Reason:      "payment",
    }
    
    resp, err := client.RequestPayment(req)
    if err != nil {
        // Handle error
        return
    }
    
    // Use response (contains reference number)
    reference := resp.Reference
    
    // Authorize payment
    authReq := kacha.PaymentAuthorizeRequest{
        Reference: reference,
        OTP:       657894,
    }
    
    authResp, err := client.AuthorizePayment(authReq)
    if err != nil {
        // Handle error
        return
    }
    
    // Payment authorized
}
```

## Project Structure

```
kacha-psp/
├── cmd/
│   └── server/
│       └── main.go          # Example server implementation
├── internal/
│   └── kacha/
│       ├── client.go        # Kacha API client
│       ├── payment.go       # Payment methods
│       └── types.go         # Request/response types
├── go.mod
├── go.sum
└── README.md
```

## Important Notes

1. **Trace Number**: Must be unique for every payment request
2. **Callback URL**: Required for Push USSD payments. Must be publicly accessible
3. **Authentication**: Uses Basic Authentication with App ID as username and API Key as password
4. **Error Handling**: All methods return errors that should be handled appropriately
5. **Production**: Remember to disable debug mode and set appropriate timeouts in production

## Testing

You can test the endpoints using the provided curl examples or tools like Postman. Make sure to:

1. Set valid Kacha credentials
2. Use valid phone numbers (Ethiopian format: 251XXXXXXXXX)
3. Ensure callback URLs are publicly accessible for Push USSD
4. Generate unique trace numbers for each request

## License

[Your License Here]

