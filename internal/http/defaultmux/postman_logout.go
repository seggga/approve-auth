package defaultmux

import (
	"net/http"
)

// PostmanLogout clears users tokens
func (rt *Router) PostmanLogout(w http.ResponseWriter, r *http.Request) {

	rt.logger.Debug("postman logout handler call")

	clearCookies(w)
	rt.logger.Debug("erase cookies")

	redirect := r.URL.Query().Get("redirect_uri")
	if redirect != "" {
		rt.logger.Debugf("redirect request to %s", redirect)

		w.Header().Set("Location", redirect)
		w.WriteHeader(http.StatusFound)
	}
}
