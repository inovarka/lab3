// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/inovarka/lab3/server/balancers"
)

// Injectors from modules.go:

func ComposeApiServer(port HttpPortNumber) (*balancersApiServer, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	store := channels.NewStore(db)
	httpHandlerFunc := balancers.HttpHandler(store)
	balancersApiServer := &balancersApiServer{
		Port:            port,
		BalancersHandler: httpHandlerFunc,
	}
	return balancersApiServer, nil
}