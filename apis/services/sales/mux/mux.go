// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"net/http"

	"github.com/evaldasg/ardanlabs-ultimate-go-service-k8s/apis/services/sales/route/sys/checkapi"
)

// WebAPIAuth constructs a http.Handler with all application routes bound.
func WebAPI() *http.ServeMux {
	mux := http.NewServeMux()

	checkapi.Routes(mux)

	return mux
}
