package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"

	"github.com/quelcom/sui-go/config"
)

var tpl = template.Must(template.New("index.html.tmpl").ParseFiles("../templates/index.html.tmpl"))

func TestGETRoot(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	conf := config.Config{}
	conf.Name = "Gopher"
	conf.Port = 3000

	page := NewPage(&conf, tpl)
	page.FrontpageHandler(response, request)

	t.Run("Page returns expected status code", func(t *testing.T) {
		got := response.Result().StatusCode
		want := 200

		if got != want {
			t.Errorf("Got %q but want %q", got, want)
		}
	})
}

func TestFindApps(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	appName := "My App"

	conf := config.Config{}
	conf.Apps = []config.App{
		{
			Name: appName,
			URL:  "https://bitwarden.mydomain.com",
			Icon: "shield-half-full",
		},
	}

	page := NewPage(&conf, tpl)
	page.FrontpageHandler(response, request)

	t.Run("Can find apps section", func(t *testing.T) {
		got := response.Body.String()
		if !strings.Contains(got, appName) {
			t.Errorf("Cannot find %q in %q", appName, got)
		}
	})
}

func TestGreetWithName(t *testing.T) {
	name := "Gopher"

	t.Run("Morning", func(t *testing.T) {
		got := greet(name, 9)
		expected := "Good morning, Gopher!"
		if got != expected {
			t.Errorf("Unexpected greeting. Got %q but expected %q", got, expected)
		}
	})
	t.Run("Afternoon", func(t *testing.T) {
		got := greet(name, 13)
		expected := "Good afternoon, Gopher!"
		if got != expected {
			t.Errorf("Unexpected greeting. Got %q but expected %q", got, expected)
		}
	})
	t.Run("Evening", func(t *testing.T) {
		got := greet(name, 18)
		expected := "Good evening, Gopher!"
		if got != expected {
			t.Errorf("Unexpected greeting. Got %q but expected %q", got, expected)
		}
	})
	t.Run("Night", func(t *testing.T) {
		got := greet(name, 0)
		expected := "Good night, Gopher!"
		if got != expected {
			t.Errorf("Unexpected greeting. Got %q but expected %q", got, expected)
		}
	})
}

func TestGreetWithoutName(t *testing.T) {
	var name string

	t.Run("Morning", func(t *testing.T) {
		got := greet(name, 9)
		expected := "Good morning!"
		if got != expected {
			t.Errorf("Unexpected greeting. Got %q but expected %q", got, expected)
		}
	})
	t.Run("Afternoon", func(t *testing.T) {
		got := greet(name, 13)
		expected := "Good afternoon!"
		if got != expected {
			t.Errorf("Unexpected greeting. Got %q but expected %q", got, expected)
		}
	})
	t.Run("Evening", func(t *testing.T) {
		got := greet(name, 18)
		expected := "Good evening!"
		if got != expected {
			t.Errorf("Unexpected greeting. Got %q but expected %q", got, expected)
		}
	})
	t.Run("Night", func(t *testing.T) {
		got := greet(name, 0)
		expected := "Good night!"
		if got != expected {
			t.Errorf("Unexpected greeting. Got %q but expected %q", got, expected)
		}
	})
}
