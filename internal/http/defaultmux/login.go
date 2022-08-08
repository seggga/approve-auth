package defaultmux

import (
	"fmt"
	"net/http"
	"time"

	"github.com/seggga/approve-auth/internal/entity"
	"github.com/seggga/approve-auth/internal/tokens"
)

// Login handle func produces JWT and refresh token
func (rt *Router) Login(w http.ResponseWriter, r *http.Request) {

	rt.logger.Debug("login handler call")

	ctxUser := r.Context().Value(ctxKeyUser{})
	user, ok := ctxUser.(*entity.UserOpts)
	if !ok {

		rt.logger.Errorw(fmt.Sprintf("cannot parse context.Value to useropts struct, %v", ctxUser),
			"package", "defaultmux",
			"method", "Login",
		)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rt.logger.Debug("user successfuly parsed from r.context")

	// create token pair
	access, refresh, err := tokens.CreateTokenPair(user.Login, rt.jwtSecretKey)
	if err != nil {

		rt.logger.Errorw(fmt.Sprintf("error creating token pair: %v", err),
			"package", "defaultmux",
			"method", "Login",
		)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rt.logger.Debug("token pair has been created")

	clearCookies(w)
	rt.logger.Debug("erase cookies")
	sendCookies(access, refresh, w)
	rt.logger.Debug("set new cookies")

	redirect := r.URL.Query().Get("redirect_uri")
	if redirect != "" {

		rt.logger.Debugf("redirect request to %s", redirect)

		w.Header().Set("Location", redirect)
		w.WriteHeader(http.StatusFound)
	}
}

func sendCookies(access, refresh string, w http.ResponseWriter) {
	// set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "access",
		Value:    access,
		Expires:  time.Now().Add(time.Minute * 1),
		HttpOnly: true,
		Path:     "/",
	})

	// set refresh cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh",
		Value:    refresh,
		Expires:  time.Now().Add(time.Minute * 60),
		HttpOnly: true,
		Path:     "/",
	})
}
