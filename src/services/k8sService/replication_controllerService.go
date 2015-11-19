package k8sService

import (
)

import (
	// "fmt"
	bs "services/baseService"
	"models/k8sModel"
	// "models/herdModel"
	"strconv"

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
	container_id string `json:"container_id"`
	image_file [] struct {
		image string  `json:"image"`
		name string `json:"name"`
	} `json:"image_file"`
	label [] struct {
		label_key string `json:"label_key"`
		label_value string  `json:"label_value"`
	} `json:"label"`
	public_ip string `json:"public_ip"`
	status string `json:"status"`
}
type selectors struct {
	label_key string `json:"label_key"`
	label_value string `json:"label_value"`
}

type Cluster struct {
	Label map[string]string `json:"label"`
	Auto_scale int `json:"auto_scale"`
	Cluster_id string `json:"cluster_id"`
	Container_info []container_infos `json:"container_info"`
	Cpu int `json:"cpu"`
	Image string `json:"image"`
	Memory int `json:"memory"`
	Selector []selectors `json:"selector"`
	Status string `json:"status"`

}
	
type ClusterList struct {
	Service_tab []Cluster `json:"data"`
}


func (this *ReplicationControllerService) GetTest() string {
	return "ReplicationControllerService"
}

func (this *ReplicationControllerService) ReplicationController() ClusterList {
	// replicationcontrollers
	_replication_controllerModel := k8sModel.ReplicationControllerModel{}
    data := _replication_controllerModel.ReplicationController()

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

		_cluster_list.Service_tab = append(_cluster_list.Service_tab, _cluster)
    }
    // fmt.Println(_cluster_list)
    // Test(_cluster_list)

	return _cluster_list
}