package open_closed_principle

import "net/http"

// a HTTP healthy check handler in long form
type healthCheckLong struct {}

func (h *healthCheckLong)ServeHTTP(resp http.ResponseWriter, _ *http.Request)  {
	resp.WriteHeader(http.StatusNoContent)
}

func healthCheckLongUsage()  {
	http.Handle("/health", &healthCheckLong{})
}