package ldap

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/config"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/pkg/log"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlaneProxy/ports"
	"github.com/go-ldap/ldap/v3"
)

type LdapConnection struct {
	ldap *ldap.Conn
}

func Connect() (*LdapConnection, error) {
	conn, err := ldap.DialURL(config.AppConfig.LdapConfig.Url)
	if err != nil {
		return nil, err
	}

	if err := conn.Bind(config.AppConfig.LdapConfig.AdminUsername, config.AppConfig.LdapConfig.AdminPassword); err != nil {
		return nil, err
	}

	log.Logger.Info("LDAP connection succeed")
	return &LdapConnection{
		ldap: conn,
	}, nil

}

var _ ports.LDAP = (*LdapConnection)(nil)

func (h *LdapConnection) Bind(username string, password string) error {
	return h.ldap.Bind(username, password)
}
func (h *LdapConnection) Add(addRequest *ldap.AddRequest) error {
	return h.ldap.Add(addRequest)
}
func (h *LdapConnection) Modify(modifyRequest *ldap.ModifyRequest) error {
	return h.ldap.Modify(modifyRequest)
}
func (h *LdapConnection) Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error) {
	return h.ldap.Search(searchRequest)
}
func (h *LdapConnection) Close() (err error) {
	return h.ldap.Close()
}
