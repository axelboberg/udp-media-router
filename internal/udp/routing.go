package udp
import (
	"sync"
)

type SafeRouting struct {
	mu sync.Mutex
	up map[string]string
	down map[string]map[string]bool
}

// Get the keys from a map
// as a slice of strings
func keys (m map[string]bool) []string {
	keys := make([]string, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Insert a new routing record safely
// by making sure that a destination
// can only have one source at a time
func (r *SafeRouting) Route (from string, to string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if _, ok := r.up[to]; ok {
		prev_from := r.up[to]
		delete(r.up, to)
		
		if _, ok := r.down[prev_from]; ok {
			delete(r.down[prev_from], to)
		}
	}
	
	if _, ok := r.down[from]; !ok {
		r.down[from] = make(map[string]bool)
	}
	
	r.down[from][to] = true
	r.up[to] = from
}

// Get all destinations for a
// source as a slice of strings
func (r *SafeRouting) Destinations (host string) []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return keys(r.down[host])
}