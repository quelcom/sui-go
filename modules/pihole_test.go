package modules

import (
	"strings"
	"testing"
)

func TestGetPihole(t *testing.T) {
	fetchPihole = mockFetchPihole
	defer func() { fetchPihole = actualFetch }()

	t.Run("Return Pi-hole stats", func(t *testing.T) {
		got := GetPiholeStats(map[string]string{})
		expectedSubstring := "Status: enabled. Ads blocked today: 34.6%"
		if !strings.HasPrefix(got, expectedSubstring) {
			t.Errorf("Cannot find %q in %q", expectedSubstring, got)
		}
	})
}

func mockFetchPihole(url string) ([]byte, error) {
	return []byte(`{"ads_percentage_today":34.56789,"status":"enabled"}`), nil
}
