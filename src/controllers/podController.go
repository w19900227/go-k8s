package controllers

import (
    "github.com/emicklei/go-restful"
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
    // response.WriteEntity(r)
    return
}

func (this *PodController) GetList(request *restful.Request, response *restful.Response) {
	fmt.Println("[GetList] Pod List")
    // response.WriteEntity(r)
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
	
    fmt.Print(pod_name)
    // response.WriteEntity(r)
    return
}

func (this *PodController) Delete(request *restful.Request, response *restful.Response) {
    fmt.Println("[Delete] Delete Pod")
    pod_name := request.PathParameter("name")
	
    fmt.Print(pod_name)
    // response.WriteEntity(r)
    return
}