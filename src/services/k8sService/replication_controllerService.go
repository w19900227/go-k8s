package k8sService

import (
	"fmt"
)

import (
	// "fmt"
	bs "services/baseService"
	"models/k8sModel"
	"models/herdModel"
	herd_format "models/herdModel/format"
	k8s_format "models/k8sModel/format"
	"strconv"
	"encoding/json"
	"strings"

	// "../models"
	// k8s "./k8s"
)

type ReplicationControllerService struct {
	bs.BaseService
	// k8s.GetK8s
	// bm.BaseModel
}

// func (this *ReplicationControllerService) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }
type image_file struct {
	Name string `json:"name,omitempty"`
	Image string  `json:"image,omitempty"`
}
type pod struct {
	Container_name string `json:"container_name,omitempty"`
	Image_file []image_file `json:"image_file,omitempty"`
	Label map[string]string `json:"label,omitempty"`
	Public_ip string `json:"public_ip,omitempty"`
	Host_ip string `json:"host_ip,omitempty"`
	Mem int `json:"mem"`
	Cpu int `json:"cpu"`
	Status string `json:"status,omitempty"`
}

type Cluster_format struct {
	Auto_scale int `json:"auto_scale"`
	Pods []pod `json:"container_info,omitempty"`
	Status string `json:"status,omitempty"`

	Cluster_name string `json:"cluster_name,omitempty"`
	Port int `json:"port,omitempty"`
	Container_count int `json:"container_count"`
	Image string `json:"image,omitempty"`
	Mem int `json:"mem"`
	Cpu int `json:"cpu"`
	Selector map[string]string `json:"selector,omitempty"`
	Label map[string]string `json:"label,omitempty"`
}

type Cluster struct {
	Status string `json:"status,omitempty"`
	Errno string `json:"errno,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	Data Cluster_format `json:"data,omitempty"`
}
	
type ClusterList struct {
	Status string `json:"status,omitempty"`
	Errno string `json:"errno,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	Data []Cluster_format `json:"data,omitempty"`
}


func (this *ReplicationControllerService) GetTest() string {
	return "ReplicationControllerService"
}

func (this *ReplicationControllerService) GetReplicationControllerList() ClusterList {
	// replicationcontrollers
	_replication_controllerModel := k8sModel.ReplicationControllerModel{}
    data := _replication_controllerModel.GetReplicationControllerList()

	_podModel := k8sModel.PodModel{}
    _pod_data := _podModel.GetPodList()

	var _cluster_list ClusterList

    for _, data_items := range data.Items {
		metadata_name := data_items.Metadata.Name
    	if metadata_name == "kube-dns-v3" {
    		continue
    	}

		var _cluster Cluster_format
		_cluster = this.RcByName(data_items)

		replicas := strconv.Itoa(data_items.Status.Replicas)
		_cluster.Status = replicas + "/" + replicas

		var herd_service herd_format.Data
		herd_service.Clusters = append(herd_service.Clusters, metadata_name)

		var scale_herd herdModel.ScaleHerdModel
		scale_herd_data := scale_herd.GetScale(herd_service)

		if len(scale_herd_data.Clusters) > 0 {
			_cluster.Auto_scale = scale_herd_data.Clusters[0].Enable_auto_scale
		}

		for _, _pod_data_item := range _pod_data.Items {
			_metadata_name_arr := strings.Split(_pod_data_item.Metadata.Name, "-")
			_metadata_name := strings.Join(_metadata_name_arr[:len(_metadata_name_arr)-1], "-")
			if metadata_name != _metadata_name {
				continue
			}
			var _pod pod
			_pod.Container_name = _pod_data_item.Metadata.Name
			_pod.Label = _pod_data_item.Metadata.Labels
			_pod.Public_ip = _pod_data_item.Status.PodIP
			_pod.Host_ip = _pod_data_item.Status.HostIP
			_pod.Status = _pod_data_item.Status.Phase

			var getHerd herdModel.GetHerdModel
			getHerd_data := getHerd.PostData(herd_service)

			for _, cluster := range getHerd_data.Clusters {
				if len(cluster.Containers) > 0 {
					for _, _container := range cluster.Containers {
						_cluster.Cpu += _container.Cpu
						_cluster.Mem += _container.Mem
						if _pod_data_item.Metadata.Name == _container.Container_name {
							_pod.Cpu = _container.Cpu
							_pod.Mem = _container.Mem
						}
					}
				}
			}
			var _image_file image_file
			_image_file.Name = _pod_data_item.Spec.Containers[0].Name
			_image_file.Image = _pod_data_item.Spec.Containers[0].Image
			_pod.Image_file = append(_pod.Image_file, _image_file)
			
			_cluster.Pods = append(_cluster.Pods, _pod)
		}

		_cluster_list.Status = "ok"
		_cluster_list.Data = append(_cluster_list.Data, _cluster)
    }

	return _cluster_list
}

