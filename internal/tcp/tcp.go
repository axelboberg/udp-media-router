package tcp

import (
	"io"
	"net"
	"log"
	"bufio"
)

type listener func (io.Reader)

// Try to read the data from an io-reader
// as a string until the next linebreak
func ReadString (listener io.Reader) (string, error) {
	str, err := bufio.NewReader(listener).ReadString('\n')
	if err != nil {
		return "", err
	}
	return str, nil
}

// Create a TCP-server listening on the specified port.
// The port should be formatted as :<port>.
// Example:
//	:3001
func CreateServer (port string, listener listener) {
	l, err := net.Listen("tcp4", port)
	
	if err != nil {
		log.Print(err)
		return
	}
	
	log.Println("[TCP] Listening on port", port)
	
	defer l.Close()
	
	for {
		c, err := l.Accept()
		if err != nil {
			log.Print(err)
			return
		}
		go listener(c)
	}
}