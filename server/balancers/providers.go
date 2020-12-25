package balancers

import "github.com/google/wire"

// Providers ist of providers for balancers components.
var Providers = wire.NewSet(NewStore, HTTPHandler)
