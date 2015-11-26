package controllers

import (
    "github.com/emicklei/go-restful"
    "services/k8sService"
)
import (
	"fmt"
)

type NodeController struct {

}

func (this *NodeController) GetList(request *restful.Request, response *restful.Response) {
	fmt.Println("[Get] NodeList Controller")
    
    var _line_service k8sService.NodeService
    response.WriteEntity(_line_service.GetPodList())
    return
}
