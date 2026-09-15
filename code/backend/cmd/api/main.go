package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ThanhNV121097/project-9143e5c4/backend/migrations"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := migrate(ctx, db); err != nil {
		log.Fatal(err)
	}
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		if err := db.PingContext(r.Context()); err != nil {
			http.Error(w, "unavailable", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
	}
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func migrate(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (version text PRIMARY KEY, applied_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP)`); err != nil {
		return fmt.Errorf("create migration table: %w", err)
	}
	entries, err := fs.ReadDir(migrations.Files, ".")
	if err != nil {
		return err
	}
	var names []string
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".up.sql") {
			names = append(names, entry.Name())
		}
	}
	sort.Strings(names)
	for _, name := range names {
		var applied bool
		if err := db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version = $1)`, name).Scan(&applied); err != nil || applied {
			if err != nil { return fmt.Errorf("check migration %s: %w", name, err) }
			continue
		}
		sqlBytes, err := migrations.Files.ReadFile(name)
		if err != nil { return err }
		tx, err := db.BeginTx(ctx, nil)
		if err != nil { return err }
		if _, err = tx.ExecContext(ctx, string(sqlBytes)); err == nil {
			_, err = tx.ExecContext(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, name)
		}
		if err != nil { _ = tx.Rollback(); return fmt.Errorf("apply migration %s: %w", name, err) }
		if err = tx.Commit(); err != nil { return err }
	}
	return nil
}
