# UDP Media Router
Route UDP data between sockets in realtime

## Table of contents
- [RMRP](#rmrp)
	- [ADD](#add)
	- [REMOVE](#remove)
	- [ROUTE](#route)

## RMRP
**Realtime Media Routing Protocol**  
The router is listening for commands as text sent directly over TCP to port 8848 by default.

### ADD
`ADD <SERVER|CLIENT> <addr>`  

**Example:**  
Listen to packets on port 3001  
`ADD SERVER :3001`  

Try to connect to an open port on localhost  
`ADD CLIENT localhost:3002`

### REMOVE  
`REMOVE <addr>`  
Removes any connections for the specified address and stops forwarding/receiving data to it. Essentially the opposite of `ADD`.

**Example:**  
Remove a connection for address `localhost:3002`  
`REMOVE localhost:3002`  

Remove a connection for address `:3001`  
`REMOVE :3001`

### ROUTE
`ROUTE <source> <destination>`

**Note:**  
Both the source and destination must have been created with `ADD` before trying to route media between them.

**Example:**  
Route data from a server listening on port 3001 and a client connected to `localhost:3002`  
`ROUTE :3001 localhost:3002`