package auth

import (
	"context"
	"fmt"
	"unicode/utf16"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/ports"
	"github.com/go-ldap/ldap/v3"
)

type SignUpRequest struct {
	FirstName   string  `json:"first_name" validate:"required"`
	LastName    string  `json:"last_name" validate:"required"`
	UserName    string  `json:"user_name" validate:"required"`
	Email       string  `json:"email" validate:"required"`
	PhoneNumber string  `json:"phone_number" validate:"required"`
	Password    string  `json:"password" validate:"required"`
	Description *string `json:"description,omitempty" validate:"omitempty"`
}
type SignUpResponse struct {
	Message string `json:"message,omitempty"`
}
type SignUpHandler struct {
	ldapCon ports.LDAP
}

func NewSignUpHandler(ldapCon ports.LDAP) *SignUpHandler {
	return &SignUpHandler{
		ldapCon: ldapCon,
	}
}

func (h *SignUpHandler) Handle(ctx context.Context, req *SignUpRequest) (*SignUpResponse, error) {

	if err := h.ldapCon.Add(AddUserRequest(req)); err != nil {
		return nil, err
	}

	if err := h.ldapCon.Modify(SetPasswordRequest(req.UserName, req.Password)); err != nil {
		return nil, err
	}

	if err := h.ldapCon.Modify(ModifyUserRequest(req.UserName)); err != nil {
		return nil, err
	}

	return &SignUpResponse{
		Message: "Kullanıcı başarıyla oluşturuldu.",
	}, nil

}
func AddUserRequest(req *SignUpRequest) *ldap.AddRequest {

	attributes := []ldap.Attribute{
		{
			Type: "objectClass",
			Vals: []string{
				"top",
				"person",
				"organizationalPerson",
				"user",
			},
		},
		{
			Type: "sAMAccountName",
			Vals: []string{req.UserName},
		},
		{
			Type: "userPrincipalName",
			Vals: []string{req.UserName + "@emrezlab.local"},
		},
		{
			Type: "cn",
			Vals: []string{req.UserName},
		},
		{
			Type: "displayName",
			Vals: []string{req.FirstName + " " + req.LastName},
		},
		{
			Type: "givenName",
			Vals: []string{req.FirstName},
		},
		{
			Type: "sn",
			Vals: []string{req.LastName},
		},
		{
			Type: "telephoneNumber",
			Vals: []string{req.PhoneNumber},
		},
		{
			Type: "mail",
			Vals: []string{req.Email},
		},
		{
			Type: "userAccountControl",
			Vals: []string{"514"},
		},
	}

	if req.Description != nil && *req.Description != "" {
		attributes = append(attributes, ldap.Attribute{
			Type: "description",
			Vals: []string{*req.Description},
		})
	}

	return &ldap.AddRequest{
		DN: "CN=" + req.UserName +
			",OU=_IT,OU=_Istanbul,OU=_Users,DC=emrezlab,DC=local",
		Attributes: attributes,
	}
}
func SetPasswordRequest(username, password string) *ldap.ModifyRequest {

	pwd := fmt.Sprintf("\"%s\"", password)
	encodedPwd := utf16.Encode([]rune(pwd))

	buf := make([]byte, len(encodedPwd)*2)
	for i, v := range encodedPwd {
		buf[i*2] = byte(v)
		buf[i*2+1] = byte(v >> 8)
	}

	return &ldap.ModifyRequest{
		DN: "CN=" + username +
			",OU=_IT,OU=_Istanbul,OU=_Users,DC=emrezlab,DC=local",
		Changes: []ldap.Change{
			{
				Operation: ldap.ReplaceAttribute,
				Modification: ldap.PartialAttribute{
					Type: "unicodePwd",
					Vals: []string{string(buf)},
				},
			},
		},
	}
}
func ModifyUserRequest(username string) *ldap.ModifyRequest {
	return &ldap.ModifyRequest{
		DN: "CN=" + username +
			",OU=_IT,OU=_Istanbul,OU=_Users,DC=emrezlab,DC=local",
		Changes: []ldap.Change{
			{
				Operation: ldap.ReplaceAttribute,
				Modification: ldap.PartialAttribute{
					Type: "userAccountControl",
					Vals: []string{"512"},
				},
			},
		},
	}
}

/*
* INFO: AD'de oluşturalan kullanıcı DISABLE(514) oluştur.
* 		Bunu başka bir request içinde modify edip ENABLE(512) yapmak gerekir.
* 		Set password ve set enable aynı requestte olmaz.
 */
