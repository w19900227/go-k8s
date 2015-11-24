package controllers

import (
    "github.com/emicklei/go-restful"
    "services/k8sService"
)
import (
	"fmt"
)

type PodController struct {

}

func (this *PodController) Get(request *restful.Request, response *restful.Response) {
	fmt.Println("[Get] Pod")
    pod_name := request.PathParameter("name")
    fmt.Print(pod_name)
	
    var _pod_service k8sService.PodService
    response.WriteEntity(_pod_service.PodByName(pod_name))
    return
}

func (this *PodController) GetList(request *restful.Request, response *restful.Response) {
	fmt.Println("[GetList] Pod List")
    var _pod_service k8sService.PodService
    response.WriteEntity(_pod_service.GetPodList())
    return
}

func (this *PodController) Post(request *restful.Request, response *restful.Response) {
	fmt.Println("[Post] Create Pod")
    // response.WriteEntity(r)
    return
}

func (this *PodController) Put(request *restful.Request, response *restful.Response) {
	fmt.Println("[Put] Editor Pod")
    pod_name := request.PathParameter("name")
	
    // response.WriteEntity(r)
    return
}

func (this *PodController) Delete(request *restful.Request, response *restful.Response) {
    fmt.Println("[Delete] Delete Pod")
    pod_name := request.PathParameter("name")
    var _pod_service k8sService.PodService

    response.WriteEntity(_pod_service.DeletePod(pod_name))
    return
}