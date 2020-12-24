//+build wireinject

package main

import (
	"github.com/google/wire"

	"ggithub.com/inovarka/lab3/server/balancers"
)

// ComposeApiServer will create an instance of CharApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*BalancersApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from channels package.
		balancers.Providers,
		// Provide BalancersApiServer instantiating the structure and injecting channels handler and port number.
		wire.Struct(new(BalancersApiServer), "Port", "BalancersHandler"),
	)
	return nil, nil
}