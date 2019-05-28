package open_closed_principle

import "net/http"

func healthCheckShort(resp http.ResponseWriter, _ *http.Request)  {
	resp.WriteHeader(http.StatusNoContent)
}

func headlthCheckShortUsage()  {
	http.Handle("/health", http.HandlerFunc(healthCheckShort))
}
