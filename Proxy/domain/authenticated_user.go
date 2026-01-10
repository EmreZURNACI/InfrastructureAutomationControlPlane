package domain

type AuthUser struct {
	Username    string
	DisplayName string
	SID         string
	Role        string
	Groups      []string
}
