package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/quelcom/sui-go/config"
	"github.com/quelcom/sui-go/modules"
)

type Page struct {
	conf *config.Config
	tpl  *template.Template
}

func NewPage(conf *config.Config, tpl *template.Template) *Page {
	for i, module := range conf.Modules {
		switch module.Name {
		case "Weather":
			dispatch(&conf.Modules[i], modules.GetWeather)
		case "Pi-hole":
			dispatch(&conf.Modules[i], modules.GetPiholeStats)
		}
	}

	return &Page{conf, tpl}
}

func dispatch(module *config.Module, f func(map[string]string) string) {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			<-ticker.C
			ch := make(chan string)
			go func() {
				ticker = time.NewTicker(module.UpdateInterval * time.Minute)
				weather := f(module.Data)
				ch <- weather
			}()

			module.Output = <-ch
		}
	}()
}

func (p *Page) FrontpageHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Add("Content-Type", "text/html")

	err := p.tpl.Execute(w, struct {
		Greeting  string
		Date      string
		Apps      []config.App
		Groups    []config.Group
		Providers []config.Provider
		Modules   []config.Module
	}{
		Greeting:  greet(p.conf.Name, time.Now().Hour()),
		Date:      time.Now().Format("Mon, Jan 02"),
		Apps:      p.conf.Apps,
		Groups:    p.conf.Groups,
		Providers: p.conf.Providers,
		Modules:   p.conf.Modules,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	elapsed := time.Since(start)
	log.Printf("Served %s in %s", r.URL.Path, elapsed)
}

// greet returns the greeting to be used in the h1 heading
func greet(name string, currentHour int) (greet string) {
	switch currentHour / 6 {
	case 0:
		greet = "Good night"
	case 1:
		greet = "Good morning"
	case 2:
		greet = "Good afternoon"
	default:
		greet = "Good evening"
	}

	if name != "" {
		return fmt.Sprintf("%s, %s!", greet, name)
	}

	return fmt.Sprintf("%s!", greet)
}
