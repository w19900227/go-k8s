package k8sService

import (
    // "github.com/emicklei/go-restful"

// // 	"net/http"
	"fmt"
	"encoding/json"
// 	"strconv"
	// "io/ioutil"
	// "strings"
// 	"io"
	// "time"
)

import (
	bs "services/baseService"
	"models/k8sModel"
	"models/herdModel"
	// k8s "./k8s"
	"strconv"

	herd_format "models/herdModel/format"
	k8s_format "models/k8sModel/format"

	// "github.com/kubernetes/kubernetes/pkg/api/types"
	// "models/k8sModel/format"

)

type ServiceService struct {
	bs.BaseService
	// bm.BaseModel
}

// func (this *BaseService) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

// func (this *ServiceService) GetTest() string {
// 	getHerd := herdModel.GetHerdModel{}
// 	a := getHerd.GetPagInfo()
// 	fmt.Println(a.Services)
// 	// this.Test(a)
// 	return "ss"
// }

type Test2 struct {
	Status string `json:"status"`
}

type sub_Ports struct {
	Port int `json:"port"`
	External_port int `json:"external_port"`
}
type Service struct {
	Status string `json:"status,omitempty"`
	Errno string `json:"errno,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	Label map[string]string `json:"label,omitempty"`
	Ports []sub_Ports `json:"ports,omitempty"`
	MachineStatus string `json:"machine_status,omitempty"`
	Cpu int `json:"cpu"`
	Memory int `json:"memory"`
	Service_name string `json:"service_name,omitempty"`
}	
type ServiceList struct {
	Status string `json:"status,omitempty"`
	Errno string `json:"errno,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	Data []Service `json:"data,omitempty"`
}
// func (this *ServiceService) Service_bak() types.ServiceList {
//     service_model := k8sModel.ServiceModel{}
//     data := service_model.GetServiceList()
//     return data
//     // if len(data.Items) < 1 {
//     // 	fmt.Println("not found data")
//     // 	return 
//     // }
	
// 	var _service_list ServiceList
// 	var _service service
// 	var label labels
// 	var sub_Port sub_Ports

//  	for _, data_items := range data.Items {
// 		metadata_name := data_items.Metadata.Name
// 		_service.Service_name = metadata_name
// 		// fmt.Println(metadata_name)
// 		if metadata_name != "kube-dns" && metadata_name != "kubernetes" {
// 			// fmt.Println(metadata_name)
// 			if data_items.Metadata.Labels == nil {
// 				label.Label_key = "Null"
// 				label.Label_value = "Null"
// 				_service.Label = append(_service.Label, label)
// 			} else {
// 				for key, value := range data_items.Metadata.Labels {
// 					label.Label_key = key
// 					label.Label_value = value
// 					_service.Label = append(_service.Label, label)
// 				}

// 				var getHerd herdModel.GetHerdModel
// 				getHerd.Data.Services = append(getHerd.Data.Services, metadata_name)
// 				getHerd_data := getHerd.PostData()
// 				var cpuAndMem struct {
// 					cpu []int
// 					mem []int
// 				}

// 				for _, service := range getHerd_data.Services {
// 					if len(service.Clusters) > 0 {
// 						for _, cluster :=  range service.Clusters {
// 							if len(cluster.Containers) > 0 {
// 								for _, container := range cluster.Containers {
// 									cpuAndMem.cpu = append(cpuAndMem.cpu, container.Cpu)
// 									cpuAndMem.mem = append(cpuAndMem.cpu, container.Mem)
// 								}
// 							}
// 						}
// 					}
// 				}

				
// 				for _, port := range data_items.Spec.Ports {
// 					sub_Port.Port = port.Port
// 					sub_Port.External_port = port.NodePort
// 					_service.Ports = append(_service.Ports, sub_Port)
// 				}

// 				if len(data_items.Spec.Selector) < 1 {
// 					_service.Status = "Forward"
// 					_service.Cpu = 0
// 					_service.Memory = 0
// 				} else {
// 					cpu := strconv.Itoa(len(cpuAndMem.cpu))
// 					mem := strconv.Itoa(len(cpuAndMem.mem))
// 					_service.Status = cpu+"/"+mem
// 					_service.Cpu = len(cpuAndMem.cpu)
// 					_service.Memory = len(cpuAndMem.mem)
// 				}
// 			}
// 		}
// 		_service_list.Service_tab = append(_service_list.Service_tab, _service)
// 	}
//     return _service_list

// 	// response, _ := ioutil.ReadAll(_service_list)
	
// 	// response, _ := ioutil.ReadAll(js)
	
// 	// fmt.Println(data.Items[0].Status.Replicas)

//     // fmt.Println(data.ApiVersion)
// }

func (this *ServiceService) Service() ServiceList {
    service_model := k8sModel.ServiceModel{}
    data := service_model.GetServiceList()
	var _service_list ServiceList
	_service_list.Status = "ok"
 	for _, data_items := range data.Items {
		metadata_name := data_items.Metadata.Name
		if metadata_name == "kube-dns" || metadata_name == "kubernetes" {
			continue
		}

		data := service_model.GetService(metadata_name)
		_service := this.service_by_name(metadata_name, data)
		_service_list.Data = append(_service_list.Data, _service)
	}
    return _service_list

	// response, _ := ioutil.ReadAll(_service_list)
	
	// response, _ := ioutil.ReadAll(js)
	
	// fmt.Println(data.Items[0].Status.Replicas)

    // fmt.Println(data.ApiVersion)
}

