// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/evaldasg/ardanlabs-ultimate-go-service-k8s/apis/services/sales/route/sys/checkapi"
	"github.com/evaldasg/ardanlabs-ultimate-go-service-k8s/foundation/web"
)

// WebAPIAuth constructs a http.Handler with all application routes bound.
func WebAPI(shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown)

	checkapi.Routes(mux)

	return mux
}
