package queries

import _ "embed"

var (
	//go:embed get_user.sql
	GetUserQuery string
)
