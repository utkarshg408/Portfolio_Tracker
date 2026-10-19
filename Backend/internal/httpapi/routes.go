package httpapi

import (
	"database/sql"
	"encoding/json"
	"github.com/utkarshg408/Portfolio_Tracker/Backend/internal/market"
	"net/http"
)

func NewMux(db *sql.DB, marketClient *market.FinnhubClient) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health)
	mux.HandleFunc("GET /ready", ready(db))
	mux.HandleFunc("GET /api/v1/market/quote", quote(marketClient))
	return mux
}
func health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
func ready(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.PingContext(r.Context()); err != nil {
			writeJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "database unavailable"})
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
	}
}
func quote(client *market.FinnhubClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := client.Quote(r.Context(), r.URL.Query().Get("symbol"))
		if err != nil {
			writeJSON(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, result)
	}
}
func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
