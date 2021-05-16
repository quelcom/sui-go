package modules

import (
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	fetchWeather = mockFetchWeather
	defer func() { fetchWeather = actualFetch }()

	data := map[string]string{"city": "Helsinki"}

	t.Run("Return weather info", func(t *testing.T) {
		got := GetWeather(data)
		expectedSubstring := "Temperature in Helsinki: 30 degrees, current condition is: Partly cloudy"
		if !strings.HasPrefix(got, expectedSubstring) {
			t.Errorf("Cannot find %q in %q", expectedSubstring, got)
		}
	})
}

func mockFetchWeather(url string) ([]byte, error) {
	return []byte(`{"current_condition": [{"temp_C":"30","weatherDesc": [{"value": "Partly cloudy"}]}]}`), nil
}
