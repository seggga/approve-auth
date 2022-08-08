package defaultmux

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// PostmanAuthMW ...
func (rt *Router) PostmanAuthMW(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			rt.logger.Debug("Authentication with json in body call")

			// get credentials from request's body
			// check method
			if r.Method != http.MethodPost {

				rt.logger.Debugf("wrong method: got %s, expected POST", r.Method)

				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}

			// copy r.Body into body to read it's content twice:
			// in this middleware function and in the next handler
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {

				rt.logger.Debugf("Error reading body, %v", err)

				http.Error(w, "can't read body", http.StatusBadRequest)
				return
			}
			// restore request's Body
			r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			// read passed credentials (from the body)
			var creds Credentials
			err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&creds)
			if err != nil {

				rt.logger.Debugf("error parsing credentials from body, %v", err)

				http.Error(w, "{}", http.StatusInternalServerError)
				return
			}
			// find password hash for specified user
			user, err := rt.stor.ReadUser(creds.Username)
			if err != nil {

				rt.logger.Debugf("error getting user data, %v", err)

				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			passHash := user.PassHash

			// compare password and hash
			// if !checkPasswordHash(creds.Password, passHash) {
			if !checkPasswordHash(creds.Password, passHash) {

				rt.logger.Debugf("wrong password for user %s", creds.Username)

				http.Error(w, "{}", http.StatusForbidden)
				return
			}

			rt.logger.Debugf("user %s authenticated", creds.Username)

			ctx := context.WithValue(r.Context(), ctxKeyUser{}, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
