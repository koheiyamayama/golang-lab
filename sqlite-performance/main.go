package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	ctx := context.Background()
	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		slog.ErrorContext(ctx, "Error opening database", slog.Any("error", err))
		os.Exit(1)
	}
	defer db.Close()

	v := db.QueryRowContext(ctx, "select sqlite_version()")

	vv := ""
	v.Scan(&vv)
	fmt.Println(vv)
}
