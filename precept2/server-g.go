/*****************************************************************************
 * Name:
 * NetId:
 *
 * Description: A simple TCP server using the Go net API running on 
 *              a IPV4 IP address 127.0.0.1 and given port. 
 *
 * Usage:       ./server-g [server-port]
 *
 * Example:     ./server-g 8999
 *
 *
 *****************************************************************************/

package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

const RECV_BUFFER_SIZE = 2048

func main() {
	// get port number from command line
    __________

	// Create a TCP socket and bind to IP address 127.0.0.1 and port
	// Listen on socket for new connections
    __________

	// Create a writer for stdout
	writer := bufio.NewWriter(os.Stdout)

	// Loop waiting for connections
	for {
		// look for a client to connect
        ___________

        // create an input buffer
		message := make([]byte, RECV_BUFFER_SIZE)

		// read the data sent by the client
        ___________

        // write the data to stdout
		____________

        // clean up/close the connection
        ____________
	}
}
