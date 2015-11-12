package main

import (
	// "fmt"
)
import (
    "github.com/emicklei/go-restful"
    "net/http"
    "log"


	"controllers"
	// "./services"
)


func main() {
	// fmt.Printf("study test")

	// listService := services.ListService{}
	// data := listService.Service()

	// listService.Cluster()
	// data := listService.Cluster()

	// fmt.Println("====================================")
	// fmt.Println(data)
	
	
    // restful.Add(UserRouter())
    restful.Add(ServiceRouter())
    restful.Add(ReplicationControllerRouter())
    restful.Add(PodRouter())



    // log.Fatal(http.ListenAndServe(":9090", nil))
    log.Fatal(http.ListenAndServe(":31001", nil))
}


func ServiceRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/service").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)

	// listService := controllers.ServiceController{}
	var listService controllers.ServiceController

    router.Route(router.GET("/{name}").To(listService.Get))
    router.Route(router.GET("/").To(listService.GetList))
    router.Route(router.POST("/").To(listService.Post))
    router.Route(router.PUT("/{name}").To(listService.Put))
    router.Route(router.DELETE("/{name}").To(listService.Delete))
 
    return router
}

func ReplicationControllerRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/cluster").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)

	// listReplicationController := controllers.ReplicationControllerController{}
	var listReplicationController controllers.ReplicationControllerController

    router.Route(router.GET("/{name}").To(listReplicationController.Get))
    router.Route(router.GET("/").To(listReplicationController.GetList))
    router.Route(router.POST("/").To(listReplicationController.Post))
    router.Route(router.PUT("/{name}").To(listReplicationController.Put))
    router.Route(router.DELETE("/{name}").To(listReplicationController.Delete))
 
    return router
}

func PodRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/container").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)

	// listPod := controllers.PodController{}
	var listPod controllers.PodController

    router.Route(router.GET("/{name}").To(listPod.Get))
    router.Route(router.GET("/").To(listPod.GetList))
    router.Route(router.POST("/").To(listPod.Post))
    router.Route(router.PUT("/{name}").To(listPod.Put))
    router.Route(router.DELETE("/{name}").To(listPod.Delete))
 
    return router
}
