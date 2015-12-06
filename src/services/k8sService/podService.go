package k8sService

import (
)

import (
	// "fmt"
	// bs "./baseService"
	// "../models"
	// k8s "./k8s"
)

import (
	bs "services/baseService"
	"models/k8sModel"
	"models/herdModel"
	herd_format "models/herdModel/format"
	k8s_format "models/k8sModel/format"
	"strings"
	// "strconv"

	// herd_format "models/herdModel/format"
)

type PodService struct {
	bs.BaseService
	// k8s.GetK8s
	// bm.BaseModel
}

// func (this *PodService) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }
type port struct {
	Phase string `json:"phase,omitempty"`
	Public_ip string `json:"public_ip,omitempty"`
	Host_ip string `json:"host_ip,omitempty"`
}
type Container struct {
	Name string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
}
type Pod_format struct {
	Container_name string `json:"container_name"`
	Service_name string `json:"service_name,omitempty"`
	Cpu int `json:"cpu"`
	Mem int `json:"mem"`
	Ports port `json:"ports,omitempty"`
	Containers []Container `json:"containers,omitempty"`
	Label map[string]string `json:"label,omitempty"`
}
type Pod struct {
	Status string `json:"status,inline"`
	Errno string `json:"errno,inline"`
	Errmsg string `json:"errmsg,inline"`
	Data Pod_format `json:"data,inline"`
}

type PodList struct {
	Status string `json:"status,inline"`
	Errno string `json:"errno,inline"`
	Errmsg string `json:"errmsg,inline"`
	Data []Pod_format `json:"data,inline"`
}

func (this *PodService) GetPodList() PodList {
	_podModel := k8sModel.PodModel{}
	data := _podModel.GetPodList()

	var _pod_list PodList
	for _, data_items := range data.Items {
		_metadata_name_arr := strings.Split(data_items.Metadata.Name, "-")
		_metadata_name := strings.Join(_metadata_name_arr[:len(_metadata_name_arr)-1], "-")
		if _metadata_name == "kube-dns-v3" {
			continue
		}

		data := _podModel.PodByName(data_items.Metadata.Name)
		_pod := this.PodByName(data_items.Metadata.Name, data)
		_pod_list.Data = append(_pod_list.Data, _pod)
	}
	return _pod_list
}

func (this *PodService) GetPod(name string) Pod {
	var _pod Pod

	_podModel := k8sModel.PodModel{}
	data := _podModel.PodByName(name)

	_metadata_name_arr := strings.Split(data.Metadata.Name, "-")
	_metadata_name := strings.Join(_metadata_name_arr[:len(_metadata_name_arr)-1], "-")
	if _metadata_name == "kube-dns-v3" {
		_pod.Status = "fail"
		return _pod
	}

	_pod.Data = this.PodByName(name, data)
	_pod.Status = "ok"
	return _pod
}

func (this *PodService) DeletePod(name string) Pod {

	_podModel := k8sModel.PodModel{}
	_podModel.DeletePod(name)

	var _pod Pod
	_pod.Status = "ok"
	return _pod
}

func (this *PodService) PodByName(name string, data k8s_format.Pod) Pod_format {
	
	var _pod Pod_format
	// _metadata_name_arr := strings.Split(data.Metadata.Name, "-")
	// _metadata_name := strings.Join(_metadata_name_arr[:len(_metadata_name_arr)-1], "-")
	// if _metadata_name == "kube-dns-v3" {
	// 	return _pod
	// }
	_pod.Label = data.Metadata.Labels

	_pod.Container_name = data.Metadata.Name
	_service_name_arr := strings.Split(data.Metadata.GenerateName, "-")
	_service_name := strings.Join(_service_name_arr[:len(_service_name_arr)-1], "-")
	_pod.Service_name = _service_name

	var herd_service herd_format.Data
	herd_service.Containers = append(herd_service.Containers, data.Metadata.Name)

	var getHerd herdModel.GetHerdModel
	getHerd_data := getHerd.PostData(herd_service)
	_herd_container := getHerd_data.Containers[0]
	_pod.Cpu = _herd_container.Cpu
	_pod.Mem = _herd_container.Mem

	_pod.Ports.Phase = data.Status.Phase
	_pod.Ports.Public_ip = data.Status.PodIP
	_pod.Ports.Host_ip = data.Status.HostIP
	// Test(data.Status.PodIP)

	// if data.Status.PodIP != nil {
	// 	_pod.Status.Public_ip = data.Status.PodIP
	// }
	// if data.Status.HostIP != nil {
	// 	_pod.Status.Host_ip = "2"
	// }

	var _container Container
	for _, _spec_container := range data.Spec.Containers {
		_container.Name = _spec_container.Name
		_container.Image = _spec_container.Image
	}
	_pod.Containers = append(_pod.Containers, _container)
	return _pod
}

