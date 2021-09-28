package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", timeHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	q := r.URL.Query()
	m := make(map[string]interface{})

	m["time"] = time.Now().UTC()
	if tz := q.Get("tz"); tz != "" {
		if loc, err := time.LoadLocation(tz); err == nil {
			m["time"] = time.Now().In(loc)
		} else {
			m["error"] = "Unknown TimeZone"
			m["time"] = nil
		}
	}

	_ = json.NewEncoder(w).Encode(m)
}
