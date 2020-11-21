package rmrp

import (
	"errors"
	"strings"
)

type validator func ([]string) ([]string, error)

var verbs = map[string]validator{
	// ROUTE
	// Format:
	//	ROUTE <SOURCE> <DESTINATION>
	// Example:
	//	ROUTE :3001 localhost:3002 - Will start routing the data from a server on port 3001 
	//															 to a client connected to localhost:3002
	"ROUTE": func (parts []string) ([]string, error) {
		if len(parts) != 3 {
			return nil, errors.New("Malformed command. Should follow format \"ROUTE <input> <output>\"")
		}
		return parts, nil
	},
	
	// ADD
	// Format:
	//	ADD <SERVER|CLIENT> <addr>
	// Example:
	//	ADD SERVER :3001 					- Starts listening for data on localhost:3001
	//	ADD CLIENT localhost:3002 - Try to create a socket by connecting to localhost:3002
	"ADD": func (parts []string) ([]string, error) {
		if len(parts) != 3 {
			return nil, errors.New("Malformed command. Should follow format \"ADD <type> <addr>\"")
		}
		
		if parts[1] != "SERVER" && parts[1] != "CLIENT" {
			return nil, errors.New("Malformed command. Type must be either \"SERVER\" or \"CLIENT\"")
		}
		
		return parts, nil
	},
}

// Parse a string as a RMRP-command
func Parse (cmd string) ([]string, error) {
	clean := strings.TrimSuffix(cmd, "\n")
	parts := strings.Split(clean, " ")
	
	if validator, ok := verbs[parts[0]]; ok {
		return validator(parts)
	}
	
	err := errors.New("Invalid command")
	return nil, err
}