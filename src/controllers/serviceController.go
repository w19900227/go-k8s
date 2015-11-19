package controllers

import (
    "github.com/emicklei/go-restful"
    "services/k8sService"
)
import (
	"fmt"
    "encoding/json"
    "io/ioutil"

)

type ServiceController struct {

}

func (this *ServiceController) Get(request *restful.Request, response *restful.Response) {
	fmt.Println("[Get] Service")
    service_name := request.PathParameter("name")
    
    var list_service k8sService.ServiceService
    // fmt.Println(list_service.Service())

    response.WriteEntity(list_service.ServiceByName(service_name))
    // response.WriteEntity(r)
    return
}

func (this *ServiceController) GetList(request *restful.Request, response *restful.Response) {
	fmt.Println("[GetList] Service List")
	var list_service k8sService.ServiceService
	// fmt.Println(list_service.Service())

    response.WriteEntity(list_service.Service())
    // response.WriteEntity(r)
    return
}

func (this *ServiceController) Post(request *restful.Request, response *restful.Response) {
	fmt.Println("[Post] Create Service")
    var list_service k8sService.ServiceService

    body, err := ioutil.ReadAll(request.Request.Body)
    
    if err != nil {
        response.WriteEntity("fail")
        return 
    }

    response.WriteEntity(list_service.CreateService(body))
    // response.WriteEntity(r)
    return
}

func (this *ServiceController) Put(request *restful.Request, response *restful.Response) {
	fmt.Println("[Put] Editor Service")
    service_name := request.PathParameter("name")
	
    var list_service k8sService.ServiceService

    body, err := ioutil.ReadAll(request.Request.Body)
    
    if err != nil {
        response.WriteEntity("fail")
        return 
    }
    response.WriteEntity(list_service.UpdateService(service_name, body))
    // response.WriteEntity(r)
    return
}

func (this *ServiceController) Delete(request *restful.Request, response *restful.Response) {
    fmt.Println("[Delete] Delete Service")

    service_name := request.PathParameter("name")
    var list_service k8sService.ServiceService

    response.WriteEntity(list_service.DeleteService(service_name))
    return
}


func Test(result interface{}) {
    js, _ := json.Marshal(result)
    fmt.Println(string(js))
}