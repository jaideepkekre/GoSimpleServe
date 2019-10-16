// package main

package SimpleServer

import (
	"log"
	"net/http"
)

type EndPointMux struct {
	URI    string
	GET    func(w http.ResponseWriter, r *http.Request)
	PUT    func(w http.ResponseWriter, r *http.Request)
	POST   func(w http.ResponseWriter, r *http.Request)
	DELETE func(w http.ResponseWriter, r *http.Request)
	PATCH  func(w http.ResponseWriter, r *http.Request)
}

// Plugs in Auth, supports JWT
func authMiddleware(h http.Handler) http.Handler {
	log.Println("Inflating AuthMiddleWare")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		// w.WriteHeader(400)
	})
}

//Inflates the Function to HTTP verb map
func (epm EndPointMux) ResourceHandler() http.Handler {
	log.Println("Inflating ResourceHandler")
	wrappedIndex := make(map[string]http.Handler)

	if epm.GET != nil {
		log.Println("GETTER INFLATED")
		wrappedIndex[http.MethodGet] = authMiddleware(http.HandlerFunc(epm.GET))
	}

	if epm.POST != nil {
		log.Println("POSTER INFLATED")
		wrappedIndex[http.MethodPost] = authMiddleware(http.HandlerFunc(epm.POST))
	}

	if epm.PUT != nil {
		log.Println("PUTTER INFLATED")
		wrappedIndex[http.MethodPut] = authMiddleware(http.HandlerFunc(epm.PUT))
	}

	if epm.PATCH != nil {
		log.Println("PATCH INFLATED")
		wrappedIndex[http.MethodPatch] = authMiddleware(http.HandlerFunc(epm.PATCH))
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("ResourceHandler Invoked for " + r.Method)

		switch r.Method {
		case http.MethodGet:
			if epm.GET != nil {
				wrappedIndex[http.MethodGet].ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}

		case http.MethodPost:
			if epm.POST != nil {
				wrappedIndex[http.MethodPost].ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}

		case http.MethodPut:
			if epm.PUT != nil {
				wrappedIndex[http.MethodPut].ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		case http.MethodDelete:
			if epm.DELETE!= nil {
				wrappedIndex[http.MethodDelete].ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		case http.MethodPatch:
			if epm.PATCH!= nil {
				wrappedIndex[http.MethodPatch].ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	})

}

func CreateEndpoint(uri string) EndPointMux {
	EPM := EndPointMux{URI:uri}
	return EPM
}

// func sample(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(203)
// }

// func main() {
// 	mux := http.NewServeMux()

// 	epm := CreateEndpoint("/sample")
// 	epm2 := CreateEndpoint("/sample/kekre")

// 	epm.GET = sample
// 	epm2.POST = sample
	
// 	mux.Handle(epm.URI, epm.ResourceHandler())
// 	mux.Handle(epm2.URI, epm.ResourceHandler())
	
// 	http.ListenAndServe(":8092", mux)
// }

