package k8sModel

import (
	bm "models/baseModel"
	"encoding/json"
	"libs/request"
	"models/k8sModel/format"
	// format "models/k8sModel/tmpformat"
	// "github.com/kubernetes/kubernetes/pkg/api"
	// "k8s.io/kubernetes/pkg/api"
)
import (
	"fmt"
)

type ServiceModel struct {
	bm.BaseModel
	req request.Request
}

// func (this *ServiceModel) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

// curl -H "Content-Type: application/json" -X GET  http://192.168.12.8:8080/api/v1/namespaces/default/services
// func (this *ServiceModel) GetServiceList() (api.ServiceList) {
func (this *ServiceModel) GetServiceList() (format.ServiceList) {
	body := this.req.Get("http://192.168.12.8:8080/api/v1/namespaces/default/services")
	// body := this.getPagInfo_data()//在伺服器run必須移除此行

	// var data api.ServiceList
	var data format.ServiceList
	json.Unmarshal(body, &data)
	// fmt.Println(err)
	// fmt.Println(b)
	return data
}
func (this *ServiceModel) GetService(name string) (format.Service) {
	body := this.req.Get("http://192.168.12.8:8080/api/v1/namespaces/default/services/"+name)
	// body := this.getPagInfo_data()//在伺服器run必須移除此行

	// var data api.ServiceList
	var data format.Service
	json.Unmarshal(body, &data)
	// fmt.Println(err)
	// fmt.Println(b)
	return data
}

func (this *ServiceModel) CreateService(data format.Service) (format.Service) {
	body := this.req.Post("http://192.168.12.8:8080/api/v1/namespaces/default/services", data)
	
	var _service_data format.Service
	json.Unmarshal(body, &_service_data)

	return _service_data
}

func (this *ServiceModel) UpdateService(name string, data format.Service) (format.Service) {
	body := this.req.Put("http://192.168.12.8:8080/api/v1/namespaces/default/services/"+name, data)
	
	var _service_data format.Service
	json.Unmarshal(body, &_service_data)

	return _service_data
}

func (this *ServiceModel) DeleteService(name string) (format.Service) {
	body := this.req.Delete("http://192.168.12.8:8080/api/v1/namespaces/default/services/"+name, nil)
	
	var _service_data format.Service
	json.Unmarshal(body, &_service_data)
	
	return _service_data
}

func Test(result interface{}) {
	js, _ := json.Marshal(result)
	fmt.Println(string(js))
}



