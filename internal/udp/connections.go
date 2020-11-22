package udp

import (
	"net"
	"sync"
)

type Connections struct {
	mu sync.Mutex
	conns map[string]*net.UDPConn
}

// Add a reference to a
// connection by its address
func (cs *Connections) Add (addr string, conn *net.UDPConn) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	
	cs.conns[addr] = conn
}

// Remove the reference to a
// connection by its address
func (cs *Connections) Remove (addr string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	
	delete(cs.conns, addr)
}

// Get a reference to a
// connection by its address
func (cs *Connections) Get (addr string) (*net.UDPConn, bool) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	
	conn, ok := cs.conns[addr]
	return conn, ok
}