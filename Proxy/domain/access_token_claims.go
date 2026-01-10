package domain

type AccessTokenClaims struct {
	UserID      string   `json:"user_id"`
	Permissions []string `json:"permissions"`
	Timestamp   string   `json:"timestamp"`
	Signature   string   `json:"signature"`
}
