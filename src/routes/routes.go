package routes

import (
	// "fmt"
)
import (
    "github.com/emicklei/go-restful"
	"controllers"
    "filter"
)

type Routes struct {

}

func (this *Routes) IndexRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)
    var filter filter.Filter
    router.Filter(filter.AuthToken)

    // listService := controllers.ServiceController{}
    var overview controllers.OverviewController
    router.Route(router.GET("/overview").To(overview.Get))

    var namespace controllers.NamespaceController
    router.Route(router.GET("/namespace").To(namespace.Get))

 
    return router
}


func (this *Routes) ServiceRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/service").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)
    var filter filter.Filter
    router.Filter(filter.AuthToken)

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

func (this *Routes) ServiceRoutingRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/service/routing").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)
    var filter filter.Filter
    router.Filter(filter.AuthToken)
    
    // listService := controllers.ServiceController{}
    var listService controllers.ServiceController
    // router.Route(router.GET("/{name}").To(listService.Get))
    // router.Route(router.GET("/").To(listService.GetList))
    router.Route(router.POST("/").To(listService.CreateServiceRouting))
    // router.Route(router.PUT("/{name}").To(listService.Put))
    // router.Route(router.DELETE("/{name}").To(listService.Delete))
    // router.Route(router.POST("/both").To(listService.CreateBoth))
    // router.Route(router.POST("/routing").To(listService.CreateServiceRouting))

    return router
}

func (this *Routes) ReplicationControllerRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/cluster").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)
    var filter filter.Filter
    router.Filter(filter.AuthToken)
    
	var listReplicationController controllers.ReplicationControllerController
    router.Route(router.GET("/{name}").To(listReplicationController.Get))
    router.Route(router.GET("/").To(listReplicationController.GetList))
    router.Route(router.POST("/").To(listReplicationController.Post))
    router.Route(router.PUT("/{name}").To(listReplicationController.Put))
    router.Route(router.DELETE("/{name}").To(listReplicationController.Delete))
 
    return router
}

func (this *Routes) PodRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/container").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)
    var filter filter.Filter
    router.Filter(filter.AuthToken)
    
    var listPod controllers.PodController
    router.Route(router.GET("/{name}").To(listPod.Get))
    router.Route(router.GET("/").To(listPod.GetList))
    router.Route(router.POST("/").To(listPod.Post))
    router.Route(router.PUT("/{name}").To(listPod.Put))
    router.Route(router.DELETE("/{name}").To(listPod.Delete))
 
    return router
}

func (this *Routes) NodeRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/machine").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)
    var filter filter.Filter
    router.Filter(filter.AuthToken)
    
    var listNode controllers.NodeController
    router.Route(router.GET("/").To(listNode.GetList))
 
    return router
}

func (this *Routes) EventRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/event").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)
    var filter filter.Filter
    router.Filter(filter.AuthToken)
    
    var listNode controllers.EventController
    router.Route(router.GET("/").To(listNode.GetList))
 
    return router
}

func (this *Routes) DockerRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/docker").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)
    var filter filter.Filter
    router.Filter(filter.AuthToken)
    
    var listDocker controllers.DockerController
    router.Route(router.POST("/search").To(listDocker.SearchImages))
 
    return router
}

func (this *Routes) LineRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/line").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)
    var filter filter.Filter
    router.Filter(filter.AuthToken)
    
    var line controllers.LineController
    router.Route(router.POST("/").To(line.Post))
    router.Route(router.DELETE("/").To(line.Delete))
 
    return router
}

func (this *Routes) BubbleRouter() *restful.WebService {
    router := new(restful.WebService)
    router.
        Path("/bubble").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)
    var filter filter.Filter
    router.Filter(filter.AuthToken)
    
    var bubble controllers.BubbleController
    router.Route(router.GET("/info").To(bubble.GetInfo))
    router.Route(router.GET("/relate").To(bubble.GetRelate))
 
    return router
}
