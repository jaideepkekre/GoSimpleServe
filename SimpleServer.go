package SimpeServer

import (
	"log"
	"net/http"
)

// Plugs in Auth, supports JWT
func AuthMiddleware(h http.Handler, comment string) http.Handler {
	log.Println("Inflating AuthMiddleWare  for " + comment)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("AuthMiddleWare Invoked")
		h.ServeHTTP(w, r)
		// w.WriteHeader(400)
	})
}

//Inflates the Function to HTTP verb map
func ResourceHandler(muxer map[string]func(w http.ResponseWriter, r *http.Request), comment string) http.Handler {
	log.Println("Inflating ResourceHandler for " + comment)
	wrappedIndex := make(map[string]http.Handler)
	for method, fun := range muxer {
		wrappedIndex[method] = AuthMiddleware(http.HandlerFunc(fun), method)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("ResourceHandler Invoked")

		switch r.Method {
		case http.MethodGet:
			wrappedIndex[http.MethodGet].ServeHTTP(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	})

}

