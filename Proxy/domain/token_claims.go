package domain

type TokenClaims struct {
	NickName    string   `json:"nickname"`
	DisplayName string   `json:"display_name"`
	SID         string   `json:"sid"`
	Role        string   `json:"role"`
	Groups      []string `json:"groups"`
	Permissions []string `json:"permissions"`
}
