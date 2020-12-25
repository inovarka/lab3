package balancers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/inovarka/lab3/server/tools"
)

// HTTPHandlerFunc is Plants HTTP handler.
type HTTPHandlerFunc http.HandlerFunc

// HTTPHandler creates a new instance of plants HTTP handler.
func HTTPHandler(store *Store) HTTPHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListBalancers(store, rw)
		} else if r.Method == "PATCH" {
			handleMachineUpdate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListBalancers(store *Store, rw http.ResponseWriter) {
	res, err := store.ListBalancers()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}

func handlePlantUpdate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var machine Machine
	if err := json.NewDecoder(r.Body).Decode(&plant); err != nil {
		log.Printf("Error decoding machine input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.UpdateMachine(machine.ID, machine.isWorking)
	if err == nil {
		tools.WriteJsonOk(rw, &machine)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}
