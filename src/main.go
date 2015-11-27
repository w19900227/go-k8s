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
	
	
    restful.Add(IndexRouter())
    restful.Add(ServiceRouter())
    restful.Add(ReplicationControllerRouter())
    restful.Add(PodRouter())
    restful.Add(NodeRouter())
    restful.Add(EventRouter())



    // log.Fatal(http.ListenAndServe(":9090", nil))
    log.Fatal(http.ListenAndServe(":31001", nil))
}


func IndexRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)

    // listService := controllers.ServiceController{}
    var overview controllers.OverviewController

    router.Route(router.GET("/overview").To(overview.Get))

    var line controllers.LineController

    router.Route(router.POST("/line").To(line.Post))
    router.Route(router.DELETE("/line").To(line.Delete))

    var bubble controllers.BubbleController

    router.Route(router.GET("/bubble/info").To(bubble.GetInfo))
    router.Route(router.GET("/bubble/relate").To(bubble.GetRelate))
 
    return router
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
    router.Route(router.POST("/both").To(listService.CreateBoth))

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

func NodeRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/machine").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)

    // listNode := controllers.PodController{}
    var listNode controllers.NodeController

    // router.Route(router.GET("/{name}").To(listNode.Get))
    router.Route(router.GET("/").To(listNode.GetList))
    // router.Route(router.POST("/").To(listNode.Post))
    // router.Route(router.PUT("/{name}").To(listNode.Put))
    // router.Route(router.DELETE("/{name}").To(listNode.Delete))
 
    return router
}

func EventRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/event").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)

    // listNode := controllers.PodController{}
    var listNode controllers.EventController

    // router.Route(router.GET("/{name}").To(listNode.Get))
    router.Route(router.GET("/").To(listNode.GetList))
    // router.Route(router.POST("/").To(listNode.Post))
    // router.Route(router.PUT("/{name}").To(listNode.Put))
    // router.Route(router.DELETE("/{name}").To(listNode.Delete))
 
    return router
}
