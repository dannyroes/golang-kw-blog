/*
A fairly simple wrapper on the net/http package. This allows us to easily swap
in another web package if we need more robust functionality.
*/

package blog

import (
	"fmt"
	"net/http"
)

const (
	ListenAddr = ":8080"
)

//Define a generic function for routes and a generic response type
//These can wrap functionality from other libraries if net/http isn't
//powerful enough in the future
type RouteFunc func(params map[string][]string) WebResponse
type WebResponse struct {
	Body         string
	RedirectUrl  string
	RedirectCode int
}

//Add a route we'll listen for. WebResponse will either set the HTML of the response
//or cause us to redirect to another page.
func CreateRoute(route string, routeFunction RouteFunc) {
	http.HandleFunc(route, func(response http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		webResponse := routeFunction(request.Form)
		if webResponse.RedirectUrl != "" {
			http.Redirect(response, request, webResponse.RedirectUrl, webResponse.RedirectCode)
		} else {
			response.Write([]byte(webResponse.Body))
		}
	})
}

//Adds our routes and actually starts the webserver.
func Run() {
	addRoutes()
	fmt.Println("Starting webserver")
	http.ListenAndServe(ListenAddr, nil)
}
