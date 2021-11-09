package rates

import (
	"context"
	stdErr "errors"
	"fmt"

	stdHttp "net/http"

	"template/internal/model"

	"gitlab.ar.bsch/santander-tecnologia/santander-go-kit/platform/errors"
	"gitlab.ar.bsch/santander-tecnologia/santander-go-kit/platform/transport/http"
)

var (
	badRequestMessage    = stdHttp.StatusText(stdHttp.StatusBadRequest)
	badRequestCodeString = fmt.Sprintf("%v", stdHttp.StatusBadRequest)
)

type MyEntityService interface {
	GetEntity(ctx context.Context, id string) (model.MyEntity, error)
}

type MyEntityController struct {
	Service MyEntityService
}

func (cc MyEntityController) GetEntity(w stdHttp.ResponseWriter, r *stdHttp.Request) error {
	ctx := r.Context()
	query := r.URL.Query()

	id := query.Get("id")
	if id == "" {
		apiErr := errors.NewWrappedAPIError(
			badRequestCodeString,
			badRequestMessage,
			stdErr.New("id is required"))
		return http.RespondError(w, badRequestCodeString, badRequestMessage, apiErr)
	}

	entity, err := cc.Service.GetEntity(ctx, id)
	if err != nil {
		return http.RespondError(
			w,
			fmt.Sprintf("%v", stdHttp.StatusInternalServerError),
			stdHttp.StatusText(stdHttp.StatusInternalServerError),
			err)
	}

	return http.Respond(w, stdHttp.StatusOK, entity)
}
