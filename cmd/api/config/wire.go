//+build wireinject

package config

import (
	"github.com/codyoss/wired/pkg/client"
	"github.com/codyoss/wired/pkg/rpc/http"
	"github.com/google/wire"
	"github.com/gorilla/mux"
)

// InitializeRouter is a injector for Router. Used by the wire auto-injection framework to generate code.
func InitializeRouter() *mux.Router {
	wire.Build(http.NewRouter, http.NewSWAPIHandlerSet, CachedServiceSWAPIService, ServiceSWAPIService, client.New, ProvideClientOptions)
	return &mux.Router{}
}
