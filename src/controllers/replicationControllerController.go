package controllers

import (
    "github.com/emicklei/go-restful"
)
import (
	"fmt"
)

type ReplicationControllerController struct {

}

func (this *ReplicationControllerController) Get(request *restful.Request, response *restful.Response) {
	fmt.Println("[Get] ReplicationController")
    rc_name := request.PathParameter("name")
	
    fmt.Print(rc_name)
    // response.WriteEntity(r)
    return
}

func (this *ReplicationControllerController) GetList(request *restful.Request, response *restful.Response) {
	fmt.Println("[GetList] ReplicationController List")
    // response.WriteEntity(r)
    return
}

func (this *ReplicationControllerController) Post(request *restful.Request, response *restful.Response) {
	fmt.Println("[Post] Create ReplicationController")
    // response.WriteEntity(r)
    return
}

func (this *ReplicationControllerController) Put(request *restful.Request, response *restful.Response) {
	fmt.Println("[Put] Editor ReplicationController")
    rc_name := request.PathParameter("name")
	
    fmt.Print(rc_name)
    // response.WriteEntity(r)
    return
}

func (this *ReplicationControllerController) Delete(request *restful.Request, response *restful.Response) {
    fmt.Println("[Delete] Delete ReplicationController")
    rc_name := request.PathParameter("name")
	
    fmt.Print(rc_name)
    // response.WriteEntity(r)
    return
}