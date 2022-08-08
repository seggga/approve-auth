package defaultmux

import (
	"net/http"
	"net/http/pprof"

	"github.com/seggga/approve-auth/internal/storage"
	"go.uber.org/zap"
)

// a struct as a key to pass User through request context
type ctxKeyUser struct{}

// Router ...
type Router struct {
	*http.ServeMux
	stor         storage.UserStorage
	jwtSecretKey string
	logger       *zap.SugaredLogger
	profilerOff  bool
}

// New creates a new router
func New(stor storage.UserStorage, jwtSecretKey string, slog *zap.SugaredLogger) *Router {
	rt := &Router{
		ServeMux:     http.NewServeMux(),
		stor:         stor,
		jwtSecretKey: jwtSecretKey,
		logger:       slog,
		profilerOff:  true,
	}

	rt.Handle("/login",
		rt.AuthMiddleware(
			http.HandlerFunc(rt.Login),
		),
	)

	rt.Handle("/logout",
		rt.AuthMiddleware(
			http.HandlerFunc(rt.Logout),
		),
	)

	rt.Handle("/i",
		rt.TokenMW(
			http.HandlerFunc(rt.Info),
		),
	)

	// to fit given postman collection
	rt.Handle("/auth/v1/login",
		rt.PostmanAuthMW(
			http.HandlerFunc(rt.PostmanLogin),
		),
	)

	rt.Handle("/auth/v1/logout",
		http.HandlerFunc(rt.PostmanLogout),
	)

	rt.Handle("/auth/v1/validate",
		rt.TokenMW(
			http.HandlerFunc(rt.Info),
		),
	)

	// http profiler: on/off
	rt.HandleFunc("/profiler-on", rt.ProfilerOn)
	rt.HandleFunc("/profiler-off", rt.ProfilerOff)

	// http profiler: program's command line
	rt.Handle("/debug/pprof/cmdline",
		rt.checkProfilerSwitch(
			http.HandlerFunc(pprof.Cmdline),
		),
	)
	// http profiler: CPU profile
	rt.Handle("/debug/pprof/profile",
		rt.checkProfilerSwitch(
			http.HandlerFunc(pprof.Profile),
		),
	)
	// http profiler: symbol
	rt.Handle("/debug/pprof/symbol",
		rt.checkProfilerSwitch(
			http.HandlerFunc(pprof.Symbol),
		),
	)
	// http profiler: execution trace
	rt.Handle("/debug/pprof/trace",
		rt.checkProfilerSwitch(
			http.HandlerFunc(pprof.Trace),
		),
	)
	// http profiler: list available profiles
	rt.Handle("/debug/pprof/",
		rt.checkProfilerSwitch(
			http.HandlerFunc(pprof.Index),
		),
	)

	// redirect to /debug/pprof/
	rt.HandleFunc("/debug/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI+"/pprof/", http.StatusMovedPermanently)
	})

	return rt
}
