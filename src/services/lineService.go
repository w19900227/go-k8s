package services

import (
	"fmt"

// 	"runtime"
// 	"time"
)

import (
	bs "services/baseService"
	"encoding/json"
	"models/k8sModel"
	// "models/herdModel"
	// "strings"
	// "strconv"
	// herd_format "models/herdModel/format"

)

type LineService struct {
	bs.BaseService
}

type data_format struct {
	Service_name string `json:"service_name,omitempty"`
	Cluster_name string `json:"cluster_name,omitempty"`
}

func (this *LineService) Post(data []byte) string {
	fmt.Println("[Post] LineService")

	var _data_format data_format
	json.Unmarshal(data, &_data_format)

	service_model := k8sModel.ServiceModel{}
    _get_service_data := service_model.GetService(_data_format.Service_name)

	replication_controller_model := k8sModel.ReplicationControllerModel{}
    _get_replication_controller_data := replication_controller_model.GetReplicationController(_data_format.Cluster_name)

	// if _get_service_data.Spec.Selector == nil {
	// 	_get_service_data.Spec.Selector = _get_replication_controller_data.Spec.Selector
	// } else {
	// 	_get_service_data.Spec.Selector = _get_replication_controller_data.Spec.Selector
	// }

    if _get_service_data.Spec.Selector == nil {
    	fmt.Println("service no selector")
    	return "service no selector"
    }

	_get_service_data.Spec.Selector = _get_replication_controller_data.Spec.Selector

	service_model.UpdateService(_data_format.Service_name, _get_service_data)
	// _put_service_data := service_model.UpdateService(_data_format.Service_name, _get_service_data)
	// if _put_service_data.Status_code == "200" {
	// 	Status = "ok"
	// } else {
	// 	Status = "fail"
	// }
	Test(_get_service_data)
    return "sss"
}
func (this *LineService) Delete(data []byte) string {
	fmt.Println("[Delete] LineService")
	
	var _data_format data_format
	json.Unmarshal(data, &_data_format)

	service_model := k8sModel.ServiceModel{}
    _get_service_data := service_model.GetService(_data_format.Service_name)

	replication_controller_model := k8sModel.ReplicationControllerModel{}
    _get_replication_controller_data := replication_controller_model.GetReplicationController(_data_format.Cluster_name)

    if _get_service_data.Spec.Selector == nil {
    	fmt.Println("service no selector")
    	return "service no selector"
    }

    if _get_service_data.Spec.Selector["name"] != _get_replication_controller_data.Spec.Selector["name"] {
    	fmt.Println("status : fail")
    	return "status : fail"
    }

    _get_service_data.Spec.Selector["name"] = _get_service_data.Spec.Selector["name"]+"-giga"
    service_model.UpdateService(_data_format.Service_name, _get_service_data)




	// if _get_service_data.Spec.Selector == nil {
	// 	Status = "fail"
	// 	Errmsg = "There is not have relate between Service and Cluster"
	// }

	// _get_service_data.Spec.Selector = map[string]string{"name":_data_format.Cluster_name}
	
	Test(_get_service_data)
	
    // response.WriteEntity(r)
    return "sss"
}

