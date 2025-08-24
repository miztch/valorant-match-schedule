package country

import (
	"testing"
)

func TestCountries(t *testing.T) {
	tests := []struct {
		code           string
		expectedName   string
		expectedRegion string
	}{
		{"jp", "Japan", "PACIFIC"},
		{"us", "United States", "AMERICAS"},
		{"de", "Germany", "EMEA"},
		{"cn", "China", "CHINA"},
		{"gb", "United Kingdom", "EMEA"},
		{"kr", "South Korea", "PACIFIC"},
		{"br", "Brazil", "AMERICAS"},
	}

	for _, test := range tests {
		t.Run(test.code, func(t *testing.T) {
			info, exists := Countries[test.code]
			if !exists {
				t.Fatalf("Country code %s not found in Countries", test.code)
			}
			if info.Name != test.expectedName {
				t.Errorf("Expected name %s, got %s", test.expectedName, info.Name)
			}
			if info.Region != test.expectedRegion {
				t.Errorf("Expected region %s, got %s", test.expectedRegion, info.Region)
			}
		})
	}
}

func TestCountriesCompleteness(t *testing.T) {
	// Check for presence of some important country codes
	importantCountries := []string{"jp", "us", "de", "cn", "gb", "kr", "br", "fr", "ca", "au"}
	for _, code := range importantCountries {
		if _, exists := Countries[code]; !exists {
			t.Errorf("Important country code %s is missing from Countries", code)
		}
	}

	// Check that all regions are used
	regions := make(map[string]bool)
	for _, info := range Countries {
		regions[info.Region] = true
	}
	expectedRegions := []string{"EMEA", "AMERICAS", "PACIFIC", "CHINA"}
	for _, region := range expectedRegions {
		if !regions[region] {
			t.Errorf("Region %s is not used in any country", region)
		}
	}
}

func TestSubAreasMap(t *testing.T) {
	// Check that SubAreas is defined correctly
	expectedRegions := []string{"PACIFIC", "AMERICAS", "EMEA", "CHINA"}
	for _, region := range expectedRegions {
		if _, exists := SubAreas[region]; !exists {
			t.Errorf("SubAreas is missing region %s", region)
		}
	}
}

func TestOrganizersMap(t *testing.T) {
	// Check that Organizers is defined correctly
	expectedRegions := []string{"PACIFIC", "AMERICAS", "EMEA", "CHINA"}
	for _, region := range expectedRegions {
		if _, exists := Organizers[region]; !exists {
			t.Errorf("Organizers is missing region %s", region)
		}
	}
}

func TestInternationalEvents(t *testing.T) {
	// Check that InternationalEvents is not empty
	if len(InternationalEvents) == 0 {
		t.Error("InternationalEvents should not be empty")
	}
}
