package main

import (
	"log"
	"net/http"

	"kacha-psp/config"
	kacha "kacha-psp/kacha"

	"github.com/gin-gonic/gin"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }

    // Initialize Kacha client using username/password from config
    username := cfg.KachaUsername
    password := cfg.KachaPassword
    var client *kacha.Client
    if cfg.KachaBaseURL != "" {
        client = kacha.NewClientWithBaseURL(username, password, cfg.KachaBaseURL)
    } else {
        client = kacha.NewClient(username, password)
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

		// Respond to Kacha that we received the callback
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Callback received",
		})
	})

	// Transfer validate endpoint (B2C)
	r.POST("/api/transfer/validate", func(c *gin.Context) {
		var req kacha.TransferRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate required fields
		if req.To == "" || req.Amount <= 0 || req.Reason == "" || req.ShortCode == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "to, amount, reason, and short_code are required"})
			return
		}

		resp, err := client.ValidateTransfer(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	// Transfer endpoint (B2C)
	r.POST("/api/transfer", func(c *gin.Context) {
		var req kacha.TransferRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate required fields
		if req.To == "" || req.Amount <= 0 || req.Reason == "" || req.ShortCode == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "to, amount, reason, and short_code are required"})
			return
		}

		resp, err := client.Transfer(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

    // Start server
    log.Printf("Server starting on port %s", cfg.Port)
    if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}

