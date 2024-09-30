package reader

import (
	"net/http"

	"github.com/Jacobo0312/microservice-latency-experiment/pkg/errors"
	"github.com/Jacobo0312/microservice-latency-experiment/pkg/web"
)

func (h *handler) WithoutConcurrency(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := HomeParams{
		UserID: "1",
	}

	data, err := h.HomeService.GetSections(ctx, params.ToDomain())

	if err != nil {
		web.RespondWithError(w, errors.NewInternalServerError("error getting sections", err))
	}

	web.RespondWithJSON(w, http.StatusOK, data)

}
