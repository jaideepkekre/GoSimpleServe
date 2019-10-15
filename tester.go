//make package name main and copy this file out then test
package SimpeServer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jaideepkekre/GoSimpleServer"
)

func sample(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		var f interface{}
		body, _ := ioutil.ReadAll(r.Body)
		_ = json.Unmarshal(body, &f)
		m := f.(map[string]interface{})
		sm, _ := json.Marshal(m)
		log.Println("In Sample", sm)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(203)
		w.Write(body)

	}

}

func main() {
	index := make(map[string]func(w http.ResponseWriter, r *http.Request))
	index[http.MethodGet] = Sample
	http.Handle("/sample", SimpeServer.ResourceHandler(index, "Sample Endpoint"))
	http.ListenAndServe(":8092", nil)
}
