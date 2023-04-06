package cns

import "time"

const (
	AppName = "Dutchman"
	AppUrl  = "https://dutchman.com"

	MaxPageSize = 1000
)

var (
	AppTimeLocation = time.FixedZone("AST", 21600) // +0600

	True = true
)

// Roles
const (
	RoleGuest = "guest"
	RoleAdmin = "admin"
)

func RoleIsValid(v string) bool {
	return v == RoleGuest ||
		v == RoleAdmin
}
