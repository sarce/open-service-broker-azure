package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) poll(w http.ResponseWriter, r *http.Request) {
	instanceID := mux.Vars(r)["instance_id"]
	log.Println(instanceID)
	// TODO: Returns 200 or 410; also see spec for response body format
}