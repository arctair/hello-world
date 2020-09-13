package v1

import (
	"net/http"
)

// Router ...
type Router struct {
	versionController *VersionController
}

// NewRouter ...
func NewRouter(
	versionController *VersionController,
) *Router {
	return &Router{
		versionController,
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	serveMux := http.NewServeMux()
	serveMux.Handle("/version", router.versionController.HandlerFunc())
	serveMux.ServeHTTP(w, request)
}
