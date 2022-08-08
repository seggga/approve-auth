package defaultmux

import (
	"fmt"
	"net/http"
)

// ProfilerOn makes it possible to use /debug/pprof/* endpoints
// to get application profile
func (rt *Router) ProfilerOn(w http.ResponseWriter, r *http.Request) {

	rt.logger.Infow("profiler switched on",
		"package", "defaultmux",
	)

	rt.profilerOff = false
	fmt.Fprint(w, "http-profiler is turned on")
}

// ProfilerOff makes it impossible to use /debug/pprof/* endpoints
func (rt *Router) ProfilerOff(w http.ResponseWriter, r *http.Request) {

	rt.logger.Infow("profiler switched off",
		"package", "defaultmux",
	)

	rt.profilerOff = true
	fmt.Fprint(w, "http-profiler is turned off")
}

func (rt *Router) checkProfilerSwitch(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if rt.profilerOff {
			http.Error(w, "profiler is switched off", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
