package baseService

import (
// 	"fmt"
// 	"runtime"
// 	"time"
	"fmt"
	"encoding/json"
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
func (this *BaseService) Test(result interface{}) {
	js, _ := json.Marshal(result)
	fmt.Println(string(js))
}