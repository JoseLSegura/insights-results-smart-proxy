package server

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type HTTPServer interface {
	GetConfiguration() Configuration
	SetServ(*http.Server)
	GetServ() *http.Server
	AddEndpointsToRouter(*mux.Router)
}

type HTTPServerImpl struct{}

// Start starts the server
func Start(server HTTPServer) error {
	address := server.GetConfiguration().Address
	log.Info().Msgf("Starting HTTP server at '%s'", address)
	router := initialize(server, address)
	httpServer := &http.Server{Addr: address, Handler: router}
	var err error

	if server.GetConfiguration().UseHTTPS {
		err = httpServer.ListenAndServeTLS("server.crt", "server.key")
	} else {
		err = httpServer.ListenAndServe()
	}
	if err != nil && err != http.ErrServerClosed {
		log.Error().Err(err).Msg("Unable to start HTTP/S server")
		return err
	}

	server.SetServ(httpServer)

	return nil
}

// Stop stops server's execution
func Stop(server HTTPServer, ctx context.Context) error {
	httpServer := server.GetServ()
	return httpServer.Shutdown(ctx)
}

// initialize perform the server initialization
func initialize(server HTTPServer, address string) http.Handler {
	log.Info().Msgf("Initializing HTTP server at '%s'", address)

	router := mux.NewRouter().StrictSlash(true)
	server.AddEndpointsToRouter(router)

	return router
}
