// Copyright 2020 Red Hat, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package server contains implementation of REST API server (HTTPServer) for the
// Insights Results Smart Proxy.
package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HTTPServer contains the configuration and objects to handle a HTTP server
type MyHTTPServer struct {
	Config Configuration
	Serv   *http.Server
}

// New constructs new implementation of MyHTTPServer
func New(conf Configuration) HTTPServer {
	return &MyHTTPServer{
		Config: conf,
	}
}

func (server *MyHTTPServer) AddEndpointsToRouter(router *mux.Router) {
	fmt.Println("Configuring endpoints")
	// apiPrefix := server.Config.APIPrefix
	// openAPIURL := apiPrefix + filepath.Base(server.Config.APISpecFile)

	// common REST API endpoints
	// router.HandleFunc(apiPrefix+MainEndpoint, server.mainEndpoint).Methods(http.MethodGet)

	// OpenAPI specs
	// router.HandleFunc(openAPIURL, server.serveAPISpecFile).Methods(http.MethodGet)
}

func (server *MyHTTPServer) GetConfiguration() Configuration {
	return server.Config
}

func (server *MyHTTPServer) GetServ() *http.Server {
	return server.Serv
}

func (server *MyHTTPServer) SetServ(httpServer *http.Server) {
	server.Serv = httpServer
}
