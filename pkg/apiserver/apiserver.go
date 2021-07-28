package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Config *Config
	Router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		Config: config,
		Router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()
	return http.ListenAndServe(s.Config.BindAddr, s.Router)
}

func (s *APIServer) configureRouter() {
	s.Router.HandleFunc("/", s.HandleHello())
}
