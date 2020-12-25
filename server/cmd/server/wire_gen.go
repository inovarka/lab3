// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/inovarka/lab3/server/balancers"
)

// Injectors from modules.go:

// ComposeApiServer will create an instance of CharApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*BalancerApiServer, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	store := balancers.NewStore(db)
	httpHandlerFunc := balancers.HTTPHandler(store)
	balancerApiServer := &BalancerApiServer{
		Port:             port,
		BalancersHandler: httpHandlerFunc,
	}
	return balancerApiServer, nil
}