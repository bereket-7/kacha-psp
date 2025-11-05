package main

import (
	"kacha-psp/config"
	kacha "kacha-psp/kacha"
	"kacha-psp/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/otp/pay", func(c *gin.Context) {
		var req kacha.PaymentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Username == "" || req.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
			return
		}
		if req.Phone == "" || req.Amount <= 0 || req.TraceNumber == "" || req.Reason == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "phone, amount, trace_number, and reason are required"})
			return
		}

		client := kacha.NewClientWithBaseURL(req.Username, req.Password, cfg.KachaBaseURL)
		resp, err := client.RequestPayment(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	r.POST("/otp/authorize", func(c *gin.Context) {
		var req kacha.PaymentAuthorizeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Username == "" || req.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
			return
		}
		if req.Reference == "" || req.OTP == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "reference and otp are required"})
			return
		}

		client := kacha.NewClientWithBaseURL(req.Username, req.Password, cfg.KachaBaseURL)
		resp, err := client.AuthorizePayment(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})
	
		// Push USSD payment request endpoint
	r.POST("/pay", func(c *gin.Context) {
		var req kacha.PushUSSDRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Username == "" || req.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
			return
		}
		if req.Phone == "" || req.Amount <= 0 || req.TraceNumber == "" || req.CallbackURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "phone, amount, trace_number, and callback_url are required"})
			return
		}

		client := kacha.NewClientWithBaseURL(req.Username, req.Password, cfg.KachaBaseURL)
		kachaResp, err := client.RequestPushUSSD(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		pspResp := utils.MapPushUSSDToPSP(kachaResp, err == nil)
		c.JSON(http.StatusOK, pspResp)
	})

	r.POST("/callback", func(c *gin.Context) {
		var notification kacha.CallbackNotification
		if err := c.ShouldBindJSON(&notification); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Printf("Received callback notification: %+v", notification)
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Callback received"})
	})

	r.POST("/withdrawal/validate", func(c *gin.Context) {
		var req kacha.TransferRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Username == "" || req.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
			return
		}
		if req.To == "" || req.Amount <= 0 || req.Reason == "" || req.ShortCode == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "to, amount, reason, and short_code are required"})
			return
		}

		client := kacha.NewClientWithBaseURL(req.Username, req.Password, cfg.KachaBaseURL)
		resp, err := client.ValidateTransfer(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})


	// B2C Transfer endpoint
	r.POST("/withdrawal", func(c *gin.Context) {
		var req kacha.TransferRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Username == "" || req.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
			return
		}
		if req.To == "" || req.Amount <= 0 || req.Reason == "" || req.ShortCode == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "to, amount, reason, and short_code are required"})
			return
		}

		client := kacha.NewClientWithBaseURL(req.Username, req.Password, cfg.KachaBaseURL)
		kachaResp, err := client.Transfer(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		pspResp := utils.MapTransferToPSP(kachaResp, err == nil)
		c.JSON(http.StatusOK, pspResp)
	})

	log.Printf("Starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
