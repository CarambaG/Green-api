package server

import (
	"GREEN-API/config"
	appHandlers "GREEN-API/handlers"
	"GREEN-API/internal/logger"
	"GREEN-API/middleware"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	config *config.Config
}

func New(cfg *config.Config) *Server {
	r := mux.NewRouter()
	s := &Server{
		router: r,
		config: cfg,
	}

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	// API routes
	s.router.HandleFunc("/api/settings", appHandlers.GetSettings).Methods(http.MethodGet, http.MethodOptions)
	s.router.HandleFunc("/api/state", appHandlers.GetStateInstance).Methods(http.MethodGet, http.MethodOptions)
	s.router.HandleFunc("/api/message", appHandlers.SendMessage).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/file", appHandlers.SendFileByURL).Methods(http.MethodPost, http.MethodOptions)

	// Static files
	s.router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	// Apply middleware
	s.router.Use(middleware.CORS)
	s.router.Use(handlers.CompressHandler)
}

func (s *Server) Start(addr string) error {
	logger.Info("Server starting on %s in %s mode", addr, s.config.Env)

	return http.ListenAndServe(addr, s.router)
}
