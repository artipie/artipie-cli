package artipie

import "net/http"

// Auth sets authentification header to request.
type Auth interface {
	SetAuthHeader(req *http.Request)
}
