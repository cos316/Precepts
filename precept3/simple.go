/*****************************************************************************
 * Description: A simple web server
 *
 * Usage:       ./simple [server-port]
 *
 * Example:     ./simple 8081
 *
 *
 *****************************************************************************/
package main

import (
        "os"
        "net/http"
        "log"
)

func main() {
        // get port number from command line
        if len(os.Args) != 2 {
                log.Fatal("Usage: ./simple [server port]")
        }
        server_port := os.Args[1]

        // set up web server
        _____________________
}
