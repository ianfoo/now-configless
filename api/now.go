package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Now(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	resp := map[string]time.Time{"now": t}
	json.NewEncoder(w).Encode(resp)
	log.Printf("served request for %s in %v", r.RemoteAddr, time.Since(t))
}
