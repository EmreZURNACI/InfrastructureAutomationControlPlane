package token

import (
	"strings"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/domain"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

func Validate(authHeader string) (*domain.AccessTokenClaims, error) {

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, errTokenNotValid
	}

	tokenStr := parts[1]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errTokenSignatureInvalid
		}
		return []byte(config.AppConfig.ServerConfig.SecretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, errTokenNotValid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errTokenClaimsInvalid
	}

	rawPerms, ok := claims["X-PERMISSIONS"].([]interface{})
	if !ok {
		return nil, errTokenClaimsInvalid
	}

	perms := make([]string, 0, len(rawPerms))
	for _, p := range rawPerms {
		if ps, ok := p.(string); ok {
			perms = append(perms, ps)
		}
	}

	return &domain.AccessTokenClaims{
		UserID:      claims["X-USERID"].(string),
		Permissions: perms,
		Timestamp:   claims["X-TIMESTAMP"].(string),
		Signature:   claims["X-SIGNATURE"].(string),
	}, nil
}
