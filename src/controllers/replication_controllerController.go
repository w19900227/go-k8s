package controllers

import (
    "github.com/emicklei/go-restful"
    "services/k8sService"
    "io/ioutil"
)
import (
	"fmt"
)

type ReplicationControllerController struct {

}

func (this *ReplicationControllerController) Get(request *restful.Request, response *restful.Response) {
	fmt.Println("[Get] ReplicationController")
    rc_name := request.PathParameter("name")
    
    var list_service k8sService.ReplicationControllerService

    response.WriteEntity(list_service.GetReplicationController(rc_name))
    return
}

func (this *ReplicationControllerController) GetList(request *restful.Request, response *restful.Response) {
	fmt.Println("[GetList] ReplicationController List")
    var list_service k8sService.ReplicationControllerService

    response.WriteEntity(list_service.GetReplicationControllerList())
    return
}

func (this *ReplicationControllerController) Post(request *restful.Request, response *restful.Response) {
	fmt.Println("[Post] Create ReplicationController")
    var list_service k8sService.ReplicationControllerService

    body, err := ioutil.ReadAll(request.Request.Body)
    if err != nil {
        response.WriteEntity("fail")
    }

    response.WriteEntity(list_service.CreateReplicationController(body))
    return
}

func (this *ReplicationControllerController) Put(request *restful.Request, response *restful.Response) {
	fmt.Println("[Put] Editor ReplicationController")
    rc_name := request.PathParameter("name")

    var list_service k8sService.ReplicationControllerService

    body, err := ioutil.ReadAll(request.Request.Body)
    if err != nil {
        response.WriteEntity("fail")
    }
	
    response.WriteEntity(list_service.UpdateReplicationController(rc_name, body))
    return
}

func (this *ReplicationControllerController) Delete(request *restful.Request, response *restful.Response) {
    fmt.Println("[Delete] Delete ReplicationController")
    rc_name := request.PathParameter("name")

    var list_service k8sService.ReplicationControllerService
    response.WriteEntity(list_service.DeleteReplicationController(rc_name))
	
    return
}