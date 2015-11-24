package k8sService

import (
	"fmt"
)

import (
	// "fmt"
	bs "services/baseService"
	"models/k8sModel"
	// "models/herdModel"
	k8s_format "models/k8sModel/format"
	"strconv"
	"encoding/json"

	// herd_format "models/herdModel/format"
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
type container_infos struct {
	container_id string `json:"container_id,omitempty"`
	image_file [] struct {
		image string  `json:"image,omitempty"`
		name string `json:"name,omitempty"`
	} `json:"image_file,omitempty"`
	label map[string]string `json:"label,omitempty"`
	public_ip string `json:"public_ip,omitempty"`
	status string `json:"status,omitempty"`
}

type Cluster struct {
	Status string `json:"status,omitempty"`
	Errno string `json:"errno,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	// Data Cluster `json:"data,omitempty"`


	Auto_scale int `json:"auto_scale,omitempty"`
	Container_info []container_infos `json:"container_info,omitempty"`
	MachineStatus string `json:"machine_status,omitempty"`

	Cluster_name string `json:"cluster_name,omitempty"`
	Port int `json:"port,omitempty"`
	Container_count int `json:"container_count,omitempty"`
	Image string `json:"image,omitempty"`
	Mem int `json:"mem,omitempty"`
	Cpu int `json:"cpu,omitempty"`
	Selector map[string]string `json:"selector,omitempty"`
	Label map[string]string `json:"label,omitempty"`

}
	
type ClusterList struct {
	Status string `json:"status,omitempty"`
	Errno string `json:"errno,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	Data []Cluster `json:"data,omitempty"`
}


func (this *ReplicationControllerService) GetTest() string {
	return "ReplicationControllerService"
}

func (this *ReplicationControllerService) GetReplicationControllerList() ClusterList {
	// replicationcontrollers
	_replication_controllerModel := k8sModel.ReplicationControllerModel{}
    data := _replication_controllerModel.GetReplicationControllerList()

	var _cluster_list ClusterList

    for _, data_items := range data.Items {
		var _cluster Cluster
		metadata_name := data_items.Metadata.Name

    	if metadata_name == "kube-dns-v3" {
    		continue
    	}

		_cluster.Label = data_items.Metadata.Labels


		// for _, value := range data_items.Spec.Template.Spec.Containers {
			// _cluster.Images = value.Image
		// }
		replicas := strconv.Itoa(data_items.Status.Replicas)
		_cluster.Status = replicas + "/" + replicas



		// var herd_service herd_format.Data
		// herd_service.Clusters = append(herd_service.Clusters, metadata_name)
		// Test(herd_service)

		// var getHerd herdModel.GetScaleHerdModel
		// getHerd_data := getHerd.PostData(herd_service)
		// var getScaleHerd herdModel.GetScaleHerdModel
		// getScaleHerd.Data.Clusters = append(getScaleHerd.Data.Clusters, metadata_name)
		// getScaleHerd_data := getScaleHerd.PostData()
		// if getScaleHerd_data.Clusters == nil {
			// _cluster.Auto_scale = 0
		// } else {
			// _cluster.Auto_scale = getScaleHerd_data.Clusters[0].Enable_auto_scale
		// }

		_cluster_list.Data = append(_cluster_list.Data, _cluster)
    }
    // fmt.Println(_cluster_list)
    // Test(_cluster_list)

	return _cluster_list
}


func (this *ReplicationControllerService) GetReplicationController(rc_name string) Cluster {
	_replication_controllerModel := k8sModel.ReplicationControllerModel{}
    data := _replication_controllerModel.GetReplicationController(rc_name)
    Test(data)

	var _cluster Cluster
	_cluster.Cluster_name = data.Metadata.Name
	_cluster.Label = data.Metadata.Labels
	_cluster.Port = data.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort
	_cluster.Selector = data.Spec.Selector
	_cluster.Image = data.Spec.Template.Spec.Containers[0].Image
	_cluster.Container_count = data.Spec.Replicas
	_cluster.Cpu = 0
	_cluster.Mem = 0

	_cluster.Status = "ok"

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