package mux

import (
	"fmt"
	"net/http"

	pb "github.com/tiennv147/http-echo/pb"
)

// New ...
func New(routes []*pb.Route) *http.ServeMux {
	// Flag gets printed as a page
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.HandleFunc(
			route.GetMatch().GetPath(),
			httpEcho(route.GetResponseHeaders(), route.GetResponseBody()),
		)
	}

	// Health endpoint
	mux.HandleFunc("/status", httpHealth())

	return mux
}

func httpEcho(headers map[string]string, body string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for k, v := range headers {
			w.Header().Add(k, v)
		}
		fmt.Fprintln(w, body)
	}
}

func httpHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
