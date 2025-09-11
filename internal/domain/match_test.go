package domain

import (
	"testing"
)

func TestIsInternational(t *testing.T) {
	tests := []struct {
		name     string
		eventId  int
		expected bool
	}{
		{
			name:     "International event (exact match)",
			eventId:  2097, // Valorant Champions 2024
			expected: true,
		},
		{
			name:     "International event (partial match)",
			eventId:  209, // non-existent
			expected: false,
		},
		{
			name:     "Non-international event",
			eventId:  2098, // Southern Esports Conference 2024: Spring Playoffs
			expected: false,
		},
		{
			name:     "Empty value",
			eventId:  0,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Match{EventId: tt.eventId}
			result := m.IsInternational()
			if result != tt.expected {
				t.Errorf("isInternationalEvent(%d) = %v, want %v", tt.eventId, result, tt.expected)
			}
		})
	}
}
