package udp

import (
	"net"
	"log"
)

var conns = map[string]*net.UDPConn{}
var routing = SafeRouting{up: make(map[string]string), down: make(map[string]map[string]bool)}

func dataLoop (conn_id string) {
	for _, ok := conns[conn_id]; ok; {
		conn := conns[conn_id]
		
		buf := make([]byte, 4096)
		n, _, err := conn.ReadFromUDP(buf)
		
		if err != nil {
			log.Print(err)
			continue
		}

		// Create a slice containing
		// only the actual packet by
		// using the length returned
		// by ReadFromUDP
		conc := buf[0:n]
		
		recvs := routing.Destinations(conn_id)
		
		for _, recv := range recvs {
			recv_conn, recv_ok := conns[recv];
			if recv_ok {
				recv_conn.Write(conc)
			}	
		}
	}
}

func CreateServer (port string) {
	if _, ok := conns[port]; ok {
		log.Print("[UDP] Port", port, "is already in use")
		return
	}
	
	addr, err := net.ResolveUDPAddr("udp", port)
	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		log.Print(err)
		return
	}

	conns[port] = conn
	dataLoop(port)
}

func CreateClient (addr string) {
	dst, err := net.ResolveUDPAddr("udp", string(addr))
	conn, _ := net.DialUDP("udp", nil, dst)
	
	if err != nil {
		log.Println(err)
		return
	}
	
	conns[addr] = conn
	dataLoop(addr)
}

func Kill (conn string) {
	if conns[conn] != nil {
		conns[conn].Close()
		conns[conn] = nil
	}
}

func Route (from string, to string) {
	routing.Route(from, to)
}