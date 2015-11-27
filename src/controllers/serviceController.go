package controllers

import (
    "github.com/emicklei/go-restful"
    "services/k8sService"
    "io/ioutil"
)
import (
	"fmt"
    "encoding/json"

)

type ServiceController struct {

}

func (this *ServiceController) Get(request *restful.Request, response *restful.Response) {
	fmt.Println("[Get] Service")
    service_name := request.PathParameter("name")
    
    var list_service k8sService.ServiceService

    response.WriteEntity(list_service.GetService(service_name))
    return
}

func (this *ServiceController) GetList(request *restful.Request, response *restful.Response) {
	fmt.Println("[GetList] Service List")
	var list_service k8sService.ServiceService

    response.WriteEntity(list_service.GetServiceList())
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
    return
}

func (this *ServiceController) Delete(request *restful.Request, response *restful.Response) {
    fmt.Println("[Delete] Delete Service")

    service_name := request.PathParameter("name")
    var list_service k8sService.ServiceService

    response.WriteEntity(list_service.DeleteService(service_name))
    return
}

func (this *ServiceController) CreateBoth(request *restful.Request, response *restful.Response) {
    fmt.Println("[Post] both")

    body, err := ioutil.ReadAll(request.Request.Body)

    if err != nil {
        response.WriteEntity("fail")
        return 
    }

    var list_service k8sService.ServiceService
    response.WriteEntity(list_service.CreateBoth(body))
    return
}


func Test(result interface{}) {
    js, _ := json.Marshal(result)
    fmt.Println(string(js))
}