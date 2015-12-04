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
	// "fmt"
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
	url := this.GetK8SUrl("services", bm.Namespace)
	body := this.req.Get(url)
	// body := this.getPagInfo_data()//在伺服器run必須移除此行

	// var data api.ServiceList
	var data format.ServiceList
	json.Unmarshal(body, &data)
	// fmt.Println(err)
	// fmt.Println(b)
	return data
}
func (this *ServiceModel) GetService(name string) (format.Service) {
	url := this.GetK8SUrl("services/"+name, bm.Namespace)
	body := this.req.Get(url)
	// body := this.getPagInfo_data()//在伺服器run必須移除此行

	// var data api.ServiceList
	var data format.Service
	json.Unmarshal(body, &data)
	// fmt.Println(err)
	// fmt.Println(b)
	return data
}

func (this *ServiceModel) CreateService(data format.Service) (format.Service) {
	url := this.GetK8SUrl("services", bm.Namespace)
	body := this.req.Post(url, data)
	
	var _service_data format.Service
	json.Unmarshal(body, &_service_data)

	return _service_data
}

func (this *ServiceModel) UpdateService(name string, data format.Service) (format.Service) {
	url := this.GetK8SUrl("services/"+name, bm.Namespace)
	body := this.req.Put(url, data)
	
	var _service_data format.Service
	json.Unmarshal(body, &_service_data)

	return _service_data
}

func (this *ServiceModel) DeleteService(name string) (format.Service) {
	url := this.GetK8SUrl("services/"+name, bm.Namespace)
	body := this.req.Delete(url, nil)
	
	var _service_data format.Service
	json.Unmarshal(body, &_service_data)
	
	return _service_data
}




