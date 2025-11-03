package main

import (
	"fmt"
	"log"

	"kacha-psp/config"
	kacha "kacha-psp/internal"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Kacha client
	client := kacha.NewClient(cfg.KachaAppID, cfg.KachaAPIKey)

	// Example 1: OTP-Based Payment Flow
	fmt.Println("=== Example 1: OTP-Based Payment Flow ===")
	
	// Step 1: Request payment (sends OTP to customer)
    	paymentReq := kacha.PaymentRequest{
		Phone:       "251913609212",
		Amount:      100,
		TraceNumber: "70RNVPO548",
		Reason:      "payment",
	}

	paymentResp, err := client.RequestPayment(paymentReq)
	if err != nil {
		log.Fatalf("Failed to request payment: %v", err)
	}

	fmt.Printf("Payment requested successfully!\n")
	fmt.Printf("Reference: %s\n", paymentResp.Reference)
	fmt.Printf("Status: %s\n", paymentResp.Status)
	fmt.Printf("Message: %s\n", paymentResp.Message)

	// Step 2: Authorize payment (customer enters OTP)
	// In a real scenario, you would get the OTP from the customer
	authReq := kacha.PaymentAuthorizeRequest{
		Reference: paymentResp.Reference,
		OTP:       657894, // This would come from customer input
	}

	authResp, err := client.AuthorizePayment(authReq)
	if err != nil {
		log.Fatalf("Failed to authorize payment: %v", err)
	}

	fmt.Printf("\nPayment authorized successfully!\n")
	fmt.Printf("Status: %s\n", authResp.Status)
	fmt.Printf("Transaction ID: %s\n", authResp.TransactionID)

	// Example 2: Push USSD Payment Flow
	fmt.Println("\n=== Example 2: Push USSD Payment Flow ===")

	pushReq := kacha.PushUSSDRequest{
		Phone:       "251913609212",
		Amount:      200,
		TraceNumber: "70RNVPO549",
		CallbackURL: "https://fenan-domain.com/api/payment/callback",
	}

	pushResp, err := client.RequestPushUSSD(pushReq)
	if err != nil {
		log.Fatalf("Failed to initiate push USSD payment: %v", err)
	}

	fmt.Printf("Push USSD payment initiated successfully!\n")
	fmt.Printf("Status: %s\n", pushResp.Status)
	fmt.Printf("Trace Number: %s\n", pushResp.TraceNumber)
	fmt.Printf("Message: %s\n", pushResp.Message)
	fmt.Println("\nNote: Transaction result will be sent to callback URL")

	// Example 3: B2C Transfer Flow
	fmt.Println("\n=== Example 3: B2C Transfer Flow ===")

	// Step 1: Validate transfer
	    transferReq := kacha.TransferRequest{
	        To:        "251913609212",
		Amount:    100,
		Reason:    "fee",
		ShortCode: "7865",
	}

	validateResp, err := client.ValidateTransfer(transferReq)
	if err != nil {
		log.Fatalf("Failed to validate transfer: %v", err)
	}

	fmt.Printf("Transfer validated successfully!\n")
	fmt.Printf("Status: %s\n", validateResp.Status)
	fmt.Printf("Amount: %d\n", validateResp.Amount)
	fmt.Printf("To: %s\n", validateResp.To)
	if validateResp.CustomerInfo != nil {
		fmt.Printf("Customer Name: %s\n", validateResp.CustomerInfo.Name)
		fmt.Printf("Customer Phone: %s\n", validateResp.CustomerInfo.Phone)
	}

	// Step 2: Execute transfer (only if validation was successful)
	if validateResp.Status == "PREPARED" {
		transferResp, err := client.Transfer(transferReq)
		if err != nil {
			log.Fatalf("Failed to execute transfer: %v", err)
		}

		fmt.Printf("\nTransfer executed successfully!\n")
		fmt.Printf("Status: %s\n", transferResp.Status)
		fmt.Printf("Transaction ID: %s\n", transferResp.TransactionID)
		fmt.Printf("Reference: %s\n", transferResp.Reference)
	} else {
		fmt.Println("\nTransfer validation did not return PREPARED status, skipping transfer execution")
	}
}