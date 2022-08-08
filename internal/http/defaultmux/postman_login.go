package defaultmux

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/seggga/approve-auth/internal/entity"
	"github.com/seggga/approve-auth/internal/tokens"
)

// PostmanLogin handle func produces JWT and refresh token
func (rt *Router) PostmanLogin(w http.ResponseWriter, r *http.Request) {

	rt.logger.Debug("postman login call")

	ctxUser := r.Context().Value(ctxKeyUser{})
	user, ok := ctxUser.(*entity.UserOpts)
	if !ok {

		rt.logger.Errorw(fmt.Sprintf("cannot parse context to useropts struct, %v", ctxUser),
			"package", "defaultmux",
			"method", "PostmanLogin",
		)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rt.logger.Debugf("user %s passed", user.Login)

	// create token pair
	token, refresh, err := tokens.CreateTokenPair(user.Login, rt.jwtSecretKey)
	if err != nil {

		rt.logger.Errorw(fmt.Sprintf("error creating token pair: %v", err),
			"package", "defaultmux",
			"method", "PostmanLogin",
		)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rt.logger.Debug("token pair created sucessfuly")

	tokenPair := &TokenPair{
		AccessToken:  token,
		RefreshToken: refresh,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tokenPair); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	redirect := r.URL.Query().Get("redirect_uri")
	if redirect != "" {
		rt.logger.Debugf("redirect request to %s", redirect)

		w.Header().Set("Location", redirect)
		w.WriteHeader(http.StatusFound)
	}
}
