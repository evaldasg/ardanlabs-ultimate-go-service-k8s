package checkapi

import "github.com/evaldasg/ardanlabs-ultimate-go-service-k8s/foundation/web"

func Routes(app *web.App) {
	app.HandleFunc("GET /liveness", liveness)
	app.HandleFunc("GET /readiness", readiness)
	app.HandleFunc("GET /testerror", testError)
}
