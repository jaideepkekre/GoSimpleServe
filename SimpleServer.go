// package main
package SimpleServer

import (
	"log"
	"net/http"
)

// Plugs in Auth, supports JWT
func AuthMiddleware(h http.Handler, comment string) http.Handler {
	log.Println("Inflating AuthMiddleWare  for " + comment)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		log.Println("ResourceHandler Invoked for " + r.Method + " on " + comment)

		switch r.Method {
		case http.MethodGet:
			wrappedIndex[http.MethodGet].ServeHTTP(w, r)
		case http.MethodPost:
			wrappedIndex[http.MethodPost].ServeHTTP(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	})

}

func MakeMap() map[string]func(w http.ResponseWriter, r *http.Request) {
	return make(map[string]func(w http.ResponseWriter, r *http.Request))
}

func GET(muxer map[string]func(w http.ResponseWriter, r *http.Request), f func(w http.ResponseWriter, r *http.Request)) map[string]func(w http.ResponseWriter, r *http.Request) {
	muxer[http.MethodGet] = f
	return muxer
}

func POST(muxer map[string]func(w http.ResponseWriter, r *http.Request), f func(w http.ResponseWriter, r *http.Request)) map[string]func(w http.ResponseWriter, r *http.Request) {
	muxer[http.MethodPost] = f
	return muxer
}

// func sample(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(203)
// }

// func main() {
// 	index := MakeMap()
// 	index = GET(index, sample)
// 	index = POST(index, sample)
// 	http.Handle("/sample", ResourceHandler(index, "Sample Endpoint"))
// 	http.ListenAndServe(":8092", nil)
// }

