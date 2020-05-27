/*
Copyright © 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Entry point to the insights results smart proxy
package main

import (
	"flag"
	"os"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/RedHatInsights/insights-results-smart-proxy/server"
)

const (
	// ExitStatusOK means that the tool finished with success
	ExitStatusOK = 0

	// ExitStatusServerError means that the HTTP server cannot be initialized
	ExitStatusServerError = 1
)

var serverInstance server.HTTPServer

// startService starts service and returns error code
func startService() int {
	serverCfg := server.Configuration{
		Address: ":8080",
	}
	serverInstance = server.New(serverCfg)

	err := server.Start(serverInstance)
	if err != nil {
		log.Error().Err(err).Msg("HTTP(s) start error")
		return ExitStatusServerError
	}

	return ExitStatusOK
}

// handleCommand select the function to be called depending on command argument
func handleCommand(command string) int {
	switch command {
	case "start-service":
		return startService()

	case "print-version":
		printVersionInfo()
		return ExitStatusOK
	}

	return ExitStatusOK
}

func main() {
	var args []string
	flag.Parse()

	args = flag.Args()

	command := "start-service"
	if len(args) >= 1 {
		command = strings.ToLower(strings.TrimSpace(args[0]))
	}

	os.Exit(handleCommand(command))
}
