package rest

import (
	"github.com/Jacobo0312/microservice-latency-experiment/pkg/web"
)

type Handler interface {
	RouteURLs(app *web.Application)
}

// API constructs a http.Handler with all application routes defined.

func API(handlers ...Handler) *web.Application {
	app := web.NewWebApplication()

	for _, handler := range handlers {
		handler.RouteURLs(app)
	}

	return app
}
