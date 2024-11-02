// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/evaldasg/ardanlabs-ultimate-go-service-k8s/apis/services/api/mid"
	"github.com/evaldasg/ardanlabs-ultimate-go-service-k8s/apis/services/sales/route/sys/checkapi"
	"github.com/evaldasg/ardanlabs-ultimate-go-service-k8s/foundation/logger"
	"github.com/evaldasg/ardanlabs-ultimate-go-service-k8s/foundation/web"
)

// WebAPIAuth constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger, shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log))

	checkapi.Routes(mux)

	return mux
}
