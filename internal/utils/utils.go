package utils

import "database/sql"

func ToNullString(s string) sql.NullString{
	if s == ""{
		return sql.NullString{}
	}
	return sql.NullString{Valid: true, String: s}
}

func FromNullStr(ns sql.NullString) string{
	if ns.Valid{
		return ns.String
	}
	return ""
}