package controllers

import (
    "github.com/emicklei/go-restful"
    "services/k8sService"
)
import (
	"fmt"
)

type EventController struct {

}

func (this *EventController) GetList(request *restful.Request, response *restful.Response) {
	fmt.Println("[Get] EventList Controller")
    
    var _line_service k8sService.EventService
    response.WriteEntity(_line_service.GetEventList())
    return
}
