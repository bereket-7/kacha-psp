package main

import (
	"log"
	"net/http"
	"os"

	"kacha-psp/internal/kacha"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get credentials from environment variables
	appID := os.Getenv("KACHA_APP_ID")
	apiKey := os.Getenv("KACHA_API_KEY")
	baseURL := os.Getenv("KACHA_BASE_URL") // Optional, defaults to https://api.kacha.com

	if appID == "" || apiKey == "" {
		log.Fatal("KACHA_APP_ID and KACHA_API_KEY environment variables are required")
	}

	// Initialize Kacha client
	var client *kacha.Client
	if baseURL != "" {
		client = kacha.NewClientWithBaseURL(appID, apiKey, baseURL)
	} else {
		client = kacha.NewClient(appID, apiKey)
	}

	// Set up Gin router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Payment request endpoint (OTP-based)
	r.POST("/api/payment/request", func(c *gin.Context) {
		var req kacha.PaymentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate required fields
		if req.Phone == "" || req.Amount <= 0 || req.TraceNumber == "" || req.Reason == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "phone, amount, trace_number, and reason are required"})
			return
		}

		resp, err := client.RequestPayment(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	// Payment authorize endpoint (OTP-based)
	r.POST("/api/payment/authorize", func(c *gin.Context) {
		var req kacha.PaymentAuthorizeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate required fields
		if req.Reference == "" || req.OTP == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "reference and otp are required"})
			return
		}

		resp, err := client.AuthorizePayment(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	// Push USSD payment request endpoint
	r.POST("/api/payment/push-ussd", func(c *gin.Context) {
		var req kacha.PushUSSDRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate required fields
		if req.Phone == "" || req.Amount <= 0 || req.TraceNumber == "" || req.CallbackURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "phone, amount, trace_number, and callback_url are required"})
			return
		}

		resp, err := client.RequestPushUSSD(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	// Callback endpoint for Push USSD notifications
	r.POST("/api/payment/callback", func(c *gin.Context) {
		var notification kacha.CallbackNotification
		if err := c.ShouldBindJSON(&notification); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Process the callback notification
		log.Printf("Received callback notification: %+v", notification)

		// TODO: Implement your business logic here
		// For example:
		// - Update database with transaction status
		// - Send notification to user
		// - Process order fulfillment

		// Respond to Kacha that we received the callback
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Callback received",
		})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

