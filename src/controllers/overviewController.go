package controllers

import (
    "github.com/emicklei/go-restful"
    "services"
)
import (
    "fmt"
)

type OverviewController struct {

}

func (this *OverviewController) Get(request *restful.Request, response *restful.Response) {
	fmt.Println("[Get] Overview")
	
    var _overview_service services.OverviewService
	// fmt.Println(_overview_service.GetServiceCount())
	// fmt.Println(_overview_service.GetReplicationControllerCount())
	// fmt.Println(_overview_service.GetPodCount())
    // fmt.Println(_overview_service.GetMachineCount())
	// fmt.Println(_overview_service.GetAllCount())
    // response.WriteEntity("overview")
    response.WriteEntity(_overview_service.GetAllCount())
    return
}

