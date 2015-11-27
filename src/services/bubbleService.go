package services

import (
	"fmt"

// 	"runtime"
// 	"time"
)

import (
	bs "services/baseService"
	// "encoding/json"
	"models/k8sModel"
	"services/k8sService"
	// "models/herdModel"
	// "strings"
	// "strconv"
	// herd_format "models/herdModel/format"

)

type BubbleService struct {
	bs.BaseService
}

type item struct {
	Services []string `json:"services,omitempty"`
	Clusters []string `json:"clusters,omitempty"`
	Containers []string `json:"containers,omitempty"`
}
type relate struct {
	Cluster_name string `json:"cluster_name"`
	Service []string `json:"service"`
	Container []string `json:"container"`
}

type data struct {
	Items item `json:"items,omitempty"`
	Relates []relate `json:"relates,omitempty"`
	Services []k8sService.Service_format `json:"services,omitempty"`
	Clusters []k8sService.Cluster_format `json:"clusters,omitempty"`
	Pods []k8sService.Pod_format `json:"containers,omitempty"`
}
type bubble struct {
	Status string `json:"status,omitempty"`
	Errno string `json:"errno,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	Data data `json:"data,omitempty"`
}


func (this *BubbleService) GetRelate() bubble {
	fmt.Println("[Get] Relate Service")

    var _bubble bubble
    var _item item

	var service_model k8sModel.ServiceModel
    service_data := service_model.GetServiceList()
    // cluster_to_service := map[string]string{}
    for _, service := range service_data.Items {
    	_item.Services = append(_item.Services, service.Metadata.Name)
    }

	var pod_model k8sModel.PodModel
    pod_data := pod_model.GetPodList()
    // Test(pod_data)
    for _, pod := range pod_data.Items {
    	// Test(pod)
    	_item.Containers = append(_item.Containers, pod.Metadata.Name)
    }

	var rc_model k8sModel.ReplicationControllerModel
    rc_data := rc_model.GetReplicationControllerList()
    for _, rc := range rc_data.Items {
    	_item.Clusters = append(_item.Clusters, rc.Metadata.Name)
    	var _relate relate
    	_relate.Cluster_name = rc.Metadata.Name
    	for _, service := range service_data.Items {
	    	// _item.Services = append(_item.Services, service.Metadata.Name)
	    	if service.Spec.Selector != nil && 
	    	service.Spec.Selector["name"] != "" && 
	    	rc.Spec.Selector["name"] == service.Spec.Selector["name"] {
	    		_relate.Service = append(_relate.Service, service.Metadata.Name)
	    	}
	    }

	    for _, pod := range pod_data.Items {
	    	if pod.Metadata.Labels != nil && 
	    	pod.Metadata.Labels["name"] != "" && 
	    	rc.Spec.Selector["name"] == pod.Metadata.Labels["name"] {
	    		_relate.Container = append(_relate.Container, pod.Metadata.Name)
	    	}
	    }

    	_bubble.Data.Relates = append(_bubble.Data.Relates, _relate)
    }

    _bubble.Data.Items = _item
	_bubble.Status = "ok"
	return _bubble
}
func (this *BubbleService) GetInfo() bubble {
	fmt.Println("[Get] Info Service")

	var _bubble bubble

	var rc_service k8sService.ReplicationControllerService
	rc_data := rc_service.GetReplicationControllerList()
    _bubble.Data.Clusters = rc_data.Data

	var ser_service k8sService.ServiceService
	ser_data := ser_service.GetServiceList()
    _bubble.Data.Services = ser_data.Data

	var pod_service k8sService.PodService
	pod_data := pod_service.GetPodList()
    _bubble.Data.Pods = pod_data.Data

	_bubble.Status = "ok"
	return _bubble
}
