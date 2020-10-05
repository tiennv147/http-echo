package mux

import (
	"fmt"
	"net/http"
	"time"

	pb "github.com/tiennv147/http-echo/pb"
)

// New ...
func New(routes []*pb.Route) *http.ServeMux {
	// Flag gets printed as a page
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.HandleFunc(
			route.GetMatch().GetPath(),
			httpEcho(route),
		)
	}

	// Health endpoint
	mux.HandleFunc("/status", httpHealth())

	return mux
}

func httpEcho(route *pb.Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headers := route.GetResponseHeaders()
		body := route.GetResponseBody()
		delay := route.GetDelay()
		if delay != nil && delay.GetSeconds() > 0 {
			time.Sleep(delay.AsDuration())
		}
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
