package defaultmux

import (
	"context"
	"fmt"
	"net/http"
)

// AuthMiddleware implements password check middleware
func (rt *Router) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			rt.logger.Debug("auth middleware call")

			// check basic-auth
			username, password, ok := r.BasicAuth()
			if !ok {

				rt.logger.Debugw("no credentials passed",
					"package", "defaultmux",
					"method", "AuthMiddleware",
				)

				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// find password hash for specified user
			user, err := rt.stor.ReadUser(username)
			if err != nil {

				rt.logger.Debugw(fmt.Sprintf("error getting user data, %v", err),
					"package", "defaultmux",
					"method", "AuthMiddleware",
				)

				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			passHash := user.PassHash

			// compare password and hash
			// if !checkPasswordHash(creds.Password, passHash) {
			if !checkPasswordHash(password, passHash) {

				rt.logger.Debug(fmt.Sprintf("wrong password for user %s with ID %s", user.Login, user.ID.String()))

				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			rt.logger.Debug(fmt.Sprintf("authentication passed with user %s, ID %s", user.Login, user.ID.String()))

			ctx := context.WithValue(r.Context(), ctxKeyUser{}, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
