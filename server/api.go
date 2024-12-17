package server

import (
	"encoding/json"
	"net/http"

	"github.com/humamalamin/distributed-cache/cache"
)

type Server struct {
	cache *cache.Cache
}

func NewServer(c *cache.Cache) *Server {
	return &Server{cache: c}
}

func (s *Server) Start(port string) {
	http.HandleFunc("/get", s.getHandler)
	http.HandleFunc("/set", s.setHandler)
	http.ListenAndServe(":"+port, nil)
}

func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if value, ok := s.cache.Get(key); ok {
		json.NewEncoder(w).Encode(map[string]string{"value": value})
		return
	}
	http.Error(w, "key not found", http.StatusNotFound)
}

func (s *Server) setHandler(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	json.NewDecoder(r.Body).Decode(&req)
	s.cache.Set(req["key"], req["value"])
	w.WriteHeader(http.StatusOK)
}
