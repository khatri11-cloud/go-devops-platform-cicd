package handlers

import (
	"encoding/json"
	"net/http"
)

type VersionResponse struct {
	Version string `json:"version"`
}

func VersionHandler(version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(VersionResponse{
			Version: version,
		})
	}
}
