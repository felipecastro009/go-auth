package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type Handler struct {
	Method  string
	Handler http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]Handler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]Handler),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, method string, handler http.HandlerFunc) {
	s.Handlers[path] = Handler{
		Method:  method,
		Handler: handler,
	}
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		switch handler.Method {
		case http.MethodGet:
			s.Router.Get(path, handler.Handler)
		case http.MethodPost:
			s.Router.Post(path, handler.Handler)
		case http.MethodPut:
			s.Router.Put(path, handler.Handler)
		case http.MethodPatch:
			s.Router.Patch(path, handler.Handler)
		case http.MethodDelete:
			s.Router.Delete(path, handler.Handler)
		default:
			log.Fatalf("Method %s not supported for path %s", handler.Method, path)
		}
	}
	log.Println("Starting server on port", s.WebServerPort)
	err := http.ListenAndServe(s.WebServerPort, s.Router)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
