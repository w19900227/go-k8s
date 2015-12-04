package k8sModel

import (
	bm "models/baseModel"
	"encoding/json"
	"libs/request"
	"models/k8sModel/format"
)

type PodModel struct {
	bm.BaseModel
	req request.Request
}

//curl -H "Content-Type: application/json" -X GET  http://192.168.12.8:8080/api/v1/namespaces/default/pods
func (this *PodModel) GetPodList() (format.PodList) {
	url := this.GetK8SUrl("pods", bm.Namespace)
	body := this.req.Get(url)

    var data format.PodList
	json.Unmarshal(body, &data)
	return data
}

func (this *PodModel) PodByName(name string) (format.Pod) {
	url := this.GetK8SUrl("pods/"+name, bm.Namespace)
	body := this.req.Get(url)

    var data format.Pod
	json.Unmarshal(body, &data)
	return data
}

func (this *PodModel) DeletePod(name string) (format.Pod) {
	url := this.GetK8SUrl("pods/"+name, bm.Namespace)
	body := this.req.Delete(url, nil)
	
	var _pod_data format.Pod
	json.Unmarshal(body, &_pod_data)
	return _pod_data
}




