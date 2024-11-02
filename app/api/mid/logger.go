package mid

import (
	"context"
	"net/http"

	"github.com/evaldasg/ardanlabs-ultimate-go-service-k8s/foundation/logger"
	"github.com/evaldasg/ardanlabs-ultimate-go-service-k8s/foundation/web"
)

func Logger(log *logger.Logger) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			log.Info(ctx, "request started", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)
			// Start
			err := handler(ctx, w, r)
			// Finish
			log.Info(ctx, "request finished", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			return err
		}

		return h
	}

	return m
}
