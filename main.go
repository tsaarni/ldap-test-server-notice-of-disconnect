package main

import (
	"log"
	"net"

	"github.com/nmcclain/ldap"
)

func main() {
	s := ldap.NewServer()
	handler := ldapHandler{}
	s.BindFunc("", handler)
	s.SearchFunc("", handler)

	tls := false

	if tls {
		if err := s.ListenAndServeTLS("localhost:9011", "/home/tsaarni/packages/openldap/tests/testrun/tls/certs/localhost.crt", "/home/tsaarni/packages/openldap/tests/testrun/tls/private/localhost.key"); err != nil {
			log.Fatal("LDAP Server Failed: %s", err.Error())
		}
	} else {
		if err := s.ListenAndServe("localhost:9011"); err != nil {
			log.Fatal("LDAP Server Failed: %s", err.Error())
		}
	}
}

type ldapHandler struct {
}

func (h ldapHandler) Bind(bindDN, bindSimplePw string, conn net.Conn) (ldap.LDAPResultCode, error) {
	if bindDN == "" && bindSimplePw == "" {
		return ldap.LDAPResultSuccess, nil
	}
	if bindDN == "cn=Barbara Jensen,ou=Information Technology Division,ou=People,dc=example,dc=com" && bindSimplePw == "bjensen" {
		return ldap.LDAPResultSuccess, nil
	}
	if bindDN == "cn=manager,dc=example,dc=com" && bindSimplePw == "secret" {
		return ldap.LDAPResultSuccess, nil
	}
	return ldap.LDAPResultInvalidCredentials, nil
}

func (h ldapHandler) Search(boundDN string, searchReq ldap.SearchRequest, conn net.Conn) (ldap.ServerSearchResult, error) {
	entries := []*ldap.Entry{
		&ldap.Entry{"cn=ned," + searchReq.BaseDN, []*ldap.EntryAttribute{
			&ldap.EntryAttribute{"cn", []string{"ned"}},
			&ldap.EntryAttribute{"uidNumber", []string{"5000"}},
			&ldap.EntryAttribute{"accountStatus", []string{"active"}},
			&ldap.EntryAttribute{"uid", []string{"ned"}},
			&ldap.EntryAttribute{"description", []string{"ned"}},
			&ldap.EntryAttribute{"objectClass", []string{"posixAccount"}},
		}},
		&ldap.Entry{"cn=trent," + searchReq.BaseDN, []*ldap.EntryAttribute{
			&ldap.EntryAttribute{"cn", []string{"trent"}},
			&ldap.EntryAttribute{"uidNumber", []string{"5005"}},
			&ldap.EntryAttribute{"accountStatus", []string{"active"}},
			&ldap.EntryAttribute{"uid", []string{"trent"}},
			&ldap.EntryAttribute{"description", []string{"trent"}},
			&ldap.EntryAttribute{"objectClass", []string{"posixAccount"}},
		}},
	}
	return ldap.ServerSearchResult{entries, []string{}, []ldap.Control{}, ldap.LDAPResultSuccess}, nil
}
