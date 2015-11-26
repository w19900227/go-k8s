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

type line struct {
	Status string `json:"status,omitempty"`
	Errno string `json:"errno,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	Data Overview `json:"data,omitempty"`
}

func (this *LineService) Post(data []byte) line {
	fmt.Println("[Post] LineService")

	var _data_format data_format
	json.Unmarshal(data, &_data_format)

	service_model := k8sModel.ServiceModel{}
    _get_service_data := service_model.GetService(_data_format.Service_name)

	replication_controller_model := k8sModel.ReplicationControllerModel{}
    _get_replication_controller_data := replication_controller_model.GetReplicationController(_data_format.Cluster_name)

    var _line line

    if _get_service_data.Spec.Selector == nil {
    	_line.Status = "fail"
    	_line.Errmsg = "service no selector"
    	return _line
    }

	_get_service_data.Spec.Selector = _get_replication_controller_data.Spec.Selector

	service_model.UpdateService(_data_format.Service_name, _get_service_data)
	// _put_service_data := service_model.UpdateService(_data_format.Service_name, _get_service_data)
	// if _put_service_data.Status_code == "200" {
	// 	Status = "ok"
	// } else {
	// 	Status = "fail"
	// }
    _line.Status = "ok"
    return _line
}
func (this *LineService) Delete(data []byte) line {
	fmt.Println("[Delete] LineService")
	
	var _data_format data_format
	json.Unmarshal(data, &_data_format)

	service_model := k8sModel.ServiceModel{}
    _get_service_data := service_model.GetService(_data_format.Service_name)

	replication_controller_model := k8sModel.ReplicationControllerModel{}
    _get_replication_controller_data := replication_controller_model.GetReplicationController(_data_format.Cluster_name)

    var _line line

    if _get_service_data.Spec.Selector == nil {
    	_line.Status = "fail"
    	_line.Errmsg = "service no selector"
    	return _line
    }

    if _get_service_data.Spec.Selector["name"] != _get_replication_controller_data.Spec.Selector["name"] {
    	_line.Status = "fail"
    	_line.Errmsg = "service and cluster not same"
    	return _line
    }

    _get_service_data.Spec.Selector["name"] = _get_service_data.Spec.Selector["name"]+"-giga"
    service_model.UpdateService(_data_format.Service_name, _get_service_data)

    _line.Status = "ok"
    return _line
}

