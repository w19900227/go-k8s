package controllers

import (
    "github.com/emicklei/go-restful"
    "services"
)
import (
	"fmt"
)

type BubbleController struct {

}

func (this *BubbleController) GetRelate(request *restful.Request, response *restful.Response) {
	fmt.Println("[Get] Relate Controller")
    
    var _line_service services.BubbleService
    response.WriteEntity(_line_service.GetRelate())
    return
}
func (this *BubbleController) GetInfo(request *restful.Request, response *restful.Response) {
	fmt.Println("[Get] Info Controller")
	
    var _line_service services.BubbleService
    response.WriteEntity(_line_service.GetInfo())
    return
}
