package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/domain"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(claims domain.TokenClaims) (string, error) {

	timestamp := time.Now().UTC().Format(time.RFC3339Nano)

	claim := jwt.MapClaims{
		"iss":           "System",
		"sub":           config.AppConfig.ServerConfig.AppName,
		"nick":          claims.NickName,
		"username":      claims.DisplayName,
		"groups":        claims.Groups,
		"X-TIMESTAMP":   timestamp,
		"X-SIGNATURE":   buildSignature(timestamp),
		"X-USERID":      claims.SID,
		"X-PERMISSIONS": claims.Permissions,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).
		SignedString([]byte(config.AppConfig.ServerConfig.SecretKey))

	if err != nil {
		return "", errTokenSignatureInvalid
	}

	return token, nil
}
func buildSignature(timestamp string) string {

	canonical := timestamp + "\n"

	h := hmac.New(sha256.New, []byte(config.AppConfig.ServerConfig.PrivateKey))
	h.Write([]byte(canonical))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signature
}
