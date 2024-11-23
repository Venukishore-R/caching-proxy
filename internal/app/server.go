package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Venukishore-R/caching-proxy/internal/proxy-server/proxy"
)

type Server struct {
	Port       int
	Origin     string
	ClearCache bool
	Proxy      *proxy.Proxy
}

func (app *Server) StartServer() {
	if app.ClearCache {
		fmt.Println("Clearing Cache...")
		app.Proxy.ClearCache()
	}

	if app.Origin != "" || app.Port != 0 {
		fmt.Printf("Starting caching proxy server on port %d and forwarding requests to %s\n", app.Port, app.Origin)
		http.Handle("/", app.Proxy)

		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.Port), nil))
	}
}
