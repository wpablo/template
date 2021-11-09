package server

import (
	"context"

	"gitlab.ar.bsch/santander-tecnologia/santander-go-kit/platform/config"

	"github.com/go-chi/chi"

	"template/internal/service"

	"template/internal/controller/rates"

	santanderCtx "gitlab.ar.bsch/santander-tecnologia/santander-go-kit/platform/context"
	contextMiddleware "gitlab.ar.bsch/santander-tecnologia/santander-go-kit/platform/middleware/context"
	"gitlab.ar.bsch/santander-tecnologia/santander-go-kit/platform/transport/http/circuitbreaker"
)

func setupHandlers(router chi.Router, cfg config.Configuration) {
	myExampleHeaderKey := santanderCtx.Key("X-My-Header")
	ratesCc := ratesController()
	router.Route("/myentity", func(r chi.Router) {
		r.Use(contextMiddleware.PropagateHeaders(map[santanderCtx.Key]santanderCtx.Setter{
			myExampleHeaderKey: func(ctx context.Context, val string) context.Context {
				return context.WithValue(ctx, myExampleHeaderKey, val)
			},
		}))
		r.Get("/", circuitbreaker.Apply(ratesCc.GetEntity))
	})

}

// configuro el servicio que es invocado por el controlador
func ratesController() rates.MyEntityController {
	ratesSvc := service.NewMyService()
	return rates.MyEntityController{
		Service: ratesSvc,
	}
}
