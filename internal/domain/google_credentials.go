package domain

import (
	"encoding/json"
	"fmt"
)

// ValidateServiceAccountKeyJSON checks whether the provided bytes represent a Google service account key JSON.
func ValidateServiceAccountKeyJSON(b []byte) error {
	if !json.Valid(b) {
		return fmt.Errorf("invalid json")
	}

	var tmp struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("failed to unmarshal json: %w", err)
	}
	if tmp.Type != "service_account" {
		return fmt.Errorf("expected type=service_account, got %q", tmp.Type)
	}
	return nil
}
