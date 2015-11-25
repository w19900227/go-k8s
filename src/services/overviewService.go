package services

import (
	"fmt"
	"encoding/json"

// 	"runtime"
// 	"time"
)

import (
	bs "services/baseService"
	"models/k8sModel"
	"models/herdModel"
	"strings"
	"strconv"
	herd_format "models/herdModel/format"

)

type OverviewService struct {
	bs.BaseService
}

type service_info struct {
	Name string `json:"name"` 
	Container_run int `json:"container_run"`  
	Container_total int `json:"container_total"` 
	Cpu int `json:"cpu"` 
	Mem float64 `json:"mem"` 
}
type machine_info struct {
	Name string `json:"name"`
	Ip string `json:"ip"`
	Cpu int `json:"cpu"`
	Mem float64 `json:"mem"`
	Status  string `json:"status"`
}
type Overview struct {
	Status string `json:"status,omitempty"`
	Errno string `json:"errno,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`

	ReplicationController_count int `json:"cluster_count,omitempty"`
	Pod_count int `json:"container_count,omitempty"`
	Machine_count int `json:"machine_count,omitempty"`
	Service_count int `json:"service_count,omitempty"`
	Service_info []service_info `json:"service_info,omitempty"`
	Machine_info []machine_info `json:"machine_info,omitempty"`
}
type OverviewList struct {
	Status string `json:"status,omitempty"`
	Errno string `json:"errno,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	Data Overview `json:"data,omitempty"`
}

func (this *OverviewService) GetAllCount() OverviewList {
	var _overview_list OverviewList
	_overview_list.Status = "ok"

	_machine_count := this.GetMachineCount()
	_overview_list.Data.Machine_count = _machine_count.Data.Machine_count
	
	_service_count := this.GetServiceCount()
	_overview_list.Data.Service_count = _service_count.Data.Service_count

	_replication_controller_count := this.GetReplicationControllerCount()
	_overview_list.Data.ReplicationController_count = _replication_controller_count.Data.ReplicationController_count

	_pod_count := this.GetPodCount()
	_overview_list.Data.Pod_count = _pod_count.Data.Pod_count

	_service_info := this.GetServiceInfo()
	_overview_list.Data.Service_info = _service_info.Data.Service_info

	_machine_info := this.GetMachineInfo()
	_overview_list.Data.Machine_info = _machine_info.Data.Machine_info

	return _overview_list
}

func (this *OverviewService) GetServiceCount() OverviewList {
	service_model := k8sModel.ServiceModel{}
    data := service_model.GetServiceList()

	var _overview_list OverviewList
	_overview_list.Status = "ok"

	var i int = 0
 	for _, data_items := range data.Items {
		metadata_name := data_items.Metadata.Name
		if metadata_name == "kube-dns" || metadata_name == "kubernetes" {
			continue
		}
		i = i + 1
	}
	_overview_list.Data.Service_count = i
	return _overview_list
}

func (this *OverviewService) GetReplicationControllerCount() OverviewList {
	_replication_controllerModel := k8sModel.ReplicationControllerModel{}
    data := _replication_controllerModel.GetReplicationControllerList()

    var _overview_list OverviewList
	_overview_list.Status = "ok"

	var i int = 0
    for _, data_items := range data.Items {
		metadata_name := data_items.Metadata.Name

    	if metadata_name == "kube-dns-v3" {
    		continue
    	}
		i = i + 1
    }
	_overview_list.Data.ReplicationController_count = i
	return _overview_list
}

func (this *OverviewService) GetPodCount() OverviewList {
	_podModel := k8sModel.PodModel{}
	data := _podModel.GetPodList()

    var _overview_list OverviewList
	_overview_list.Status = "ok"

	var i int = 0
	for _, data_items := range data.Items {
		_metadata_name_arr := strings.Split(data_items.Metadata.Name, "-")
		_metadata_name := strings.Join(_metadata_name_arr[:len(_metadata_name_arr)-1], "-")
		if _metadata_name == "kube-dns-v3" {
			continue
		}

		i = i + 1
    }
	_overview_list.Data.Pod_count = i
	return _overview_list
}

func (this *OverviewService) GetMachineCount() OverviewList {
	_nodeModel := k8sModel.NodeModel{}
	data := _nodeModel.GetNodeList()

    var _overview_list OverviewList
	_overview_list.Status = "ok"
	_overview_list.Data.Machine_count = len(data.Items)
	return _overview_list
}

//未完成
func (this *OverviewService) GetServiceInfo() OverviewList {
	service_model := k8sModel.ServiceModel{}
    data := service_model.GetServiceList()

	rc_model := k8sModel.ReplicationControllerModel{}
    rc_data := rc_model.GetReplicationControllerList()

    machine_data := this.GetMachineInfo()
    var total_mem float64 = 0
    for _, _machine_data := range machine_data.Data.Machine_info {
    	total_mem += _machine_data.Mem
    }

	var _overview_list OverviewList
	_overview_list.Status = "ok"

 	for _, data_items := range data.Items {
		metadata_name := data_items.Metadata.Name
		if metadata_name == "kube-dns" || metadata_name == "kubernetes" {
			continue
		}

		var _service_info service_info
		_service_info.Name = metadata_name

		var herd_service herd_format.Data
		herd_service.Services = append(herd_service.Services, metadata_name)

		var getHerd herdModel.GetHerdModel
		getHerd_data := getHerd.PostData(herd_service)

		var cpu int = 0
		var mem int = 0
		for _, service_herd := range getHerd_data.Services {
			if len(service_herd.Clusters) > 0 {
				for _, cluster_herd := range service_herd.Clusters {
					if len(cluster_herd.Containers) > 0 {
						for _, container_herd := range cluster_herd.Containers {
							cpu += container_herd.Cpu
							mem += container_herd.Mem
						}
					}
				}
			}
		}
		_service_info.Cpu = cpu
		_service_info.Mem = (float64(mem)/total_mem)*100
		for _, _rc_data_itme := range rc_data.Items {
			if data_items.Spec.Selector != nil {
				if data_items.Spec.Selector["name"] == _rc_data_itme.Spec.Selector["name"] {
					_service_info.Container_total += _rc_data_itme.Spec.Replicas
					_service_info.Container_run += _rc_data_itme.Spec.Replicas
					
				}
			}
		}
		_overview_list.Data.Service_info = append(_overview_list.Data.Service_info, _service_info)
	}

	return _overview_list
}

func (this *OverviewService) GetMachineInfo() OverviewList {
	_nodeModel := k8sModel.NodeModel{}
    data := _nodeModel.GetNodeList()

	var _overview_list OverviewList
	_overview_list.Status = "ok"
 	for _, data_items := range data.Items {
		var _machine_info machine_info
 		_machine_info.Name = data_items.Metadata.Name
 		_machine_info.Ip = data_items.Status.Addresses[0].Address
 		_machine_info.Cpu, _ = strconv.Atoi(data_items.Status.Capacity["cpu"])

 		mem_str := strings.Split(data_items.Status.Capacity["memory"], "Ki")[0]
 		mem, _ := strconv.ParseFloat(mem_str, 64)
 		_machine_info.Mem = mem/1024

 		if data_items.Status.Conditions[0].Status == "True" {
 			_machine_info.Status = "Ready"
 		} else {
 			_machine_info.Status = data_items.Status.Conditions[0].Reason
 		}

 		_overview_list.Data.Machine_info = append(_overview_list.Data.Machine_info, _machine_info)
	}
	return _overview_list
}
func Test(result interface{}) {
	js, _ := json.Marshal(result)
	fmt.Println(string(js))
}

func (this *OverviewService) AABB() string {
	fmt.Println("ssswww")
	// return "sssss"
	return this.GetBaseUrl()
}

