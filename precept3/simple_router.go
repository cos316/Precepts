/*****************************************************************************
 * Description: A simple router for a simple web server.
 *		Uses DeafultServeMux
 *
 *****************************************************************************/
package main

import (
        "net/http"
        "io"
)

func configure_routes() {
        // create a function for  each route: http://localhost:[port]/[routename]
        routeAHandler := ________________
        routeBHandler := ________________
	...


        // for each route pattern, register the handler
        http.HandleFunc(______________)
        http.HandleFunc(______________)

}
