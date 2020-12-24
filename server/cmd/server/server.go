package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/inovarka/lab3/server/balancers"
)

type HttpPortNumber int

// balancersApiServer configures necessary handlers and starts listening on a configured port.
type BalancersApiServer struct {
	Port HttpPortNumber

	BalancersHandler balancers.HttpHandlerFunc

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *BalancersApiServer) Start() error {
	if s.BalancersHandler == nil {
		return fmt.Errorf("bakancers HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/balancers", s.BalancersHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stops will shut down previously started HTTP server.
func (s *BalancersApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}