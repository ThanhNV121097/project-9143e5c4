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
	mux.HandleFunc("GET /v1/greeting", getGreeting(db))
	mux.HandleFunc("PUT /v1/greeting", putGreeting(db))
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
	}
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func getGreeting(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var greeting string
		if err := db.QueryRowContext(r.Context(), `SELECT text FROM greetings WHERE id = 1`).Scan(&greeting); err != nil {
			writeStoreError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"greeting": greeting})
	}
}

func putGreeting(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Greeting *string `json:"greeting"`
		}
		dec := json.NewDecoder(io.LimitReader(r.Body, 1<<20))
		dec.DisallowUnknownFields()
		if err := dec.Decode(&body); err != nil || body.Greeting == nil {
			writeError(w, http.StatusBadRequest, "MALFORMED_REQUEST", "Request body is malformed.")
			return
		}
		if dec.Decode(&struct{}{}) != io.EOF {
			writeError(w, http.StatusBadRequest, "MALFORMED_REQUEST", "Request body is malformed.")
			return
		}
		greeting := strings.TrimSpace(*body.Greeting)
		if greeting == "" {
			writeError(w, http.StatusUnprocessableEntity, "VALIDATION_FAILED", "Greeting must not be empty.")
			return
		}
		if err := db.QueryRowContext(r.Context(), `UPDATE greetings SET text = $1, updated_at = CURRENT_TIMESTAMP WHERE id = 1 RETURNING text`, greeting).Scan(&greeting); err != nil {
			writeStoreError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"greeting": greeting})
	}
}

func writeStoreError(w http.ResponseWriter, err error) {
	if isUnavailable(err) {
		writeError(w, http.StatusServiceUnavailable, "UNAVAILABLE", "Service unavailable.")
		return
	}
	writeError(w, http.StatusInternalServerError, "INTERNAL", "Internal server error.")
}

func isUnavailable(err error) bool {
	return errors.Is(err, context.Canceled) || strings.Contains(err.Error(), "connection refused")
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	data, err := json.Marshal(body)
	if err != nil {
		http.Error(w, "internal", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, _ = w.Write(data)
}

func writeError(w http.ResponseWriter, status int, code string, message string) {
	writeJSON(w, status, map[string]map[string]string{"error": map[string]string{"code": code, "message": message}})
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
