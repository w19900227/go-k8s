package k8sModel

import (
	bm "models/baseModel"
	"encoding/json"
	"libs/request"
	"models/k8sModel/format"
)

type ReplicationControllerModel struct {
	bm.BaseModel
	req request.Request
}

func (this *ReplicationControllerModel) init() {
}

// curl -H "Content-Type: application/json" -X GET  http://192.168.12.8:8080/api/v1/namespaces/ReplicationController/replicationcontrollers
func (this *ReplicationControllerModel) GetReplicationControllerList() (format.ReplicationControllerList) {
	url := this.GetK8SUrl("replicationcontrollers", bm.Namespace)
	body := this.req.Get(url)
	// body := this.req.Get("http://192.168.12.8:8080/api/v1/namespaces/default/replicationcontrollers")

    var data format.ReplicationControllerList
	json.Unmarshal(body, &data)
	return data
}

func (this *ReplicationControllerModel) GetReplicationController(name string) (format.ReplicationController) {
	url := this.GetK8SUrl("replicationcontrollers/"+name, bm.Namespace)
	body := this.req.Get(url)
	// body := this.req.Get("http://192.168.12.8:8080/api/v1/namespaces/default/replicationcontrollers/"+name)

    var data format.ReplicationController
	json.Unmarshal(body, &data)
	return data
}

func (this *ReplicationControllerModel) CreateReplicationController(data format.ReplicationController) (format.ReplicationController) {
	url := this.GetK8SUrl("replicationcontrollers", bm.Namespace)
	body := this.req.Post(url, data)
	
	var _replication_controller_data format.ReplicationController
	json.Unmarshal(body, &_replication_controller_data)

	return _replication_controller_data
}

func (this *ReplicationControllerModel) UpdateReplicationController(name string, data format.ReplicationController) (format.ReplicationController) {
	url := this.GetK8SUrl("replicationcontrollers/"+name, bm.Namespace)
	body := this.req.Put(url, data)
	
	var _replication_controller_data format.ReplicationController
	json.Unmarshal(body, &_replication_controller_data)

	return _replication_controller_data
}

func (this *ReplicationControllerModel) DeleteReplicationController(name string) (format.ReplicationController) {
	url := this.GetK8SUrl("replicationcontrollers/"+name, bm.Namespace)
	body := this.req.Delete(url, nil)
	
	var _replication_controller_data format.ReplicationController
	json.Unmarshal(body, &_replication_controller_data)
	
	return _replication_controller_data
}