func sum(a []int) int {
	var i int
	for _, v := range a {
		i += v
	}
	return i
}


func (this *ReplicationControllerService) GetReplicationController(rc_name string) Cluster {
	_replication_controllerModel := k8sModel.ReplicationControllerModel{}
    data := _replication_controllerModel.GetReplicationController(rc_name)

	var _cluster Cluster
	_cluster.Data = this.RcByName(data)

	_cluster.Status = "ok"
	return _cluster
}

func (this *ReplicationControllerService) RcByName(data k8s_format.ReplicationController) Cluster_format {

	var _cluster Cluster_format
	_cluster.Cluster_name = data.Metadata.Name
	_cluster.Label = data.Metadata.Labels
	_cluster.Port = data.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort
	_cluster.Selector = data.Spec.Selector
	_cluster.Image = data.Spec.Template.Spec.Containers[0].Image
	_cluster.Container_count = data.Spec.Replicas
	_cluster.Cpu = 0
	_cluster.Mem = 0
	return _cluster
}

func (this *ReplicationControllerService) CreateReplicationController(data []byte) Cluster {
	type rc_format struct {
		Cluster_name string `json:"cluster_name"`
		Port string `json:"port"`
		Container_count string `json:"container_count"`
		Cpu string `json:"cpu"`
		Image string `json:"image"`
		Label map[string]string `json:"label"`
	}
	var _rc_format rc_format
	json.Unmarshal(data, &_rc_format)

	var k8s_rc k8s_format.ReplicationController
	var k8s_rc_container k8s_format.Container
	var k8s_rc_container_port k8s_format.ContainerPort

	k8s_rc.Kind = "ReplicationController"
	k8s_rc.Metadata.Name = _rc_format.Cluster_name
	k8s_rc.Metadata.Labels = _rc_format.Label
	k8s_rc.Metadata.Labels["name"] = _rc_format.Cluster_name
	k8s_rc.Spec.Replicas, _ = strconv.Atoi(_rc_format.Container_count)
	k8s_rc.Spec.Selector = map[string]string{"name":_rc_format.Cluster_name}
	k8s_rc.Spec.Template.Metadata.Labels = map[string]string{"name":_rc_format.Cluster_name}
	k8s_rc_container.Name = _rc_format.Cluster_name
	k8s_rc_container.Image = _rc_format.Image
	k8s_rc_container_port.ContainerPort, _ = strconv.Atoi(_rc_format.Port)
	k8s_rc_container.Ports = append(k8s_rc_container.Ports, k8s_rc_container_port)
	k8s_rc.Spec.Template.Spec.Containers = append(k8s_rc.Spec.Template.Spec.Containers, k8s_rc_container)

	_replication_controllerModel := k8sModel.ReplicationControllerModel{}
    _replication_controllerModel.CreateReplicationController(k8s_rc)

	var _cluster Cluster
	_cluster.Status = "ok"

	return _cluster
}

func (this *ReplicationControllerService) UpdateReplicationController(rc_name string, data []byte) Cluster {
	type rc_format struct {
		// Cluster_name string `json:"cluster_name"`
		Port string `json:"port"`
		Container_count string `json:"container_count"`
		Cpu string `json:"cpu"`
		Image string `json:"image"`
		Label map[string]string `json:"label"`
	}
	var _rc_format rc_format
	json.Unmarshal(data, &_rc_format)

	_replication_controllerModel := k8sModel.ReplicationControllerModel{}
    _k8s_rc := _replication_controllerModel.GetReplicationController(rc_name)

	var _cluster Cluster
	_k8s_rc.Metadata.Labels = _rc_format.Label
	_k8s_rc.Spec.Replicas, _ = strconv.Atoi(_rc_format.Container_count)
	// _k8s_rc.Spec.Selector["name"] = _rc_format.Cluster_name
	// _k8s_rc.Spec.Template.Metadata.Labels["name"] = _rc_format.Cluster_name
	// _k8s_rc.Spec.Template.Spec.Containers[0].Name = _rc_format.Cluster_name
	_k8s_rc.Spec.Template.Spec.Containers[0].Image = _rc_format.Image
	_k8s_rc.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort, _ = strconv.Atoi(_rc_format.Port)
	
	_replication_controllerModel.UpdateReplicationController(rc_name, _k8s_rc)

	_cluster.Status = "ok"

	return _cluster
}

func (this *ReplicationControllerService) DeleteReplicationController(rc_name string) Cluster {
	_replication_controllerModel := k8sModel.ReplicationControllerModel{}
    _k8s_rc := _replication_controllerModel.GetReplicationController(rc_name)
	_k8s_rc.Spec.Replicas = 0
	_replication_controllerModel.UpdateReplicationController(rc_name, _k8s_rc)
    _replication_controllerModel.DeleteReplicationController(rc_name)



	var _cluster Cluster
	_cluster.Status = "ok"

	return _cluster
}

func tt() {
	fmt.Println("ssstttt")
}