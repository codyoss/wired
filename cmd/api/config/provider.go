package config

import (
	"github.com/codyoss/wired/pkg/client"
	"github.com/codyoss/wired/pkg/rpc/http"
	"github.com/codyoss/wired/pkg/swapi"
	"github.com/google/wire"
)

// CachedServiceSWAPIService binds swapi.CachedService to the http.SWAPIService interface
var CachedServiceSWAPIService = wire.NewSet(swapi.NewCachedService, wire.Bind(new(http.SWAPIService), new(swapi.CachedService)))

// ServiceSWAPIService binds swapi.Service to the swapi.SWAPIService interface
var ServiceSWAPIService = wire.NewSet(swapi.NewService, wire.Bind(new(swapi.SWAPIService), new(swapi.Service)))

// ProvideClientOptions sets Options used for client.
func ProvideClientOptions() []client.Option { return nil }
