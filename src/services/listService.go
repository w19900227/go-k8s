package services

import (
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
	bs "./baseService"
	"../models/k8sModel"
	"../models/herdModel"
	// k8s "./k8s"
	"strconv"

)

type ListService struct {
	bs.BaseService
	// bm.BaseModel
}

// func (this *BaseService) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

func (this *ListService) GetTest() string {
	getHerd := herdModel.GetHerdModel{}
	a := getHerd.GetPagInfo()
	fmt.Println(a.Services)
	// this.Test(a)
	return "ss"
}

type Test2 struct {
	Status string `json:"status"`
}

type labels struct {
	Label_key string `json:"label_key"`
	Label_value string `json:"label_value"`
}
type sub_Ports struct {
	Port int `json:"port"`
	External_port int `json:"external_port"`
}
type service_lists struct {
	Label []labels `json:"label"`
	Ports []sub_Ports `json:"ports"`
	Status string `json:"status"`
	Cpu int `json:"cpu"`
	Memory int `json:"memory"`
	Service_name string `json:"service_name"`
}	
type results struct {
	Service_tab []service_lists `json:"service_tab"`
}
func (this *ListService) Service() results {
    service_model := k8sModel.ServiceModel{}
    data := service_model.GetServiceList()

    // if len(data.Items) < 1 {
    // 	fmt.Println("not found data")
    // 	return 
    // }
	
	var result results
	var service_list service_lists
	var label labels
	var sub_Port sub_Ports

 	for _, data_items := range data.Items {
		metadata_name := data_items.Metadata.Name
		service_list.Service_name = metadata_name
		// fmt.Println(metadata_name)
		if metadata_name != "kube-dns" && metadata_name != "kubernetes" {
			// fmt.Println(metadata_name)
			if data_items.Metadata.Labels == nil {
				label.Label_key = "Null"
				label.Label_value = "Null"
				service_list.Label = append(service_list.Label, label)
			} else {
				for key, value := range data_items.Metadata.Labels {
					label.Label_key = key
					label.Label_value = value
					service_list.Label = append(service_list.Label, label)
				}

				var getHerd herdModel.GetHerdModel
				getHerd.Data.Services = append(getHerd.Data.Services, metadata_name)
				getHerd_data := getHerd.PostData()
				var cpuAndMem struct {
					cpu []int
					mem []int
				}

				for _, service := range getHerd_data.Services {
					if len(service.Clusters) > 0 {
						for _, cluster :=  range service.Clusters {
							if len(cluster.Containers) > 0 {
								for _, container := range cluster.Containers {
									cpuAndMem.cpu = append(cpuAndMem.cpu, container.Cpu)
									cpuAndMem.mem = append(cpuAndMem.cpu, container.Mem)
								}
							}
						}
					}
				}

				
				for _, port := range data_items.Spec.Ports {
					sub_Port.Port = port.Port
					sub_Port.External_port = port.NodePort
					service_list.Ports = append(service_list.Ports, sub_Port)
				}

				if len(data_items.Spec.Selector) < 1 {
					service_list.Status = "Forward"
					service_list.Cpu = 0
					service_list.Memory = 0
				} else {
					cpu := strconv.Itoa(len(cpuAndMem.cpu))
					mem := strconv.Itoa(len(cpuAndMem.mem))
					service_list.Status = cpu+"/"+mem
					service_list.Cpu = len(cpuAndMem.cpu)
					service_list.Memory = len(cpuAndMem.mem)
				}
			}
		}
		result.Service_tab = append(result.Service_tab, service_list)
	}
    return result

	// response, _ := ioutil.ReadAll(result)
	
	// response, _ := ioutil.ReadAll(js)
	
	// fmt.Println(data.Items[0].Status.Replicas)

    // fmt.Println(data.ApiVersion)
}


func Test(result interface{}) {
	js, _ := json.Marshal(result)
	fmt.Println(string(js))
}

type cluster_labels struct {
	Label_key string `json:"label_key"`
	Label_value string `json:"label_value"`
}
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
type cluster_service_lists struct {
	Label []cluster_labels `json:"label"`
	Auto_scale int `json:"auto_scale"`
	Cluster_id string `json:"cluster_id"`
	Container_info []container_infos `json:"container_info"`
	Cpu int `json:"cpu"`
	Image string `json:"image"`
	Memory int `json:"memory"`
	Selector []selectors `json:"selector"`
	Status string `json:"status"`

}	
	
type cluster_results struct {
	Service_tab []cluster_service_lists `json:"service_tab"`
}
func (this *ListService) Cluster() string {
	// replicationcontrollers
	cluster_model := k8sModel.ClusterModel{}
    data := cluster_model.GetClusterList()
    // Test(data)


	var result cluster_results
	var service_list cluster_service_lists
    var label cluster_labels

    for _, data_items := range data.Items {
		metadata_name := data_items.Metadata.Name
		// fmt.Println(metadata_name)
		if metadata_name != "kube-dns-v3" {
			if data_items.Metadata.Labels == nil {
				label.Label_key = "Null"
				label.Label_value = "Null"
				service_list.Label = append(service_list.Label, label)
			} else {
				for key, value := range data_items.Metadata.Labels {
					label.Label_key = key
					label.Label_value = value
					service_list.Label = append(service_list.Label, label)
				}

				// for _, value := range data_items.Spec.Template.Spec.Containers {
					// service_list.Images = value.Image
				// }
				replicas := strconv.Itoa(data_items.Status.Replicas)
				service_list.Status = replicas + "/" + replicas

				var getScaleHerd herdModel.GetScaleHerdModel
				getScaleHerd.Data.Clusters = append(getScaleHerd.Data.Clusters, metadata_name)
				getScaleHerd_data := getScaleHerd.PostData()
				if getScaleHerd_data.Clusters == nil {
					service_list.Auto_scale = 0
				} else {
					// service_list.Auto_scale = getScaleHerd_data.Clusters[0].Enable_auto_scale
				}
				

			}

		}
		result.Service_tab = append(result.Service_tab, service_list)
    }
    // fmt.Println(result)
    Test(result)

	return this.GetBaseUrl()
}

func (this *ListService) GetCluster() string {
	return this.GetBaseUrl()
}

// func (this *ListService) GetContainer() string {
// 	return this.GetBaseUrl()
// }

// func (this *ListService) GetMachine() string {
// 	return this.GetBaseUrl()
// }
