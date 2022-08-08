package defaultmux

import (
	"net/http"
	"time"
)

// Logout clears users tokens
func (rt *Router) Logout(w http.ResponseWriter, r *http.Request) {

	rt.logger.Debug("logout handler call")

	clearCookies(w)
	rt.logger.Debug("erase cookies")

	redirect := r.URL.Query().Get("redirect_uri")
	if redirect != "" {

		rt.logger.Debugf("redirect request to %s", redirect)

		w.Header().Set("Location", redirect)
		w.WriteHeader(http.StatusFound)
	}

}

func clearCookies(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "access",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:    "refresh",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	})
}
