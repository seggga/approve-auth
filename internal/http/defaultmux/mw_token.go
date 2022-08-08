package defaultmux

import (
	"context"
	"fmt"
	"net/http"

	"github.com/seggga/approve-auth/internal/tokens"
)

// TokenMW ...
func (rt *Router) TokenMW(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			rt.logger.Debug("token middleware call")

			accessCookie, err1 := r.Cookie("access")
			refreshCookie, err2 := r.Cookie("refresh")

			if err1 != nil && err2 != nil {

				rt.logger.Debug("no token passed with cookies")

				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				http.Error(w, "{}", http.StatusForbidden)
				return
			}

			var (
				valid bool
				login string
			)
			// check access token
			if accessCookie != nil {
				rt.logger.Debug("got access token")

				valid, login = tokens.CheckToken(accessCookie.Value, rt.jwtSecretKey)
			}

			if valid {
				rt.logger.Debug("access token is valid")

				user, err := rt.stor.ReadUser(login)
				if err != nil {

					rt.logger.Warnw(fmt.Sprintf("with valid access token error getting user data %s, %v", login, err),
						"package", "defaultmux",
						"method", "TokenMW",
					)

					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				rt.logger.Debugf("user %s authenticated", user.Login)

				ctx := context.WithValue(r.Context(), ctxKeyUser{}, user)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			rt.logger.Debug("access token is not valid or was not passed")

			// check refresh token
			if refreshCookie != nil {
				rt.logger.Debug("got refresh token")

				valid, login = tokens.CheckToken(refreshCookie.Value, rt.jwtSecretKey)
			}

			if valid {
				rt.logger.Debug("refresh token is valid")

				user, err := rt.stor.ReadUser(login)
				if err != nil {

					rt.logger.Warnw(fmt.Sprintf("with valid refresh token error getting user data %s, %v", login, err),
						"package", "defaultmux",
						"method", "TokenMW",
					)

					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				rt.logger.Debugf("user %s authenticated", user.Login)

				// create token pair
				access, refresh, err := tokens.CreateTokenPair(login, rt.jwtSecretKey)
				if err != nil {

					rt.logger.Errorw(fmt.Sprintf("with valid refresh token error creating token pair: %v", err),
						"package", "defaultmux",
						"method", "TokenMW",
					)

					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				rt.logger.Debug("new token pari is created")

				clearCookies(w)
				rt.logger.Debug("erase cookies")
				sendCookies(access, refresh, w)
				rt.logger.Debug("set new cookies")

				ctx := context.WithValue(r.Context(), ctxKeyUser{}, user)
				next.ServeHTTP(w, r.WithContext(ctx))
				return

			}

			w.WriteHeader(http.StatusUnauthorized)
			return

		},
	)
}
