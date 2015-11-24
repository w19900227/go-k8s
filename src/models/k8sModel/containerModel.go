package k8sModel

import (
	"fmt"

	bm "models/baseModel"
	"encoding/json"
	"libs/request"
	"models/k8sModel/format"
)

type PodModel struct {
	bm.BaseModel
	req request.Request
}

// func (this *PodModel) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

// func (this *PodModel) GetContainer() (format.Pod) {}

//curl -H "Content-Type: application/json" -X GET  http://192.168.12.8:8080/api/v1/namespaces/default/pods
func (this *PodModel) GetPodList() (format.PodList) {
	body := this.req.Get("http://192.168.12.8:8080/api/v1/namespaces/default/pods")
	// body := this.getPagInfo_data()//在伺服器run必須移除此行

    var data format.PodList
	json.Unmarshal(body, &data)
	// fmt.Println(err)
	// fmt.Println(b)
	return data
}
func (this *PodModel) PodByName(name string) (format.Pod) {
	body := this.req.Get("http://192.168.12.8:8080/api/v1/namespaces/default/pods/"+name)
	// body := this.getPagInfo_data()//在伺服器run必須移除此行
    var data format.Pod
	json.Unmarshal(body, &data)
	// fmt.Println(err)
	// fmt.Println(b)
	return data
}
func (this *PodModel) DeletePod(name string) (format.Pod) {
	body := this.req.Delete("http://192.168.12.8:8080/api/v1/namespaces/default/pods/"+name, nil)
	
	var _pod_data format.Pod
	json.Unmarshal(body, &_pod_data)
	
	return _pod_data
}

func (this *PodModel) Test22() {
	fmt.Println("ssss")
}





