package baseService

import (
// 	"fmt"
// 	"runtime"
// 	"time"
	// "fmt"
	// "encoding/json"
)

type BaseService struct {
	baseurl string 
	namespace string 
	hred_ip string 

	// baseurl string = "http://192.168.6.13:8080/api/v1"
	// namespace string = "/namespaces/default/"
	// hred_ip string = "http://192.168.6.14:8090"
}

// func (this *BaseService) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

func (this *BaseService) GetBaseUrl() string {
	return "http://192.168.6.13:8080/api/v1" + "/namespaces/default/"
}

func (this *BaseService) GetHredIP() string {
	return "http://192.168.6.14:8090"
	// return "http://192.168.6.14:8090" + "/namespaces/default/"
}

// func (this *BaseService) JsonRequest(action string, data ) {
	
// }




// func (this *BaseService) AjaxResponse(status string, erron string, errmsg string, data interface{}) {
// 	response := struct {
// 		Status string
// 		Erron  string
// 		Errmsg string
// 		Data   interface{}
// 	}{
// 		Status : status,
// 		Erron  : erron,
// 		Errmsg : errmsg,
// 		Data   : data,
// 	}

// 	this.Data["json"] = response
// 	this.ServeJson()
// }