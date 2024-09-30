package reader

import "github.com/Jacobo0312/microservice-latency-experiment/pkg/web"

func (h *handler) RouteURLs(app *web.Application) {

	app.Router.Get("/concurrency", h.WithConcurrency)

	app.Router.Get("/no-concurrency", h.WithoutConcurrency)

}
