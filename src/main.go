package main

import (
	// "fmt"
)
import (
    "github.com/emicklei/go-restful"
    "net/http"
    "log"
    "routes"
    "libs/auth"
    // "filter"
)


func main() {
    var routes routes.Routes
	wsContainer := restful.NewContainer()
    wsContainer.Add(routes.IndexRouter())
    wsContainer.Add(routes.ServiceRouter())
    wsContainer.Add(routes.ServiceRoutingRouter())
    wsContainer.Add(routes.ReplicationControllerRouter())
    wsContainer.Add(routes.PodRouter())
    wsContainer.Add(routes.NodeRouter())
    wsContainer.Add(routes.EventRouter())
    wsContainer.Add(routes.DockerRouter())
    wsContainer.Add(routes.LineRouter())
    wsContainer.Add(routes.BubbleRouter())
    
    var auth auth.Auth
    wsContainer.Add(auth.TokenRouter())
    wsContainer.Add(auth.UserRouter())


    // Add container filter to enable CORS
    cors := restful.CrossOriginResourceSharing{
        ExposeHeaders:  []string{"X-My-Header"},
        AllowedHeaders: []string{"Content-Type", "Accept", "token"},
        CookiesAllowed: false,
        Container:      wsContainer}
    wsContainer.Filter(cors.Filter)

    // Add container filter to respond to OPTIONS
    wsContainer.Filter(wsContainer.OPTIONSFilter)

    // var filter filter.Filter
    // wsContainer.Filter(filter.AuthToken)
	
    server := &http.Server{Addr:":8802", Handler: wsContainer}
    log.Fatal(server.ListenAndServe())
    // log.Fatal(http.ListenAndServe(":31001", nil))
}


