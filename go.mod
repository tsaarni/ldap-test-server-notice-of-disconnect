module github.com/tsaarni/ldap-test-server-notice-of-disconnect

go 1.16

replace github.com/nmcclain/ldap => ./internal/github-com-nmcclain-ldap-fork

require (
	github.com/nmcclain/asn1-ber v0.0.0-20170104154839-2661553a0484 // indirect
	github.com/nmcclain/ldap v0.0.0-20191021200707-3b3b69a7e9e3
)
