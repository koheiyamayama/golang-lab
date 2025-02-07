package main

import (
	"context"
	"database/sql"
)

type sqlite struct {
	db *sql.DB
}

func (s *sqlite) Version(ctx context.Context) (string, error) {
	v := s.db.QueryRowContext(ctx, "select sqlite_version()")
	vv := ""
	err := v.Scan(&vv)
	if err != nil {
		return "", err
	}
	return vv, nil
}
