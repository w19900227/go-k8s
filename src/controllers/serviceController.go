package controllers

import (
    "github.com/emicklei/go-restful"
    "services"
)
import (
	"fmt"
)

type ServiceController struct {

}

func (this *ServiceController) Get(request *restful.Request, response *restful.Response) {
	fmt.Println("[Get] Service")
    service_name := request.PathParameter("name")
	
    fmt.Print(service_name)
    // response.WriteEntity(r)
    return
}

func (this *ServiceController) GetList(request *restful.Request, response *restful.Response) {
	fmt.Println("[GetList] Service List")
	var list_service services.ListService
	fmt.Println(list_service.Service())
    // response.WriteEntity(r)
    return
}

func (this *ServiceController) Post(request *restful.Request, response *restful.Response) {
	fmt.Println("[Post] Create Service")
    // response.WriteEntity(r)
    return
}

func (this *ServiceController) Put(request *restful.Request, response *restful.Response) {
	fmt.Println("[Put] Editor Service")
    service_name := request.PathParameter("name")
	
    fmt.Print(service_name)
    // response.WriteEntity(r)
    return
}

func (this *ServiceController) Delete(request *restful.Request, response *restful.Response) {
    fmt.Println("[Delete] Delete Service")
    service_name := request.PathParameter("name")
	
    fmt.Print(service_name)
    // response.WriteEntity(r)
    return
}