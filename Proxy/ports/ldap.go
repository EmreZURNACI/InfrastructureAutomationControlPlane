package ports

import "github.com/go-ldap/ldap/v3"

type LDAP interface {
	Bind(username string, password string) error
	Add(addRequest *ldap.AddRequest) error
	Modify(modifyRequest *ldap.ModifyRequest) error
	Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error)
	Close() (err error)
}
