package controllers

import (
    "github.com/emicklei/go-restful"
    "services"
    "io/ioutil"
)
import (
    "fmt"
)

type LineController struct {

}

func (this *LineController) Post(request *restful.Request, response *restful.Response) {
	fmt.Println("[Post] Line")
	
    body, err := ioutil.ReadAll(request.Request.Body)
    
    if err != nil {
        response.WriteEntity("fail")
        return 
    }
    
    var _line_service services.LineService
	fmt.Println(_line_service.Post(body))
    response.WriteEntity("LineController")
    // response.WriteEntity(r)
    return
}
func (this *LineController) Delete(request *restful.Request, response *restful.Response) {
	fmt.Println("[Delete] Line")
	
    body, err := ioutil.ReadAll(request.Request.Body)
    
    if err != nil {
        response.WriteEntity("fail")
        return 
    }
	
    var _line_service services.LineService
	fmt.Println(_line_service.Delete(body))
    response.WriteEntity("LineController")
    // response.WriteEntity(r)
    return
}