func (this *ServiceService) ServiceByName(name string) Service {
	service_model := k8sModel.ServiceModel{}
	data := service_model.GetService(name)
	Test(data)
	_service := this.service_by_name(name, data)
	_service.Status = "ok"

	return _service
}

func (this *ServiceService) CreateService(data []byte) Service {
	type service_format struct {
		Service_name string `json:"service_name"`
		Port string `json:"port"`
		Label map[string]string `json:"label"`
		NodePort string `json:"nodePort"`
	}

    var _service_format service_format
	json.Unmarshal(data, &_service_format)

	var _k8s_service k8s_format.Service
	var _k8s_service_ports k8s_format.ServicePort
	_k8s_service.Kind = "Service"
	// _k8s_service.APIVersion = "v1beta3"
	_k8s_service.Metadata.Name = _service_format.Service_name
	_k8s_service.Metadata.Labels = _service_format.Label
	_k8s_service.Metadata.Labels["name"] = _service_format.Service_name

	_k8s_service_ports.Port, _ = strconv.Atoi(_service_format.Port)
	_k8s_service_ports.TargetPort, _ = strconv.Atoi(_service_format.Port)
	_k8s_service_ports.Protocol = "TCP"
	_k8s_service.Spec.Ports = append(_k8s_service.Spec.Ports, _k8s_service_ports)

	_k8s_service.Spec.Selector = map[string]string{"name":_service_format.Service_name}
	_k8s_service.Spec.Type = "NodePort"

	service_model := k8sModel.ServiceModel{}
    // rep := service_model.CreateService(_k8s_service)
    service_model.CreateService(_k8s_service)
    var _service Service
    // if rep == nil {
    // 	_service.Status = "fail"
    // 	return _service
    // }

	_service.Status = "ok"
	return _service
}

func (this *ServiceService) UpdateService(name string, data []byte) Service {
	type service_format struct {
		Service_name string `json:"service_name"`
		Port string `json:"port"`
		Label map[string]string `json:"label"`
		NodePort string `json:"nodePort"`
	}
	var _service Service

    var _service_format service_format
	json.Unmarshal(data, &_service_format)

	service_model := k8sModel.ServiceModel{}
	_k8s_service := service_model.GetService(name)

	_k8s_service.Metadata.Labels = _service_format.Label
	_k8s_service.Spec.Ports[0].Port, _ = strconv.Atoi(_service_format.Port)
	_k8s_service.Spec.Ports[0].TargetPort, _ = strconv.Atoi(_service_format.Port)
	_k8s_service.Spec.Selector["name"] = _service_format.Service_name
	
    service_model.UpdateService(name, _k8s_service)
	_service.Status = "ok"
	// _service.Metadata.Label = _service_format.Label
	return _service
}

func (this *ServiceService) DeleteService(name string) Service {
	service_model := k8sModel.ServiceModel{}
	// _service := service_model.DeleteService(name)
	service_model.DeleteService(name)
	var _service Service
	_service.Status = "ok"

	return _service
}

func (this *ServiceService) service_by_name(name string, data k8s_format.Service) Service {
    // service_model := k8sModel.ServiceModel{}
    // data := service_model.GetService(name)
	// var _service_list ServiceList

	var _service Service

	metadata_name := data.Metadata.Name
	if metadata_name == "kube-dns" || metadata_name == "kubernetes" {
		_service.Status = "fail"
		_service.Errmsg = "no found service"
		return _service
	}

	_service.Service_name = metadata_name
	_service.Label = data.Metadata.Labels

	var herd_service herd_format.Data
	herd_service.Services = append(herd_service.Services, metadata_name)

	var getHerd herdModel.GetHerdModel
	getHerd_data := getHerd.PostData(herd_service)
	var cpuAndMem struct {
		cpu []int
		mem []int
	}

	for _, service := range getHerd_data.Services {
		if len(service.Clusters) > 0 {
			for _, _cluster :=  range service.Clusters {
				if len(_cluster.Containers) > 0 {
					for _, _container := range _cluster.Containers {
						cpuAndMem.cpu = append(cpuAndMem.cpu, _container.Cpu)
						cpuAndMem.mem = append(cpuAndMem.mem, _container.Mem)
					}
				}
			}
		}
	}

	var _sub_Port sub_Ports
	for _, port := range data.Spec.Ports {
		_sub_Port.Port = port.Port
		_sub_Port.External_port = port.NodePort
		_service.Ports = append(_service.Ports, _sub_Port)
	}

	if len(data.Spec.Selector) < 1 {
		_service.MachineStatus = "Forward"
		_service.Cpu = 0
		_service.Memory = 0
	} else {
		cpu := strconv.Itoa(len(cpuAndMem.cpu))
		mem := strconv.Itoa(len(cpuAndMem.mem))
		_service.MachineStatus = cpu+"/"+mem
		_service.Cpu = len(cpuAndMem.cpu)
		_service.Memory = len(cpuAndMem.mem)
	}
		
    return _service
}

func Test(result interface{}) {
	js, _ := json.Marshal(result)
	fmt.Println(string(js))
}

	

func (this *ServiceService) GetCluster() string {
	return this.GetBaseUrl()
}

// func (this *ServiceService) GetContainer() string {
// 	return this.GetBaseUrl()
// }

// func (this *ServiceService) GetMachine() string {
// 	return this.GetBaseUrl()
// }
