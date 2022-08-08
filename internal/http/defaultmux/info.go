package defaultmux

import (
	"fmt"
	"net/http"

	"github.com/seggga/approve-auth/internal/entity"
)

// Info checks token and sends username as response
func (rt *Router) Info(w http.ResponseWriter, r *http.Request) {

	rt.logger.Debug("info handler call")

	ctxUser := r.Context().Value(ctxKeyUser{})
	user, ok := ctxUser.(*entity.UserOpts)
	if !ok {

		rt.logger.Errorw(fmt.Sprintf("cannot parse context.Value to useropts struct, %v", ctxUser),
			"package", "defaultmux",
			"method", "Whoami",
		)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rt.logger.Debug("send userdata %s", user.Login)

	fmt.Fprintf(w, "hello, %s\n", user.Login)
	return
}
