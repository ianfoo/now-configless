package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Date(rw http.ResponseWriter, r *http.Request) {
	t := time.Now()
	resp := struct {
		Now time.Time
	}{t}
	enc := json.NewEncoder(rw)
	enc.Encode(resp)
	log.Printf("served request for 'date' %s in %v", r.RemoteAddr, time.Since(t))
}
