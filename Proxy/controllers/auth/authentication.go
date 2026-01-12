package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"time"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/controllers"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/config"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/token"

	"github.com/gofiber/fiber/v2"
)

func verifySignature(timestamp, signature string) bool {

	canonical := timestamp + "\n"

	h := hmac.New(sha256.New, []byte(config.AppConfig.ServerConfig.PrivateKey))
	h.Write([]byte(canonical))
	expectedSig := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return hmac.Equal([]byte(expectedSig), []byte(signature))
}

func isValidTimestamp(timestamp string) bool {
	t, err := time.Parse(time.RFC3339Nano, timestamp)
	if err != nil {
		return false
	}

	now := time.Now().UTC()

	if t.After(now) {
		return false
	}

	return now.Sub(t) <= 10*time.Minute
}

func Authentication(c *fiber.Ctx) error {

	_token := c.Get("Authorization")
	if _token == "" {
		return c.JSON(controllers.FailureResponse(401, "Unauthorized", "Unauthorized"))
	}

	claims, err := token.Validate(_token)
	if err != nil {
		return c.JSON(controllers.FailureResponse(401, "Unauthorized", "Unauthorized"))
	}

	if !isValidTimestamp(claims.Timestamp) {
		return c.JSON(controllers.FailureResponse(401, "Unauthorized", "Unauthorized"))
	}

	if !verifySignature(claims.Timestamp, claims.Signature) {
		return c.JSON(controllers.FailureResponse(401, "Unauthorized", "Unauthorized"))
	}

	c.Locals("X-USER-PERMISSIONS", claims.Permissions)
	c.Locals("X-USER-ID", claims.UserID)

	return c.Next()
}
