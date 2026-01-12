package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/domain"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/ports"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/token"
	"github.com/go-ldap/ldap/v3"
)

type SignInRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type SignInResponse struct {
	Token *string `json:"token"`
}
type SignInHandler struct {
	repository Repository
	ldapCon    ports.LDAP
}

func NewSignInHandler(repository Repository, ldapCon ports.LDAP) *SignInHandler {
	return &SignInHandler{
		repository: repository,
		ldapCon:    ldapCon,
	}
}
func (h *SignInHandler) Handle(ctx context.Context, req *SignInRequest) (*SignInResponse, error) {

	searchResp, err := h.ldapCon.Search(GetUserRequest(req.Username))
	if err != nil {
		return nil, errors.New("kullanıcı arama hatası")
	}

	if len(searchResp.Entries) != 1 {
		return nil, errors.New("kullanıcı bulunamadı")
	}

	entry := searchResp.Entries[0]

	dn := entry.DN

	if err := h.ldapCon.Bind(dn, req.Password); err != nil {
		return nil, errors.New("kullanıcı adı veya şifre hatalı")
	}

	username := entry.GetAttributeValue("sAMAccountName")
	displayName := entry.GetAttributeValue("displayName")
	groups := entry.GetAttributeValues("memberOf")
	sid := entry.GetAttributeValue("objectSid")
	role := UserRole(groups)
	roleID := h.repository.GetRole(role)
	permissions := h.repository.GetPermissions(roleID)

	tokenClaims := domain.TokenClaims{
		NickName:    username,
		Groups:      groups,
		SID:         sid,
		DisplayName: displayName,
		Role:        role,
		Permissions: permissions,
	}

	token, err := token.CreateAccessToken(tokenClaims)
	if err != nil {
		return nil, err
	}

	return &SignInResponse{
		Token: &token,
	}, nil
}
func UserRole(groups []string) (role string) {

	for _, g := range groups {
		if strings.Contains(g, "CN=Administrators") {
			return "admin"
		}
	}
	return "user"
}
func GetUserRequest(username string) *ldap.SearchRequest {
	filter := fmt.Sprintf("(sAMAccountName=%s)", username)
	attributes := []string{
		"dn",
		"sAMAccountName",
		"displayName",
		"memberOf",
		"objectSid",
	}
	return &ldap.SearchRequest{
		BaseDN:       "DC=emrezlab,DC=local",
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       filter,
		Attributes:   attributes,
		Controls:     nil,
	}
}
