package http

import (
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type Config struct {
	Addr string
	DSN  string
}

type Server struct {
	config Config
	db     *pgx.Conn
}

func New(cnf Config, db *pgx.Conn) *Server {
	return &Server{
		config: cnf,
		db:     db,
	}
}

func (s *Server) Mount() http.Handler {
	return NewRouter(s)
}

func (s *Server) Run(h http.Handler) error {
	svr := &http.Server{
		Addr:    s.config.Addr,
		Handler: h,
	}

	slog.Info("server started on", "port", s.config.Addr)

	return svr.ListenAndServe()
}
